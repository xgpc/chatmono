/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/13 17:23
 */

package user

import (
	"chatmono/services/user/session"
	redis2 "github.com/go-redis/redis/v8"
)

var (
	_rdb *redis2.Client
)

func Init(rdb *redis2.Client) {
	_rdb = rdb
	session.Init(rdb)
}

func redis() *redis2.Client {
	return _rdb
}
