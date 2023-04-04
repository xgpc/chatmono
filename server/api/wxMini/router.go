/**
 * @Author: smono
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/9/25 0:30
 */

package wxMini

import (
	"github.com/kataras/iris/v12"
)

func Router(api iris.Party) {
	r := api.Party("/wxMini")

	r.Post("/login", PostLogin)
	r.Post("/decrypt", PostDecrypt)
	r.Post("/getMobile", GetMobile)
	r.Post("/register", PostRegister)
	r.Post("/qrCode", QrCode)
}
