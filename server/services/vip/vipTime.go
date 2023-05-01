package vip

import (
	"chatmono/pkg"
	"context"
	"errors"
	"github.com/xgpc/dsg/exce"
	"strconv"
	"time"
)

// 会员时间

type Info struct {
	UserID       uint32 // 用户ID
	VipEndTimeAt int64  // vip到期时间
}

// TimeAdd 添加有效时间
func TimeAdd(userID uint32, num time.Duration) error {
	IdStr := strconv.Itoa(int(userID))
	info := TimeGet(userID)
	at := pkg.TimeGetChineseTime()
	if info.VipEndTimeAt <= at.Unix() {
		info.VipEndTimeAt = at.Unix()
	}

	newAt := time.Unix(info.VipEndTimeAt, 0)
	newAt.Add(num)

	info.VipEndTimeAt = newAt.Unix()

	result, err := redis().Set(context.Background(), "vip:"+IdStr, info.VipEndTimeAt, 0).Result()
	if result != "ok" {
		return errors.New("添加会员时间失败")
	}
	return err
}

// CheckVip 判断是否在会员时间内
func CheckVip(userID uint32) error {
	info := TimeGet(userID)
	at := pkg.TimeGetChineseTime()
	if info.VipEndTimeAt >= at.Unix() {
		return nil
	}

	return errors.New("会员已过期")
}

// TimeGet 剩余有效时间
func TimeGet(userID uint32) Info {
	var info Info
	IdStr := strconv.Itoa(int(userID))
	i, err := redis().Get(context.Background(), "vip:"+IdStr).Int64()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, "获取VIP数据出错:"+err.Error())
	}

	info.VipEndTimeAt = i
	info.UserID = userID

	return info
}
