/**
 * @Author: smono
 * @Description:
 * @File:  middleware
 * @Version: 1.0.0
 * @Date: 2022/9/28 21:32
 */

package openAI

import (
	"github.com/kataras/iris/v12"
)

func SetSessionNum(ctx iris.Context, num int64) {
	ctx.Values().Set("num", num)
}

func GetSessionNum(ctx iris.Context) int64 {
	return ctx.Values().GetInt64Default("num", 0)
}

func GetMyID(ctx iris.Context) uint32 {
	return ctx.Values().GetUint32Default("mid", 0)
}

func SessionNum(ctx iris.Context) {
	defer func() {
		num := GetSessionNum(ctx)
		userID := GetMyID(ctx)
		SessionNumSub(userID, num)
		// 扣次数
	}()

	ctx.Next()
}

func SessionMember(ctx iris.Context) {

	// 判断是否是会员
	// 不是会员就扣次数

	ctx.Next()
}
