package conf

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus" //github.com/sirupsen/logrus
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//初始化
var (
	DB            *gorm.DB
	MongoDBClient *mongo.Client
	RedisClient   *redis.Client
)

func Mysql_Conn(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println("connect err:", err)
	}
	if err != nil {
		panic(err)
	}
	DB = db
	logging.Info("MySQL successfully connect")
}

func MongoDB_Conn(connString string) {
	clientOptions := options.Client().ApplyURI(connString)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	logging.Info("MongoDB successfully connect")
}

func Redis_Conn(RedisAddr string, RedisDbName string, RedisPw string) {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) //string to uint64
	client := redis.NewClient(&redis.Options{       //登录Redis
		Addr:     RedisAddr,
		Password: RedisPw, // 无密码，注释掉就好了
		DB:       int(db),
	})
	_, err := client.Ping().Result() //验证是否ping通
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
