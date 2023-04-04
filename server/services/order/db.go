/**
 * @Author: smono
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2022/10/11 11:40
 */

package order

import (
	redis2 "github.com/go-redis/redis/v8"
)

var _rdb *redis2.Client

func redis() *redis2.Client {
	return _rdb
}
