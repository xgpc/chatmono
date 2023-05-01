package alipay

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"testing"
)

func TestDemo(t *testing.T) {

	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient("2021003187693658", privateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).             // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).            // 设置签名类型，不设置默认 RSA2
							SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
							SetNotifyUrl("https://www.fmm.ink")  // 设置异步通知URL
	//SetAppAuthToken() // 设置第三方应用授权

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign([]byte("alipayCertPublicKey_RSA2 bytes"))

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
	// 证书内容
	//err := client.SetCertSnByContent("appCertPublicKey bytes", "alipayRootCert bytes", "alipayCertPublicKey_RSA2 bytes")
	if err != nil {
		t.Fatal(err)
	}
	bm := make(gopay.BodyMap)
	bm.Set("subject", "条码支付").
		Set("scene", "bar_code").
		Set("auth_code", "286248566432274952").
		Set("out_trade_no", "GZ201909081743431443").
		Set("total_amount", "0.01").
		Set("timeout_express", "2m")

	pay, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		return
	}

	fmt.Println(pay)

}
