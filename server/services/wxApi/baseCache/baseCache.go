// Package baseInterface
// @Author:        asus
// @Description:   $
// @File:          base
// @Data:          2021/12/2015:39
//
package baseCache

import (
	"dsg/frame"
	"fmt"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/xgpc/util"
	"gorm.io/gorm"
)

type Cache struct {
	db        *gorm.DB
	redisConn *redis.Client
	tag       string
	sync.Mutex
}

func (b *Cache) Tag() string {
	return b.tag
}

func (b *Cache) SetTag(tag string) {

	b.tag = tag
	if strings.Repeat(tag, 1) != ":" {
		b.tag = tag + ":"
	}

}

func (b *Cache) Redis() *frame.RedisCon {
	return frame.NewRedisConn(b.redisConn)
}

func (b *Cache) RD() *redis.Client {
	return b.redisConn
}

func (b *Cache) DB() *gorm.DB {
	return b.db
}

//
//func (b *Cache) Init(db *gorm.DB, conn *redis.Client) {
//	if db == nil {
//		panic(fmt.Errorf("dept ReloadDB gorm.DB == nil"))
//	}
//
//	b.db = db
//	b.redisConn = conn
//}

func (b *Cache) InitTag(db *gorm.DB, conn *redis.Client, tag string) {
	if db == nil {
		panic(fmt.Errorf("dept ReloadDB gorm.DB == nil"))
	}

	b.db = db
	b.redisConn = conn

	b.SetTag(tag)
}

func (b *Cache) cacheTag() string {
	return "cache:" + b.Tag()
}

func (b *Cache) CacheSet(key, data string, sec int) {
	b.Redis().RedisSet(b.cacheTag()+key, data, b.randomSec(sec))
}

func (b *Cache) CacheGet(key string) string {
	return b.Redis().RedisGet(b.cacheTag() + key)
}

func (b *Cache) CacheDel(key string) {
	b.Redis().RedisDel(b.cacheTag() + key)
}

// 添加随机数, 防止缓存击穿
func (b *Cache) randomSec(i int) int {
	return i + util.RandomInt(3)
}
