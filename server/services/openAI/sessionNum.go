package openAI

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"strconv"
)

// SessionNumGet 通过 get key 获取剩余此时
func SessionNumGet(userid uint32) int64 {
	idStr := strconv.Itoa(int(userid))
	i, err := redis().Get(context.Background(), idStr).Int64()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	return i
}

// SessionNumAdd 通过 incrby key 次数 添加次数
func SessionNumAdd(userid uint32, num int64) int64 {
	idStr := strconv.Itoa(int(userid))
	i, err := redis().IncrBy(context.Background(), idStr, num).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	return i
}

// SessionNumSub 通过 decrby key 减少次数
func SessionNumSub(userid uint32, num int64) int64 {
	idStr := strconv.Itoa(int(userid))
	i, err := redis().DecrBy(context.Background(), idStr, num).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	return i
}
