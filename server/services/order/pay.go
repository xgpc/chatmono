// Package order
// @Author:        asus
// @Description:   $
// @File:          pay
// @Data:          2022/1/1921:53
//
package order

import (
	"chatmono/services/order/sn"
	"chatmono/services/order/wechatPly"
	jsapisvc "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/util"
)

type ResOrderInfo struct {
	OrderSn string
	WxJsp   *jsapisvc.PrepayWithRequestPaymentResponse
}

func createdRemark(user uint32, price float64, desc string) string {

	md := map[string]interface{}{}
	md["user"] = user
	md["price"] = price
	md["desc"] = desc

	data, err := util.JsonEncode(md)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	return string(data)
}

func (p *order) CreatedOrderSN() string {
	return sn.GetSn()
}

// Pay 下单
func (p *order) Pay(userID uint32, price float64, openid string, desc string, ProfitSharing bool) *ResOrderInfo {

	// 生成orderSn
	OrderSn := p.CreatedOrderSN()

	remark := createdRemark(userID, price, desc)

	// 调起支付
	req := wechatPly.OrderRequest{
		Price:         uint32(price * 100),
		Desc:          remark,
		Attach:        "",
		OpenId:        openid,
		OrderSn:       OrderSn,
		TimeExpire:    "5m",
		ProfitSharing: ProfitSharing,
	}

	res, err := p.Repay(&req)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	resOrder := ResOrderInfo{
		OrderSn: OrderSn,
		WxJsp:   res,
	}

	// 存储数据 防止再次支付
	p.cachePaySet(&resOrder)

	return &resOrder
}
