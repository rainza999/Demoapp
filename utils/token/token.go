package token

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
	RtExpires   int64
}
type AccessDetails struct {
	AccessUuid string
	UserId     string
	Name       string
}

func GenerateToken(user_id, name string) (*TokenDetails, error) {

	var err error

	lifespan, err := strconv.Atoi(os.Getenv("MINUTE_LIFESPAN"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(lifespan)).Unix()
	td.AccessUuid = uuid.NewV4().String()

	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = user_id
	atClaims["name"] = name
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return td, nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		log.Println(token)
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyTokenApi(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return token, nil
}

func TokenValidApi(c *gin.Context) error {
	token, err := VerifyTokenApi(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadataApi(c *gin.Context) (*AccessDetails, error) {
	token, err := VerifyTokenApi(c)
	if err != nil {
		log.Printf("token is nil and %s", err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}
		name, ok := claims["name"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
			Name:       name,
		}, nil
	}
	return nil, err
}
