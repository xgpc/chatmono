package score

import (
	"chatmono/services/admin"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	p := api.Party("/score", user.Login)

	// 查询他人积分
	p.Post("/query", admin.Admin, Query)

	// 我的积分获取情况
	p.Post("/list", List)

	// 我的总积分
	p.Post("/total", Total)

	// 积分购买
	p.Post("/pay", Pay)

}
