/**
 * @Author: smono
 * @Description:
 * @File:  random
 * @Version: 1.0.0
 * @Date: 2022/10/2 15:05
 */

package random

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/util"
	"time"
)

var _rdb *redis2.Client

func Init(rdb *redis2.Client) {
	_rdb = rdb
}

func redis() *redis2.Client {
	return _rdb
}

func key(mobile string) string {
	return "mobile:" + mobile
}

func CreatedMobileCode(mobile string) string {
	code := createdCode()
	_, err := redis().Set(context.Background(), key(mobile), code, time.Minute*5).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	return code
}

func CheckMobileCode(mobile, code string) bool {
	res, err := redis().Get(context.Background(), key(mobile)).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, fmt.Sprintf("手机验证码获取失败: %s", err.Error()))
	}
	if res == code {
		return true
	}
	return false
}

func createdCode() string {
	code := util.RandomNumber(6)
	return code
}
