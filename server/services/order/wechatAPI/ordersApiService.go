package wechatAPI

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"log"
)

const Ledger = `LEDGER`

func (h *Hand) OrdersApiService() *profitsharing.OrdersApiService {
	svc := profitsharing.OrdersApiService{Client: h.client}
	return &svc
}

func (h *Hand) CreateOrder(orderSn, TransactionId, Desc, openID string, Amount float64) (*profitsharing.OrdersEntity, error) {
	AmountInt64 := int64(Amount * 100)
	resp, result, err := h.OrdersApiService().CreateOrder(
		context.Background(),

		profitsharing.CreateOrderRequest{
			Appid:      core.String(h.appID),
			OutOrderNo: core.String(Ledger + orderSn),
			Receivers: []profitsharing.CreateOrderReceiver{profitsharing.CreateOrderReceiver{
				Account:     core.String(openID),
				Amount:      core.Int64(AmountInt64),
				Description: core.String(Desc),
				Type:        core.String("PERSONAL_OPENID"),
			}},

			TransactionId:   core.String(TransactionId),
			UnfreezeUnsplit: core.Bool(true)})
	if err != nil {
		// 处理错误
		log.Printf("call CreateOrder err:%s", err)
		return nil, err
	}
	//处理返回结果
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}

func (h *Hand) QueryOrder(orderSn string) (*profitsharing.OrdersEntity, error) {

	resp, result, err := h.OrdersApiService().QueryOrder(
		context.Background(),
		profitsharing.QueryOrderRequest{
			OutOrderNo: core.String(Ledger + orderSn)},
	)
	if err != nil {
		// 处理错误
		log.Printf("call CreateOrder err:%s", err)
		return nil, err
	}
	//处理返回结果
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	return resp, nil
}
