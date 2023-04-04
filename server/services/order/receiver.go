package order

import (
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"github.com/xgpc/dsg/exce"
)

func (p *order) AddReceiver(openID, name string) *profitsharing.AddReceiverResponse {
	rsp, err := p.Api.AddReceiver(openID, name)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	return rsp
}

func (p *order) DeleteReceiver(openID string) *profitsharing.DeleteReceiverResponse {
	rsp, err := p.Api.DeleteReceiver(openID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	return rsp
}
