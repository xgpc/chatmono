package refund

import (
	"chatmono/models"
	"chatmono/services/order"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/log"
)

type orderQuitReq struct {
	ID           uint32 `json:"id"`                                        // 订单ID   -> product_order_id
	RefundStatus int    `gorm:"column:refund_status" json:"refund_status"` // 退款状态:[3驳回退款, 4同意退款]
	RefundDetail string `gorm:"column:refund_detail" json:"refund_detail"` // 退款详情: 驳回理由, 同意退款自动生成
}

// Quit 订单退款
// @Summary      订单退款
// @Description  订单退款
// @Accept       json
// @Produce      json
// @param        root  body  orderQuitReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response
// @Router       /api/order/refund/quit [Post]
func Quit(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param orderQuitReq
	p.Init(&param)
	mgr := models.ProductOrderMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	if info.PayStatus != models.PayStatusOK {
		exce.ThrowSys(exce.CodeRequestError, "订单未支付, 不能退款")
	}

	if info.RefundStatus != models.RefundStatusQuitExamine {
		exce.ThrowSys(exce.CodeRequestError, "该订单未申请退款")
	}

	switch param.RefundStatus {
	case models.RefundStatusQuitOK:
		// 微信退款
		res, err := order.Handle.Refund(info.OrderSn, info.Price)
		if err != nil {
			exce.ThrowSys(exce.CodeRequestError, err.Error())
		}
		info.PayStatus = models.PayStatusQuit
		info.RefundStatus = param.RefundStatus
		info.RefundDetail = order.GetChannel(res.Channel)
		log.Error(res)
	case models.RefundStatusQuitRe:
		info.RefundStatus = param.RefundStatus
		info.RefundDetail = param.RefundDetail
	default:
		exce.ThrowSys(exce.CodeRequestError, "未知状态")
	}

	err = p.DB().Model(&info).Where(models.ProductOrderColumns.ID, param.ID).Save(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(info)
}
