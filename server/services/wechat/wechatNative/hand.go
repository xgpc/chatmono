package wechatNative

import (
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	_db   *gorm.DB
	_conf *Config
	_rdb  *redis2.Client
)

type Config struct {
	AppID     string `json:"app_id"`
	MchID     string `json:"mch_id"`
	Key       string `json:"key"`
	NotifyURL string `json:"notify_url"`
}

func Init(db *gorm.DB, rdb *redis2.Client, conf Config) {
	_rdb = rdb
	_db = db
}

func db() *gorm.DB {
	return _db
}

func redis() *redis2.Client {
	return _rdb
}
