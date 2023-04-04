package refund

import (
	"chatmono/models"
	"chatmono/services/admin"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type orderRefundApplyReq struct {
	ID           uint32 `json:"id"`                                        // 订单ID   -> product_order_id
	RefundReason string `gorm:"column:refund_reason" json:"refund_reason"` // 退款原因
}

// RefundApply 退款申请
// @Summary      退款申请
// @Description  退款申请
// @Accept       json
// @Produce      json
// @param        root  body  orderRefundApplyReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response
// @Router       /api/order/refund/apply [Post]
func RefundApply(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param orderRefundApplyReq
	p.Init(&param)

	mgr := models.ProductOrderMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	if info.PayStatus != models.PayStatusOK {
		exce.ThrowSys(exce.CodeRequestError, "订单未支付, 不能申请退款")
	}

	switch info.RefundStatus {
	case models.RefundStatusQuitOK:
		exce.ThrowSys(exce.CodeRequestError, "订单已退款")
	case models.RefundStatusQuitExamine:
		exce.ThrowSys(exce.CodeRequestError, "订单已申请退款, 请勿重复申请")
	}

	if info.CreatedUserID != p.MyId() {
		if !admin.IsAdmin(p.MyId()) {
			exce.ThrowSys(exce.CodeRequestError, "只能本人申请退款, 或者管理员申请退款")
		}
	}

	info.RefundStatus = models.RefundStatusQuitExamine
	info.RefundReason = param.RefundReason

	err = p.DB().Where(models.ProductOrderColumns.ID, param.ID).Save(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(info)
}
