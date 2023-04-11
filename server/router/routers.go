package router

import (
	"chatmono/api/dict"
	"chatmono/api/openAI"
	"chatmono/api/order"
	"chatmono/api/product"
	"chatmono/api/session"
	"chatmono/api/user"
	_ "chatmono/docs"
	"chatmono/services/fileServer"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/middleware/recover"
	"github.com/xgpc/dsg/middleware"
	"github.com/xgpc/dsg/version"

	"github.com/rs/cors"
)

func Routers(api *iris.Application) {
	//    跨域
	c := cors.AllowAll()
	api.WrapRouter(c.ServeHTTP)
	//    中间件
	api.Use(middleware.ExceptionLog)

	//    版本
	api.Get("/", version.Version)

	api.Get("/swagger/{any:path}", swagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	// 文件服务
	fileServer.Router(api)

	a := api.Party("/api")
	//oss.Router(a)
	dict.Router(a)
	user.Router(a)

	//wxMini.Router(a)
	//wxOfficial.Router(a)

	product.Router(a)
	order.Router(a)
	//sign.Router(a)
	//score.Router(a)
	//referrer.Router(a)
	//receiver.Router(a)
	openAI.Router(a)
	session.Router(a)
}
