/**
 * @Author: smono
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/9/25 0:30
 */

package user

import (
	user2 "chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	r := api.Party("/user")

	r.Post("/query", user2.Login, Query)
	r.Post("/up", user2.Login, PostUp)
	r.Post("/get", user2.Login, Get)
	r.Post("/login", PostLogin)
	r.Post("/code", Code)

	//
	r.Post("/logon", Logon)

}
