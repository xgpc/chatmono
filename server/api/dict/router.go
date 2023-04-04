/**
 * @Author: smono
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/9/25 0:30
 */

package dict

import (
	"chatmono/services/admin"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	r := api.Party("/dict")

	r.Post("/get", Get)
	r.Post("/del", user.Login, admin.Admin, PostDel)
	r.Post("/query", PostQuery)
	r.Post("/list", List)
	r.Post("/created", user.Login, admin.Admin, PostCreated)
	r.Post("/up", user.Login, admin.Admin, Up)
}
