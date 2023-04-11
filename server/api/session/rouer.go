package session

import (
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	a := api.Party("/session", user.Login)

	a.Post("/all", All)
	a.Post("/set", Set)
	a.Post("/get", Get)
	a.Post("/del", Del)

}
