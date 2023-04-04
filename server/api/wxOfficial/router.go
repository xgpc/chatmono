package wxOfficial

import "github.com/kataras/iris/v12"

func Router(api iris.Party) {
	p := api.Party("/wxOfficial")

	// 公众号分享
	p.Post("/share", Share)
	p.Post("/login", PostLogin)
	p.Post("/info", Info)
	p.Post("/register", PostRegister)
}
