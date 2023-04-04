/**
 * @Author: smono
 * @Description:
 * @File:  code
 * @Version: 1.0.0
 * @Date: 2022/9/21 14:05
 */

package user

import (
	"chatmono/services/msg"
	"chatmono/services/msg/random"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type CodeReq struct {
	Mobile string `json:"mobile"`
}

// Code 发送验证码
// @Summary      发送验证码
// @Description  发送验证码, 填写手机号,后端将发送验证码到指定手机, 然后调用验证码登录接口 填入对应的验证码登录(5分钟内失效)
// @Accept       json
// @Produce      json
// @param        root  body  CodeReq  true  "参数"
// @Tags         用户管理
// @Success      200  {object}  render.Response
// @Router       /api/user/code [Post]
func Code(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param CodeReq
	p.Init(&param)

	code := random.CreatedMobileCode(param.Mobile)

	err := msg.SendCode(param.Mobile, code)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, (err))
	}
	p.Success()

}
