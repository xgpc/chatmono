package referrer

import (
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	p := api.Party("/referrer", user.Login)

	p.Post("/commit", Commit)
	p.Post("/check", Check)
}
