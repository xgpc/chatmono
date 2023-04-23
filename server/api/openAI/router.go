package openAI

import (
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	a := api.Party("/openAI", user.Login)

	a.Post("/info", Info)
	a.Post("/templates/list", Templates)
	a.Post("/list", List)
	a.Post("/send", Send)
	a.Post("/new", New)

}
