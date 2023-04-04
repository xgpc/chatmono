/**
 * @Author: smono
 * @Description:
 * @File:  getMobile
 * @Version: 1.0.0
 * @Date: 2022/9/30 10:39
 */

package wxMini

import (
	"chatmono/services/wechat/wechatMini"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type getPhoneReq struct {
	Code string `json:"code"`
}

// GetMobile      通过code获取手机号
// @Summary      通过code获取手机号
// @Description  通过code获取手机号
// @Accept       json
// @Produce      json
// @param        root  body  getPhoneReq  true  "参数"
// @Tags         小程序
// @Success      200  {object}  render.Response
// @Router       /api/wxMini/getMobile [Post]
func GetMobile(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param getPhoneReq
	p.Init(&param)

	res, err := wechatMini.GetPhoneNumber(param.Code)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(res)

}
