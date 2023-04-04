/**
 * @Author: smono
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/9/27 16:35
 */

package fileServer

import (
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
	"path/filepath"
)

func Router(api iris.Party) {

	// 访问地址
	api.HandleDir(Hand.UploadDir, filepath.Join(Hand.UploadRoot, Hand.UploadDir))

	r := api.Party("/fileServer", user.Login)

	r.Post("/query", Query)
	r.Post("/upload", upload)
	r.Post("/del", Del)
	r.Post("/move", Move)
	r.Post("/exist", Exist)

}
