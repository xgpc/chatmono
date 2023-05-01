package vip

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
