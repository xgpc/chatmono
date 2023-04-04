/**
 * @Author: smono
 * @Description:
 * @File:  session
 * @Version: 1.0.0
 * @Date: 2022/9/20 23:51
 */

package wxMini

import (
	"chatmono/models"
	"chatmono/services/aes"
	"chatmono/services/user"
	"chatmono/services/wechat/wechatMini"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type wxMiniLoginReq struct {
	JsCode string `json:"js_code"` // 必填
	AppID  string `json:"app_id"`
}

// PostLogin 微信登录
// @Summary      微信登录
// @Description  微信登录
// @Accept       json
// @Produce      json
// @param        root  body  wxMiniLoginReq  true  "参数"
// @Tags         小程序
// @Success      200  {object}  render.Response
// @Router       /api/wxMini/login [Post]
func PostLogin(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param wxMiniLoginReq
	p.Init(&param)

	res, err := wechatMini.GetSession(param.JsCode)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, "小程序code查询错误: "+err.Error())
	}
	resData := map[string]interface{}{}
	resData["wx_info"] = res

	// 查询是否存在
	var num int64
	err = p.DB().Model(&models.WxUser{}).Limit(1).
		Where(models.WxUserColumns.WxAppid, param.AppID).
		Where(models.WxUserColumns.OpenID, res.OpenID).Count(&num).Error

	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	if num > 0 {
		mgr := models.WxUserMgr(p.DB())
		// 理论一个openID只绑定一个user 为空不一定是错误, 可能是未注册
		wxUser, _ := mgr.GetByOption(
			mgr.WithWxAppid(param.AppID),
			mgr.WithOpenID(res.OpenID))

		var info models.User
		err := p.DB().Model(&info).Where(models.UserColumns.ID, wxUser.UserID).Limit(1).First(&info).Error
		if err != nil {
			exce.ThrowSys(exce.CodeRequestError, err.Error())
		}

		info.Mobile = aes.DeCode(info.MobileData)
		info.MobileData = nil

		resData["user"] = info
		if wxUser.UserID != 0 {
			resData["token"] = user.SetSession(wxUser.UserID)
		}
	}

	p.SuccessWithData(resData)
}
