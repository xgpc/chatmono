/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/20 10:57
 */

package wechatMini

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"github.com/xgpc/dsg/exce"
	"gorm.io/gorm"
)

var (
	_db         *gorm.DB
	mini        *miniprogram.MiniProgram
	handlerList map[string]interface{}
)

type Config struct {
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

func Init(conf Config) {
	wc := wechat.NewWechat()

	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     conf.AppID,
		AppSecret: conf.AppSecret,
		Cache:     memory,
	}
	mini = wc.GetMiniProgram(cfg)
}

func Hand() *miniprogram.MiniProgram {
	return mini
}

func GetAcctoken() (accessToken string, err error) {
	return Hand().GetAuth().GetAccessToken()
}

func GetUserInfo() (accessToken string, err error) {
	return
}

func GetSession(jsCode string) (auth.ResCode2Session, error) {
	return Hand().GetAuth().Code2Session(jsCode)
}

func Decrypt(sessionKey, encryptedData, iv string) (*encryptor.PlainData, error) {
	return Hand().GetEncryptor().Decrypt(sessionKey, encryptedData, iv)
}

func GetPhoneNumber(code string) (*auth.GetPhoneNumberResponse, error) {
	return Hand().GetAuth().GetPhoneNumber(code)
}

func CreatedQrCode(data qrcode.QRCoder) []byte {
	unlimit, err := Hand().GetQRCode().GetWXACodeUnlimit(data)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	return unlimit
}
