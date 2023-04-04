/**
 * @Author: smono
 * @Description:
 * @File:  middleware
 * @Version: 1.0.0
 * @Date: 2022/9/28 21:32
 */

package admin

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
)

// 所有权限管理 基于会话的ID来操作

func Super(ctx iris.Context) {
	id, err := ctx.Values().GetUint32("mid")
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoAuth, err.Error())
	}

	CheckSuper(id)

	ctx.Next()
}

func Admin(ctx iris.Context) {
	id, err := ctx.Values().GetUint32("mid")
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoAuth, err.Error())
	}

	CheckAdmin(id)

	ctx.Next()
}

func Rule(ctx iris.Context) {
	id, err := ctx.Values().GetUint32("mid")
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoAuth, err.Error())
	}

	CheckRule(id, ctx.Path())
	ctx.Next()
}

func CheckRule(id uint32, path string) {
	fmt.Println(path)
}
