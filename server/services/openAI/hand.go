package openAI

import (
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Config struct {
	OpenaiApiKey string `yaml:"openai_api_key"`
}

var _conf Config
var _rdb *redis2.Client
var _db *gorm.DB

func Init(db *gorm.DB, rdb *redis2.Client, conf Config) {
	_conf = conf
	_db = db
	_rdb = rdb
}

func redis() *redis2.Client {
	return _rdb
}
func db() *gorm.DB {
	return _db
}
