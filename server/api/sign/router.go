package sign

import (
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	r := api.Party("/sign", user.Login)

	// 人员签到
	r.Post("/", Sign)
	// 查询今日是否签到
	r.Post("/check", Check)
	// 签到信息
	r.Post("/info", Info)
}
