// Package orderPay
// @Author:        asus
// @Description:   $
// @File:          pay
// @Data:          2021/12/2320:12
//
package wechatPly

import (
	"testing"
)

var (
	mchid                      = "1614508080"
	mchKey                     = "SkLRqvd6pxH0ioVLjL4tSJQwLuUIOc0N"
	mchCertificateSerialNumber = "776BC1AF8E992829104094D6459A5F83696F237B"
	mchAPIv3Key                = "8edimXegiIVyytpMHz3NKoSsbON4sgMV"
	mchPrivateKeyPath          = "cert/apiclient_key.pem"
	appid                      = "wx236d7a807086d1d2"
	jsapiNotifyUrl             = "https://cs.dousougou.com:19512"
)

func TestPay(t *testing.T) {

	// 获取jsapi支付客户端
	jsapi, err := NewJsapi(mchid, mchCertificateSerialNumber, mchAPIv3Key, mchPrivateKeyPath, jsapiNotifyUrl, appid)
	if err != nil {
		t.Error(err.Error())
		return
	}

	// 测试订单退款
	//str, err := jsapi.Refund("CYGH47147661cac37128", "CYGH47147661cac37128TD", "卖家申请退款", 1, 1)
	//if err != nil {
	//	fmt.Printf("err:%+v\n\n", err)
	//	t.Error(err.Error())
	//	return
	//}
	//t.Log(str)

	// 测试退款订单查询
	//str, err = jsapi.RefundQuery("CYGH47147661cac37128TD")
	//if err != nil {
	//	t.Error(err.Error())
	//	return
	//}
	//
	//t.Log(str)

	//
	orderRequest := &OrderRequest{
		Price:   1,
		Desc:    "测试商品",
		OpenId:  "oCTIB59ibNmj19OJMaQnXKzYkugs",
		OrderSn: "testXCGH14115661c67ad72069",
	}
	//
	//调起支付
	res, err := jsapi.PayUnifiedOrder(orderRequest)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)

	//// 获取jsapi支付客户端
	//native, err := NewNative(mchid, mchCertificateSerialNumber, mchAPIv3Key, mchPrivateKeyPath, jsapiNotifyUrl, appid)
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//
	////调起支付
	//res, err = native.PayUnifiedOrder(orderRequest)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//t.Log(res)
	//
	//mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyPath)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//ctx := context.Background()
	//// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	//opts := []core.ClientOption{
	//	option.WithWechatPayAutoAuthCipher(mchid, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	//}
	//
	//client, err := core.NewClient(ctx, opts...)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//svc := app.AppApiService{Client: client}
	//resp, result, err := svc.Prepay(ctx,
	//	app.PrepayRequest{
	//		Appid:       core.String(appid),
	//		Mchid:       core.String(mchid),
	//		Description: core.String("Image形象店-深圳腾大-QQ公仔"),
	//		OutTradeNo:  core.String(orderRequest.OrderSn),
	//		Attach:      core.String("自定义数据说明"),
	//		NotifyUrl:   core.String("https://www.weixin.qq.com/wxpay/pay.php"),
	//		Amount: &app.Amount{
	//			Total: orderRequest.GetInt64Price(),
	//		},
	//	},
	//)
	//
	//if err != nil {
	//	// 处理错误
	//	t.Logf("call Prepay err:%s", err)
	//} else {
	//	// 处理返回结果
	//	t.Logf("status=%d resp=%s", result.Response.StatusCode, resp)
	//}

}

func TestCheck(t *testing.T) {
	// 获取jsapi支付客户端
	jsapi, err := NewJsapi(mchid, mchCertificateSerialNumber, mchAPIv3Key, mchPrivateKeyPath, jsapiNotifyUrl, appid)
	if err != nil {
		t.Error(err.Error())
		return
	}

	//
	//调起支付
	res, err := jsapi.QueryOrder("water16432559985")
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}
