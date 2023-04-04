package receiver

import (
	"chatmono/models"
	"chatmono/services/order"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type receiverReceivedReq struct {
	ID uint32 `json:"id"` // 订单ID
}

// Received 确认分账
// @Summary      确认分账
// @Description  确认分账
// @Accept       json
// @Produce      json
// @param        root  body  receiverReceivedReq  true  "参数"
// @Tags         分销管理
// @Success      200  {object}  render.Response
// @Router       /api/receiver/received [Post]
func Received(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverReceivedReq
	p.Init(&param)
	var info models.ProductOrder
	var rsp interface{}

	// 获取订单详情
	err := p.DB().Model(&info).Where(models.ProductOrderColumns.ID, param.ID).Preload("Product").First(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 分销金额
	if info.ReferrerID == 0 {
		exce.ThrowSys(exce.CodeRequestError, "推荐人ID为空, 该订单无法分账")
	}
	// 获取分账人 信息
	wxMgr := models.WxUserMgr(p.DB())
	wxUser, err := wxMgr.GetFromUserID(info.ReferrerID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, "未找到分销用户信息, 分销失败")
	}
	if wxUser.IsProfitsharing != models.IsProfitsharingTrue {
		exce.ThrowSys(exce.CodeRequestError, "用户未开通分销模式, 分销失败")
	}

	rsp, err = order.Handle.Api.CreateOrder(info.OrderSn, info.TransactionID, "分销金", wxUser.OpenID, info.ProfitsharingMoney)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 修改状态
	info.ProfitsharingStatus = models.ProfitsharingStatusTrue
	err = p.DB().Model(&info).Where(models.ProductOrderColumns.ID, param.ID).
		Save(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(rsp)
}
