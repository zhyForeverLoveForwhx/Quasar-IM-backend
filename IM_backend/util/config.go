package util

import (
	"time"

	"github.com/spf13/viper"
)

//Config stores all configuration of the application
//The values are read by viper from a config file or environment varibles
//Mysql: mysql://username:password@xxx.xxx.xxx.xxx:port/DB_Name?sslmode=disable&charset=utf8&parseTime=true
//MongoDB: mongodb://username:password@127.0.0.1:27017/admin
type Config struct {
	AppMode             string        `mapstructure:"APP_MODE"`
	HttpPort            string        `mapstructure:"HTTP_PORT"`
	MysqlDBSource       string        `mapstructure:"MYSQL_DB_SOURCE"`
	MongoDBSource       string        `mapstructure:"MONGO_DB_SOURCE"`
	RedisDBSource       string        `mapstructure:"REDIS_DB_SOURCE"`
	DBTESTSOURCE        string        `mapstructure:"DB_TEST_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

//LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
