/**
 * @Author: smono
 * @Description:
 * @File:  login
 * @Version: 1.0.0
 * @Date: 2022/9/21 13:34
 */

package user

import (
	"chatmono/models"
	"chatmono/services/aes"
	"chatmono/services/msg/random"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type userLoginReq struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

// PostLogin 登录
// @Summary      登录
// @Description  登录, 通过验证码登录
// @Accept       json
// @Produce      json
// @param        root  body  userLoginReq  true  "参数"
// @Tags         用户管理
// @Success      200  {object}  render.Response
// @Router       /api/user/login [Post]
func PostLogin(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param userLoginReq
	p.Init(&param)

	if !random.CheckMobileCode(param.Mobile, param.Code) {
		exce.ThrowSys(exce.CodeRequestError, "验证码不通过")
	}

	var userInfo models.User

	err := p.DB().Model(&userInfo).
		Where(models.UserColumns.MobileData, aes.EnCode(param.Mobile)).
		Attrs(models.User{
			UserName:   "手机用户",
			MobileData: aes.EnCode(param.Mobile),
			Mobile:     aes.EnMobile(param.Mobile),
		}).
		FirstOrCreate(&userInfo).Error
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	token := user.SetSession(userInfo.ID)

	userInfo.Mobile = param.Mobile
	userInfo.MobileData = nil

	res := map[string]interface{}{}
	res["token"] = token
	res["user"] = userInfo

	p.SuccessWithData(res)

}
