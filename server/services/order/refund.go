/**
 * @Author: smono
 * @Description:
 * @File:  refund
 * @Version: 1.0.0
 * @Date: 2022/8/16 11:34
 */

package order

import (
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

type Channel string

func (e Channel) Ptr() *Channel {
	return &e
}

var ChannelDict = map[refunddomestic.Channel]string{
	refunddomestic.CHANNEL_ORIGINAL:       "原路退款",
	refunddomestic.CHANNEL_BALANCE:        "退回到余额",
	refunddomestic.CHANNEL_OTHER_BALANCE:  "原账户异常退到其他余额账户",
	refunddomestic.CHANNEL_OTHER_BANKCARD: "原银行卡异常退到其他银行卡",
}

func GetChannel(channel *refunddomestic.Channel) string {
	if ChannelDict == nil {
		return "无"
	}
	return ChannelDict[*channel]
}

// Refund 订单退款
func (p *order) Refund(orderSn string, price float64) (*refunddomestic.Refund, error) {
	priceFloat := uint32(price * 100)
	return p.wxmini.Refund(orderSn, orderSn, orderSn, priceFloat, priceFloat)
}
