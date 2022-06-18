package models

import (
	"errors"
	"fmt"
	"html"
	"net/url"
	"resume/utils/token"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   string `json:"user_id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Useflag  string `json:"userflag"`
	CreateBy string `json:"create_by"`
	CreateAt string `json:"create_at"`
	UpdateBy string `json:"update_by"`
	UpdateAt string `json:"update_at"`
}

type Users struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Useflag  string `json:"userflag"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type Pagination struct {
	Page  int           `json:"page"`
	Limit int           `json:"limit"` // size
	Total int           `json:"total"` // count all
	Sort  string        `json:"sort,omitempty"`
	Links []interface{} `json:"Links"`
}

func (u *User) TableName() string {
	return "users"
}

func CreateUser(username, password, name string) (User, error) {

	user := User{}
	id := uuid.NewV4().String()

	DB.Where("username = ?", username).Find(&user)

	if user.Username != "" {
		err := fmt.Errorf("username not available")
		log.Error(err)
		return user, err
	}

	user.UserId = id
	user.Username = Trim(username)
	user.Password = Hash(password)
	user.Name = Trim(name)
	user.Useflag = "Y"
	user.CreateBy = id
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdateBy = id
	user.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := user.SaveUser()
	if err != nil {
		log.Error(err)
		return user, err
	}

	return user, nil
}

func Login(username, password string) (map[string]string, error) {
	var user User
	data := make(map[string]string)

	err := DB.Model(User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		err := fmt.Errorf("username not found")
		log.Error(err)
		return data, err
	}

	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err := fmt.Errorf("password is incorrect")
		log.Error(err)
		return data, err
	}

	token, err := token.GenerateToken(user.UserId, user.Name)
	if err != nil {
		log.Error(err)
		return data, err
	}

	at := time.Unix(token.AtExpires, 0)
	now := time.Now()

	err = rdb.Set(ctx, token.AccessUuid, user.UserId, at.Sub(now)).Err()
	if err != nil {
		log.Error(err)
		return data, err
	}

	data["access_token"] = token.AccessToken
	data["access_uuid"] = token.AccessUuid
	data["user_id"] = user.UserId
	data["name"] = user.Name
	data["expire"] = at.Format("2006-01-02 15:04:05")

	return data, nil
}

func Profile(user_id string) (*User, error) {
	var user User
	if err := DB.Where(&User{UserId: user_id}).First(&user).Error; err != nil {
		log.Error(err)
		return nil, errors.New("user not found")
	}
	user.PrepareCreateAt()
	user.PrepareUpdateAt()
	user.PrepareGive()
	return &user, nil
}

func ChangePassword(user_id, current, new string) (*User, error) {
	var (
		user User
		err  error
	)

	err = DB.Model(User{}).Where("user_id = ?", user_id).Take(&user).Error
	if err != nil {
		err := fmt.Errorf("user not found")
		log.Error(err)
		return &user, err
	}

	err = VerifyPassword(current, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err := fmt.Errorf("password is incorrect")
		log.Error(err)
		return &user, err
	}

	if err = DB.Model(&user).Where("user_id = ? ", user.UserId).Updates(User{
		Password: Hash(new),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}).Error; err != nil {
		log.Error(err)
		return &user, err
	}
	user.PrepareCreateAt()
	user.PrepareGive()

	return &user, nil
}

func DeleteToken(authD *token.AccessDetails) (map[string]string, error) {
	var (
		resp = make(map[string]string)
		uuid = make(map[string]string)
	)

	get_access, err := rdb.Get(ctx, authD.AccessUuid).Result()
	if err != nil || get_access == "" {
		log.Error(err)
		return resp, err
	}

	allKey := rdb.Scan(ctx, 0, "", 0).Iterator()
	for allKey.Next(ctx) {
		uuid[allKey.Val()] = rdb.Get(ctx, allKey.Val()).Val()
	}
	if err := allKey.Err(); err != nil {
		log.Error(err)
		return resp, err
	}

	if len(uuid) != 0 {
		user_id := authD.UserId
		for index, each := range uuid {
			if each == user_id {
				_, err = rdb.Del(ctx, index).Result()
				if err != nil {
					log.Error(err)
				}
			}
		}
	} else {
		log.Error("not found token")
		return resp, errors.New("not found token")
	}

	return resp, nil
}

func GetListUsers(link, st, et, sort string, pg, limit int) (u []Users, p Pagination, err error) {
	var (
		users       []User
		rows        []Users
		pagination  Pagination
		links       []interface{}
		page, total int
	)

	offset := 0
	if pg != 0 {
		offset = (pg - 1) * limit
		page = pg
	} else {
		page = 1
	}

	if st == "" || et == "" {
		if sort != "desc" {
			err = DB.Order("name").Offset(offset).Limit(limit).Find(&users).Error
			if err != nil {
				log.Error(err)
				return nil, pagination, err
			}
		} else {
			err = DB.Order("name desc").Offset(offset).Limit(limit).Find(&users).Error
			if err != nil {
				log.Error(err)
				return nil, pagination, err
			}
		}
		err = DB.Model(&users).Count(&total).Error
		if err != nil {
			log.Error(err)
			return nil, pagination, err
		}
	} else {
		if sort != "desc" {
			err = DB.Order("name").Where("create_at BETWEEN ? AND ?", st, et).Offset(offset).Limit(limit).Find(&users).Error
			if err != nil {
				log.Error(err)
				return nil, pagination, err
			}
		} else {
			err = DB.Order("name desc").Where("create_at BETWEEN ? AND ?", st, et).Offset(offset).Limit(limit).Find(&users).Error
			if err != nil {
				log.Error(err)
				return nil, pagination, err
			}
		}

		err = DB.Model(&users).Where("create_at BETWEEN ? AND ?", st, et).Count(&total).Error
		if err != nil {
			log.Error(err)
			return nil, pagination, err
		}
	}

	pages := 1
	if total%limit == 0 {
		pages = (pages - 1) + (total / limit)
	} else {
		pages = pages + (total / limit)
	}

	for i := 1; i <= pages; i++ {
		params := url.Values{}
		params.Add("sort", sort)
		params.Add("page", strconv.Itoa(i))
		params.Add("limit", strconv.Itoa(limit))
		if i != page {
			if st != "" && et != "" {
				params.Add("st", st)
				params.Add("et", et)
			}
			links = append(links, link+"?"+params.Encode())
			continue
		}
	}

	pagination.Limit = limit
	pagination.Page = page
	pagination.Total = total
	pagination.Sort = sort
	pagination.Links = links

	currentTime := time.Now()
	loc := currentTime.Location()
	layout := "2006-01-02 15:04:05"

	for _, value := range users {
		cdate, _ := time.ParseInLocation(time.RFC3339, value.CreateAt, loc)
		udate, _ := time.ParseInLocation(time.RFC3339, value.UpdateAt, loc)

		rows = append(rows, Users{
			UserId:   value.UserId,
			Username: value.Username,
			Name:     value.Name,
			Useflag:  value.Useflag,
			CreateAt: cdate.Format(layout),
			UpdateAt: udate.Format(layout),
		})
	}

	return rows, pagination, nil
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		log.Error(err)
		return &User{}, err
	}
	u.PrepareGive()
	return u, nil
}

func Hash(pw string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func Trim(u string) string {
	html.EscapeString(strings.TrimSpace(strings.ReplaceAll(u, " ", "")))
	return u
}

func (u *User) PrepareCreateAt() {

	create := u.CreateAt
	currentTime := time.Now()
	loc := currentTime.Location()
	layout := "2006-01-02 15:04:05"

	t, _ := time.ParseInLocation(time.RFC3339, create, loc)
	u.CreateAt = t.Format(layout)
}

func (u *User) PrepareUpdateAt() {

	update := u.UpdateAt
	currentTime := time.Now()
	loc := currentTime.Location()
	layout := "2006-01-02 15:04:05"

	t, _ := time.ParseInLocation(time.RFC3339, update, loc)
	u.UpdateAt = t.Format(layout)
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func FetchAuth(authD *token.AccessDetails) (string, error) {
	userid, err := rdb.Get(ctx, authD.AccessUuid).Result()
	if err != nil {
		log.Error(err)
		return userid, err
	}
	return userid, nil
}
