package order

import "github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"

// RefundQuery 订单退款查询
func (p *order) RefundQuery(orderSn string) (*refunddomestic.Refund, error) {
	return p.wxmini.RefundQuery(orderSn)
}
