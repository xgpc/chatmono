package order

import "chatmono/services/order/wechatPly"

// Query 订单查询
func (p *order) Query(orderSn string) (*wechatPly.PaySuccess, error) {
	return p.wxmini.QueryOrder(orderSn)
}
