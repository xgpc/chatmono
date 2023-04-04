// Package wxApi
// @Author:        asus
// @Description:   $
// @File:          wx_test.go
// @Data:          2022/2/1016:30
//
package wxApi

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func getRedis() *redis.Client {
	opt := env.RedisOption{
		Host:         "127.0.0.1",
		Port:         6379,
		Db:           0,
		Password:     "",
		PoolSize:     1,
		MinIdleConns: 5,
	}
	re, err := db.RedisInit(opt)
	if err != nil {
		panic(err)
	}
	return re
}

func Test_wx_Init(t *testing.T) {
	// 测试发送消息

	Handel.InitTag(&gorm.DB{}, getRedis(), "wxToken", "wx236d7a807086d1d2", "7db4fc547d9de2def877b7373f48c798")

	token := Handel.getToken()
	if token == "" {
		t.Fatal("token为空")
	}

	Handel.UserNotice(
		"oCTIB51NOLUzQRH8bWOJz1ynru-w", "sfr", "17610782527", "2022-02-16", "订单备注", "www.baidu.com")
}
