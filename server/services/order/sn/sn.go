/**
 * @Author: smono
 * @Description:
 * @File:  sn
 * @Version: 1.0.0
 * @Date: 2022/7/2 17:23
 */

package sn

import (
	redis2 "github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/util"
	"strconv"
	"time"
)

var _rdb *redis2.Client

func Init(rdb *redis2.Client) {
	_rdb = rdb
}

func redis() *redis2.Client {
	return _rdb
}

// SN 编号生成器
func timeKey(t int64) string {
	return "KeySN" + util.TimeToString(t)
}

func timeOutKey(t int64) string {
	return "KeyOutSN" + util.TimeToString(t)
}

func GetSn() string {
	t := time.Now().Unix()
	key := timeKey(t)

	orderSn := frame.NewRedisConn(redis()).RedisIncr(key)

	frame.NewRedisConn(redis()).RedisExpire(key, 3600)

	return "INS" + strconv.FormatInt(t, 10) + strconv.FormatInt(orderSn, 10)
}

func GetOutSn() string {
	t := time.Now().Unix()
	key := timeOutKey(t)

	orderSn := frame.NewRedisConn(redis()).RedisIncr(key)

	frame.NewRedisConn(redis()).RedisExpire(key, 3600)

	return "OINS" + strconv.FormatInt(t, 10) + strconv.FormatInt(orderSn, 10)
}
