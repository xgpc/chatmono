// Package order
// @Author:        asus
// @Description:   $
// @File:          order
// @Data:          2022/1/1410:57
//
package order

import (
	"chatmono/services/order/sn"
	"chatmono/services/order/wechatAPI"
	"chatmono/services/order/wechatPly"
	"errors"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
)

type order struct {
	wxmini *wechatPly.JsApiPay
	Api    *wechatAPI.Hand
}

var Handle *order

func Init(conf wechatPly.Config, rdb *redis2.Client) {
	_rdb = rdb

	sn.Init(rdb)
	if Handle == nil {
		// 获取支付环境变量
		wxConfig := conf

		// 创建小程序支付句柄
		wxmini, err := wechatPly.NewJsapi(wxConfig.Mchid, wxConfig.MchCertificateSerialNumber, wxConfig.MchAPIv3Key, wxConfig.MchPrivateKeyPath, wxConfig.NotifyUrl, wxConfig.AppID)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		// 账单分账
		api := wechatAPI.NewHand(wxConfig.AppID, wxConfig.Mchid, wxConfig.MchCertificateSerialNumber, wxConfig.MchAPIv3Key, wxConfig.MchPrivateKeyPath)
		if api == nil {
			fmt.Println("wechatAPI.NewHand Create error")
			panic(errors.New("wechatAPI.NewHand Create error"))
		}

		Handle = &order{
			wxmini: wxmini,
			Api:    api,
		}
	}

}
