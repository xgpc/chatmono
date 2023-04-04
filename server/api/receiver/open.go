package receiver

import (
	"chatmono/models"
	"chatmono/services/order"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type receiverOpenReq struct {
}

// Open 开通分销模式
// @Summary      开通分销模式
// @Description  开通分销模式(需要公众号登录绑定过openID, 分销金额打进用户钱包)
// @Accept       json
// @Produce      json
// @param        root  body  receiverOpenReq  true  "参数"
// @Tags         分销管理
// @Success      200  {object}  render.Response
// @Router       /api/receiver/open [Post]
func Open(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverOpenReq
	p.Init(&param)
	var info models.WxUser
	err := p.DB().Model(&info).Where(models.WxUserColumns.UserID, p.MyId()).First(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	rsp, err := order.Handle.Api.AddReceiver(info.OpenID, "")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	info.IsProfitsharing = models.IsProfitsharingTrue
	err = p.DB().Model(&info).Where(models.WxUserColumns.ID, info.ID).Save(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(rsp)

}
