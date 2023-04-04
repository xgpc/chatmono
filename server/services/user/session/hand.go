/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/24 16:01
 */

package session

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/exce"
	"time"
)

const KeyDir = "token:"

var (
	_rdb *redis2.Client
)

func redis() *redis2.Client {
	return _rdb
}

func Init(rdb *redis2.Client) {
	_rdb = rdb
}

// Set 设置数据
func Set(key string, data interface{}) {
	_, err := redis().Set(context.Background(), KeyDir+key, data, time.Hour*720).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoLogin, "设置token失败, 请联系管理员")
	}
}

// Get 数据
func Get(key string) interface{} {

	data, err := redis().Get(context.Background(), KeyDir+key).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoLogin, "获取token失败, 请重新登录")
	}
	return data
}

// Expire 更新Key
func Expire(key string) {
	_, err := redis().Expire(context.Background(), KeyDir+key, time.Hour*720).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoLogin, "刷新token失败, 请联系管理员")
	}
}
