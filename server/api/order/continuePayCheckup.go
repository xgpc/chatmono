/**
 * @Author: smono
 * @Description:
 * @File:  continuePayCheckup
 * @Version: 1.0.0
 * @Date: 2022/10/12 9:47
 */

package order

import (
	"chatmono/models"
	"chatmono/services/order"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type ContinuePayCheckUpReq struct {
	ID uint32 `json:"id"` // 订单ID
}
type ContinuePayCheckUpRsp struct {
	ProductOrder *models.ProductOrder `json:"product_order"`
	WxPay        *order.ResOrderInfo  `json:"wx_pay"`
}

// ContinuePayCheckUp 订单继续支付
// @Summary           订单继续支付
// @Description       订单继续支付
// @Accept       json
// @Produce      json
// @param        root  body  ContinuePayCheckUpReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response{}
// @Router       /api/order/continue/pay [Post]
func ContinuePayCheckUp(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param ContinuePayCheckUpReq
	p.Init(&param)

	mgr := models.ProductOrderMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	if info.ID == 0 {
		exce.ThrowSys(exce.CodeRequestError, "未找到数据")
	}

	if info.CreatedUserID != p.MyId() {
		exce.ThrowSys(exce.CodeRequestError, "非本人下单, 请重试")
	}
	res := ContinuePayCheckUpRsp{}

	//支付状态:[1未支付, 2已取消或超时, 3支付完成, 4退款]
	switch info.PayStatus {
	case models.PayStatusWait:
		res.WxPay = order.Handle.ContinuePay(info.OrderSn)

	default:
		exce.ThrowSys(exce.CodeRequestError, "套餐已超时或支付完成,请刷新")

	}
	res.ProductOrder = &info
	p.SuccessWithData(res)
}
