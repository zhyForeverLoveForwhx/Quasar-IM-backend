package conf

import (
	"context"
	"fmt"

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

func Redis_Conn() {

}
