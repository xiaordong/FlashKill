package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func init() {
	initRedis()
	initMySQL()

}
func initMySQL() {
	user := "root"
	password := "2132047479"
	port := "3306"
	host := "localhost"
	database := "flashkill"
	charset := "utf8mb4"
	parseTime := "True"
	loc := "Local"
	dns := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", user, password, host, port, database, charset, parseTime, loc)
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		log.Fatal(err)
		os.Exit(5)
	}
	DB = db.Debug()

}
func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		PoolSize: 20,
	})
	if err := rdb.Ping().Err(); err != nil {
		log.Fatal("redis connect error")
	}
	RDB = rdb

}
