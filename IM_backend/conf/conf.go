package conf

import (
	"context"
	"strconv"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//初始化
var (
	MongoDBClient *mongo.Client
	RedisClient   *redis.Client
)

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
