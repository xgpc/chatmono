/**
 * @Author: smono
 * @Description:
 * @File:  notifyCheckUp
 * @Version: 1.0.0
 * @Date: 2022/10/11 22:53
 */

package order

import (
	"chatmono/models"
	"chatmono/services/productOrder"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type notifyCheckUpReq struct {
	ID        uint32 `json:"id"`                                  // 订单ID
	PayStatus int    `gorm:"column:pay_status" json:"pay_status"` // 支付状态:[2取消, 3支付完成]
}

// NotifyCheckUp 订单回调
// @Summary      订单回调
// @Description  订单回调
// @Accept       json
// @Produce      json
// @param        root  body  notifyCheckUpReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response
// @Router       /api/order/notify/checkup [Post]
func NotifyCheckUp(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param notifyCheckUpReq
	p.Init(&param)

	mgr := models.ProductOrderMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	if info.ID == 0 {
		exce.ThrowSys(exce.CodeRequestError, "未找到数据")
	}

	//支付状态:[1已取消或超时, 2支付完成,]
	switch param.PayStatus {
	case models.PayStatusTimeOutOrCancel:
		productOrder.CancelCheckUP(&info)
	case models.PayStatusOK:
		productOrder.DealCheckUP(&info)

	}
	p.SuccessWithData(info)
}
