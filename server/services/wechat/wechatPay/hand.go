/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/20 10:57
 */

package wechatPay

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/pay"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
	"gorm.io/gorm"
)

var (
	_db     *gorm.DB
	payHand *pay.Pay
)

type Config struct {
	AppID     string `json:"app_id"`
	MchID     string `json:"mch_id"`
	Key       string `json:"key"`
	NotifyURL string `json:"notify_url"`
}

func Init(conf Config) {
	wc := wechat.NewWechat()
	//cfg := &offConfig.Config{
	//    AppID:     conf.AppID,
	//    AppSecret: conf.AppSecret,
	//    Token:     "syd745j9b368T12PhgNxO7e0cSZ54sJQ",
	//    //EncodingAESKey: "6wdxffqIczmJUOwKyxogujy1jsRs9Kng5hZEzRrfnAg",
	//    Cache: memory,
	//}

	cfg := &payConfig.Config{
		//AppID     string `json:"app_id"`
		AppID: conf.AppID,
		//MchID     string `json:"mch_id"`
		MchID: conf.MchID,
		//Key       string `json:"key"`
		Key: conf.Key,
		//NotifyURL string `json:"notify_url"`
		NotifyURL: conf.NotifyURL,
	}
	payHand = wc.GetPay(cfg)
}

func Hand() *pay.Pay {
	return payHand
}

func Query() {
}
