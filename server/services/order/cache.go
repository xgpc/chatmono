package order

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/util"
	"time"
)

// ContinuePay 继续支付订单
func (p *order) ContinuePay(orderSn string) *ResOrderInfo {
	return p.cachePayGet(orderSn)
}

func (p *order) cachePaySet(info *ResOrderInfo) {
	data, err := util.JsonEncode(info)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	_, err = redis().Set(context.Background(), "order:"+info.OrderSn, data, time.Minute*6).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
}

func (p *order) cachePayGet(OrderSn string) *ResOrderInfo {

	data, _ := redis().Get(context.Background(), "order:"+OrderSn).Result()

	if data == "" {
		exce.ThrowSys(exce.CodeRequestError, "订单已超时或不存在")
	}
	var res ResOrderInfo
	err := util.JsonDecode([]byte(data), &res)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	return &res
}
