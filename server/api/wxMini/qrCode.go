/**
 * @Author: smono
 * @Description:
 * @File:  qrCode
 * @Version: 1.0.0
 * @Date: 2022/9/25 1:43
 */

package wxMini

import (
	"chatmono/services/wechat/wechatMini"
	"github.com/kataras/iris/v12"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"github.com/xgpc/dsg/frame"
)

type QrCodeReq struct {
	qrcode.QRCoder
}

// QrCode 生成二维码
// @Summary      生成二维码
// @Description  生成二维码
// @Accept       json
// @Produce      json
// @param        root  body  QrCodeReq  true  "参数"
// @Tags         小程序
// @Success      200  {object}  render.Response
// @Router       /api/wxMini/qrCode [Post]
func QrCode(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param QrCodeReq
	p.Init(&param)

	res := map[string]interface{}{}
	res["qrcode"] = wechatMini.CreatedQrCode(param.QRCoder)

	p.SuccessWithData(res)
}
