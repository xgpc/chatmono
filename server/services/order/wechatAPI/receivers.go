package wechatAPI

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"log"
)

func (h *Hand) Receivers() *profitsharing.ReceiversApiService {
	svc := profitsharing.ReceiversApiService{Client: h.client}
	return &svc
}

const (
	ADDRECEIVER_SYSTEM_ERROR     = 500 //系统错误	系统异常，请使用相同参数稍后重新调用
	ADDRECEIVER_PARAM_ERROR      = 400 //请求参数不符合参数格式	请使用正确的参数重新调用
	ADDRECEIVER_INVALID_REQUEST  = 400 //无效请求	请确认分账接收方是否存在
	ADDRECEIVER_NO_AUTH          = 403 //商户无权限	请开通商户号分账权限
	ADDRECEIVER_RATELIMIT_EXCEED = 429 //添加接收方频率过高	请降低频率后重试
)

//var AddReceiver  = map

// AddReceiver 申请合作
func (h *Hand) AddReceiver(openID, name string) (*profitsharing.AddReceiverResponse, error) {
	resp, result, err := h.Receivers().AddReceiver(context.Background(), profitsharing.AddReceiverRequest{
		Account:        core.String(openID),
		Appid:          core.String(h.appID),
		CustomRelation: core.String("分销商户"),
		//Name:           core.String(name),
		RelationType: profitsharing.RECEIVERRELATIONTYPE_DISTRIBUTOR.Ptr(),
		Type:         profitsharing.RECEIVERTYPE_PERSONAL_OPENID.Ptr(),
	})

	if err != nil {
		// 处理错误
		log.Printf("call AddReceiver err:%s", err)
		return nil, err
	}

	// 处理返回结果
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil

}

// DeleteReceiver 解绑合作
func (h *Hand) DeleteReceiver(openID string) (*profitsharing.DeleteReceiverResponse, error) {
	resp, result, err := h.Receivers().DeleteReceiver(context.Background(), profitsharing.DeleteReceiverRequest{
		Account: core.String(openID),
		Appid:   core.String(h.appID),
		Type:    profitsharing.RECEIVERTYPE_PERSONAL_OPENID.Ptr(),
	})
	if err != nil {
		// 处理错误
		log.Printf("call DeleteReceiver err:%s", err)
		return nil, err
	}
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}
