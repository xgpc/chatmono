/**
 * @Author: smono
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/9/27 16:35
 */

package oss

import (
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xgpc/dsg/frame"
)

func Router(api iris.Party) {

	mvc.Configure(api, func(m *mvc.Application) {
		m.Register(func(ctx iris.Context) *frame.Base {
			return frame.NewBase(ctx)
		})

		m.Party("/server", user.Login).Handle(new(UploadController)) // 文件上传
	})
}
