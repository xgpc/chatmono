package product

import (
	"chatmono/services/admin"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	p := api.Party("/product")

	// 添加商品
	p.Post("/add", user.Login, admin.Admin, Add)
	// 修改商品
	p.Post("/up", user.Login, admin.Admin, Up)
	// 删除商品
	p.Post("/del", user.Login, admin.Admin, Del)
	// 商品列表
	p.Post("/query", Query)
	// 商品详情
	p.Post("/info", Info)
}
