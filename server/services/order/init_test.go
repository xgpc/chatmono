/**
 * @Author: smono
 * @Description:
 * @File:  order_test.go
 * @Version: 1.0.0
 * @Date: 2022/8/16 11:46
 */

package order

import (
	"chatmono/services/order/wechatPly"
	"fmt"

	"testing"
)

func Start() {

	appid := "wx19c3e0e106b4545c"
	mchid := "1628218523"
	//mchkey:=  "SkLRqvd6pxH0ioVLjL4tSJQwLuUIOc0N"
	//notifyurl := "https://insurance.dousougou.com:16520" // 小程序没有回调
	mchcertificateserialnumber := "40FF1779BED5B7D745A24646CD60F43E21D7613B"
	mchapiv3key := "zljabxjjyxgskmfgsZLJABXJJKMF1028"
	mchprivatekeypath := "../../cert/apiclient_key.pem"

	// 创建小程序支付句柄
	wxmini, err := wechatPly.NewJsapi(
		mchid,
		mchcertificateserialnumber,
		mchapiv3key,
		mchprivatekeypath,
		"",
		appid)

	if err != nil {
		panic(err)
	}
	Handle.wxmini = wxmini
}

func Test_order_Check(t *testing.T) {
	Start()

	orderSn := "INS16605618671"
	data, err := Handle.Query(orderSn)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(data)

	res, err := Handle.wxmini.Refund(orderSn, orderSn, orderSn, 0.01*100, 0.01*100)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(res)

	res2, err := Handle.wxmini.RefundQuery(orderSn)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(res2)
}

func FmtSn(orderSn string) {
	res, err := Handle.Query(orderSn)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(*res.TradeState, "---", orderSn, "---", *res.TransactionId)
}

func TestQueryOrderSn(t *testing.T) {
	Start()

	FmtSn("INS166047716229")
	FmtSn("INS166047905014")

}
