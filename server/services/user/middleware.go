package user

import (
	"chatmono/services/user/session"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
)

func Login(ctx iris.Context) {
	token := ctx.GetHeader("token")
	if token != "" {
		id := GetSession(token)
		if id == 0 {
			exce.ThrowSys(exce.CodeUserNoLogin, "请登录后重试")
		}
		ctx.Values().Set("mid", id)
		// 刷新token
		defer session.Expire(token)
	} else {
		exce.ThrowSys(exce.CodeUserNoLogin, "请登录后重试")
	}

	ctx.Next()
}
