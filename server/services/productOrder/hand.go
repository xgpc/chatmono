package productOrder

import (
	"chatmono/services/order"
	"chatmono/services/order/wechatPly"
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	_db *gorm.DB
)

func Init(conf wechatPly.Config, db *gorm.DB, rdb *redis2.Client) {
	// 微信支付
	order.Init(conf, rdb)
	_db = db
	// 启动任务
	task()
}

func db() *gorm.DB {
	return _db
}
