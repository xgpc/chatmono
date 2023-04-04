package order

func (p *order) CancelOrder(orderSn string) (bool, error) {
	return p.wxmini.CancelOrder(orderSn)
}
