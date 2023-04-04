package config

import (
	"chatmono/services/fileServer"
	"chatmono/services/msg"
	"chatmono/services/openAI"
	"chatmono/services/order/wechatPly"
	"chatmono/services/oss"
	"chatmono/services/ueditor"
	"chatmono/services/wechat/wechatOfficial"
)

var Config config

type config struct {
	Secret           string                `yaml:"secret"`
	Ueditor          ueditor.Configs       `yaml:"ueditor"`
	Mini             wechatOfficial.Config `yaml:"wechatOfficial"`
	Oss              oss.Config            `yaml:"oss"`
	FileServer       fileServer.Config     `yaml:"file_server"`
	AliMsg           msg.Config            `yaml:"ali_msg"`
	WechatPay        wechatPly.Config      `yaml:"wechatPay"`
	ReferrerScoreNum int                   `yaml:"referrer_score_num"`
	OpenAIConfig     openAI.Config         `yaml:"openAIConfig"`
}
