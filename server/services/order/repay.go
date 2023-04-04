package order

import (
	"chatmono/services/order/wechatPly"
	jsapisvc "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

// Repay 支付
func (p *order) Repay(orderPay *wechatPly.OrderRequest) (*jsapisvc.PrepayWithRequestPaymentResponse, error) {
	return p.wxmini.PayUnifiedOrder(orderPay)
}
