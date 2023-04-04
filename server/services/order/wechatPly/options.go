// Package orderPay
// @Author:        asus
// @Description:   使用选项卡模式来设置对应支付体系的配置
// @File:          pay
// @Data:          2021/12/2320:03
package wechatPly

import (
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"time"
)

// OrderRequest 下单时需要的订单字段
type OrderRequest struct {
	Price         uint32 // 订单价格 分
	Desc          string // 订单描述
	Attach        string // 追加描述
	OpenId        string // 用户openId
	OrderSn       string // 订单编号
	TimeExpire    string // 订单过期时间 半小时过期: 30m
	ProfitSharing bool   // 是否分账
}

func (o *OrderRequest) GetInt64Price() *int64 {
	return core.Int64(int64(o.Price))
}

func (o *OrderRequest) GetTimeExpire() *time.Time {
	if o.TimeExpire != "" {
		expire, _ := time.ParseDuration(o.TimeExpire)
		return core.Time(time.Now().Add(expire))
	}
	return nil
}

// PaySuccess 订单返回时使用的字段
type PaySuccess struct {
	*payments.Transaction
}
