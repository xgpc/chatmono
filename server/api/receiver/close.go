/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/

package receiver

import (
	"chatmono/models"
	"chatmono/services/order"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type receiverCloseReq struct {
}

// Close 关闭分销
// @Summary      关闭分销
// @Description  关闭分销
// @Accept       json
// @Produce      json
// @param        root  body  receiverCloseReq  true  "参数"
// @Tags         分销管理
// @Success      200  {object}  render.Response
// @Router       /api/receiver/close [Post]
func Close(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverCloseReq
	p.Init(&param)

	var info models.WxUser
	err := p.DB().Model(&info).Where(models.WxUserColumns.UserID, p.MyId()).First(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	rsp, err := order.Handle.Api.DeleteReceiver(info.OpenID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	info.IsProfitsharing = models.IsProfitsharingFalse
	err = p.DB().Model(&info).Where(models.WxUserColumns.ID, info.ID).Save(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(rsp)

}
