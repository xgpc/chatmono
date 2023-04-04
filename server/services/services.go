/**
 * @Author: smono
 * @Description:
 * @File:  services
 * @Version: 1.0.0
 * @Date: 2022/9/13 15:00
 */

package services

import (
	"chatmono/config"
	"chatmono/services/admin"
	"chatmono/services/aes"
	"chatmono/services/fileServer"
	"chatmono/services/openAI"
	"chatmono/services/signServer"
	"chatmono/services/user"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/frame"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, rdb *redis.Client) {
	conf := frame.Config
	localConf := config.Config

	fmt.Printf("conf = %v", conf)
	fmt.Printf("localConf = %v", localConf)

	// 启动富文本
	//ueditor.Init(&localConf.Ueditor, conf.App.TLS)

	// 权限管理
	admin.Init(db)

	// 用户系统
	user.Init(rdb)

	// 加密功能
	aes.Init(localConf.Secret)
	// 微信功能
	//wechatMini.Init(localConf.Mini)
	//wechatOfficial.Init(localConf.Mini)

	// oss
	//oss.Init(localConf.Oss)

	// 文件服务
	fileServer.Init(localConf.FileServer)

	// 短信服务
	//msg.Init(localConf.AliMsg, rdb)

	// 商品订单系统
	//productOrder.Init(localConf.WechatPay, db, rdb)

	// 签到系统
	signServer.Init(db, rdb)

	// openAI
	openAI.Init(db, rdb, localConf.OpenAIConfig)
}
