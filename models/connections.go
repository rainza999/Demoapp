package models

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var (
	DB     *gorm.DB
	ctx    = context.Background()
	rdb    *redis.Client
	Mailer *gomail.Dialer
)

func ConnectMysql() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}
	// Conncet to Mysql
	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		log.Printf("Cannot connect to database: %s:%s\n", DbHost, DbPort)
	} else {
		log.Printf("Successfully connected to database: %s:%s\n", DbHost, DbPort)
	}
}

func ConnectRedis() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}
	// Conncet to Redis
	RdHost := os.Getenv("REDIS_HOST")
	RdPassword := os.Getenv("REDIS_PASSWORD")
	RdIndex := 1 // Redis support 16 database. You can switch a DB using an integer starting from 0 to 15
	RdPort := os.Getenv("REDIS_PORT")

	rdb = redis.NewClient(&redis.Options{
		Addr:     RdHost + ":" + RdPort,
		Password: RdPassword,
		DB:       RdIndex,
	})

	pong, err := rdb.Ping(ctx).Result()

	if err != nil || pong == "" {
		log.WithFields(log.Fields{
			"host":  RdHost,
			"index": RdIndex,
			"port":  RdPort,
		}).Fatal(err)
		fmt.Printf("Cannot connect to redis: %s:%s\n", RdHost, RdPort)
	} else {
		log.WithFields(log.Fields{
			"host":  RdHost,
			"index": RdIndex,
			"port":  RdPort,
		}).Info("Successfully connected to redis")
		fmt.Printf("Successfully connected to redis: %s:%s\n", RdHost, RdPort)
	}
}

func ConnectMailer() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}
	// Conncet to Gmail
	MHost := os.Getenv("MAILER_HOST")
	MUsername := os.Getenv("MAILER_USERNAME")
	MPassword := os.Getenv("MAILER_PASSWORD")
	MPort, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))

	mailer := gomail.NewDialer(
		MHost,
		MPort,
		MUsername,
		MPassword,
	)
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	Mailer = mailer
}
