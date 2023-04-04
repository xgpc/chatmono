/**
 * @Author: smono
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/10/11 0:02
 */

package order

import (
	"chatmono/api/order/delivery"
	"chatmono/api/order/refund"
	"chatmono/services/admin"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	r := api.Party("/order", user.Login)

	r.Post("/notify/checkup", NotifyCheckUp)
	r.Post("/continue/pay", ContinuePayCheckUp)
	r.Post("/pay", Pay)

	r.Post("/list", List)
	r.Post("/query", Query)
	r.Post("/info", Info)

	//退款
	r.Post("/refund/quit", admin.Admin, refund.Quit)
	r.Post("/refund/apply", refund.RefundApply)

	// 配送
	r.Post("/delivery/up", delivery.DeliveryUp)
}
