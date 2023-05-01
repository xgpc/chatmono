package vip

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
)

func Check(ctx iris.Context) {
	userID := ctx.Values().GetUint32Default("mid", 0)

	err := CheckVip(userID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, "不是会员, 请充值后重试")
	}

	ctx.Next()
}
