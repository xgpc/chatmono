package wxOfficial

import (
	"chatmono/services/wechat/wechatOfficial"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type wxOfficialInfoReq struct {
	OpenID string `json:"open_id"`
}

// Info 获取用户信息
// @Summary      获取用户信息
// @Description  获取用户信息
// @Accept       json
// @Produce      json
// @param        root  body  wxOfficialInfoReq  true  "参数"
// @Tags         公众号
// @Success      200  {object}  render.Response
// @Router       /api/wxOfficial/info [Post]
func Info(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param wxOfficialInfoReq
	p.Init(&param)

	info, err := wechatOfficial.UserInfo(param.OpenID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(info)

}
