/**
 * @Author: smono
 * @Description:
 * @File:  status
 * @Version: 1.0.0
 * @Date: 2022/10/11 23:26
 */

package order

const (
	SUCCESS    = "SUCCESS"
	REFUND     = "REFUND"
	NOTPAY     = "NOTPAY"
	CLOSED     = "CLOSED"
	REVOKED    = "REVOKED"
	USERPAYING = "USERPAYING"
	PAYERROR   = "PAYERROR"
)

var TradeState = map[string]string{
	SUCCESS:    "支付成功",
	REFUND:     "转入退款",
	NOTPAY:     "未支付",
	CLOSED:     "已关闭",
	REVOKED:    "已撤销（付款码支付）",
	USERPAYING: "用户支付中（付款码支付）",
	PAYERROR:   "支付失败(其他原因，如银行返回失败)",
}
