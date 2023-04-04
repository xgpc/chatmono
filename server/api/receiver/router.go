package receiver

import (
	"chatmono/services/admin"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	p := api.Party("/receiver", user.Login)

	// 成为分销商
	p.Post("/open", Open)
	// 关闭分销
	p.Post("/close", Close)
	// 检测是否开通分销
	p.Post("/check", Check)
	// 查询所有分销商列表
	p.Post("/query", admin.Admin, Query)
	// 确认分账
	p.Post("/received", admin.Admin, Received)

	p.Post("/order/query", ReceiverQuery)
	p.Post("/order/list", OrderList)
	p.Post("/order/total", Total)
}
