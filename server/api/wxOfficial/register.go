/**
 * @Author: smono
 * @Description:
 * @File:  register
 * @Version: 1.0.0
 * @Date: 2022/9/21 11:32
 */

package wxOfficial

import (
	"chatmono/config"
	"chatmono/models"
	"chatmono/services/aes"
	"chatmono/services/msg/random"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type wxMiniRegisterReq struct {
	Mobile string `json:"mobile" validate:"required"`
	OpenID string `json:"open_id" validate:"required"`
	Code   string `json:"code" validate:"required"`
}

// PostRegister 公众号注册账号
// @Summary      公众号注册账号
// @Description  公众号注册账号
// @Accept       json
// @Produce      json
// @param        root  body  wxMiniRegisterReq  true  "参数"
// @Tags          公众号
// @Success      200  {object}  render.Response
// @Router       /api/wxOfficial/register [Post]
func PostRegister(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param wxMiniRegisterReq
	p.Init(&param)

	// 效验手机号码
	if !random.CheckMobileCode(param.Mobile, param.Code) {
		exce.ThrowSys(exce.CodeRequestError, "验证码不通过")
	}

	// 获取user
	moblieData := aes.EnCode(param.Mobile)
	var info models.User

	err := p.DB().Model(&info).
		Where(models.UserColumns.MobileData, moblieData).
		Attrs(models.User{
			UserName:   "微信用户",
			MobileData: moblieData,
			Mobile:     aes.EnMobile(param.Mobile),
		}).
		FirstOrCreate(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	var wxUser models.WxUser
	err = p.DB().Where(models.WxUserColumns.UserID, info.ID).
		Where(models.WxUserColumns.OpenID, param.OpenID).
		Attrs(models.WxUser{
			UserID:  info.ID,
			OpenID:  param.OpenID,
			WxAppid: config.Config.Mini.AppID,
		}).
		FirstOrCreate(&wxUser).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	token := user.SetSession(info.ID)

	p.SuccessWithData(map[string]interface{}{
		"user":  info,
		"token": token,
	})
}
