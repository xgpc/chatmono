/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/20 10:57
 */

package wechatOfficial

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"gorm.io/gorm"
)

var (
	_db     *gorm.DB
	account *officialaccount.OfficialAccount
)

type Config struct {
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

func Init(conf Config) {
	wc := wechat.NewWechat()

	memory := cache.NewMemory()

	cfg := &offConfig.Config{
		AppID:     conf.AppID,
		AppSecret: conf.AppSecret,
		Token:     "syd745j9b368T12PhgNxO7e0cSZ54sJQ",
		//EncodingAESKey: "6wdxffqIczmJUOwKyxogujy1jsRs9Kng5hZEzRrfnAg",
		Cache: memory,
	}
	account = wc.GetOfficialAccount(cfg)
}

func Hand() *officialaccount.OfficialAccount {
	return account
}

func GetAcctoken() (accessToken string, err error) {
	return Hand().GetAccessToken()
}

func GetTicket() (string, error) {
	token, err := GetAcctoken()
	if err != nil {
		return "", err
	}
	return Hand().GetJs().GetTicket(token)
}

// AuthorizationCode 验证
func AuthorizationCode(code string) (*oauth.UserInfo, error) {
	oauth := Hand().GetOauth()
	acc, err := oauth.GetUserAccessToken(code)
	if err != nil {
		return nil, err
	}

	info, err := oauth.GetUserInfo(acc.AccessToken, oauth.AppID, "")
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func UserInfo(openID string) (*user.Info, error) {
	user := Hand().GetUser()

	info, err := user.GetUserInfo(openID)
	if err != nil {
		return nil, err
	}

	return info, nil
}
