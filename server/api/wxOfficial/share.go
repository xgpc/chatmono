package wxOfficial

import (
	"chatmono/services/wechat/wechatOfficial"
	"crypto/sha1"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/util"
	"io"
	"strconv"
	"time"
)

type wechatShareReq struct {
	Url   string `json:"url"`
	AppId string `json:"appId"`
}

type wechatShareRsp struct {
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
	AppID     string `json:"appId"`
}

// Share         公众号分享
// @Summary      公众号分享
// @Description  公众号分享
// @Accept       json
// @Produce      json
// @param        root  body  wechatShareReq  true  "参数"
// @Tags         公众号
// @Success      200  {object}  render.Response
// @Router       /api/wxOfficial/share [Post]
func Share(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param wechatShareReq
	p.Init(&param)

	ticket, err := wechatOfficial.GetTicket()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	noncestr := util.RandomStr(16)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signatureStr := "jsapi_ticket=" + ticket + "&noncestr=" + noncestr + "&timestamp=" + timestamp + "&url=" + param.Url

	t := sha1.New()
	io.WriteString(t, signatureStr)
	signature := fmt.Sprintf("%x", t.Sum(nil))

	data := wechatShareRsp{
		Noncestr:  noncestr,
		Timestamp: timestamp,
		Url:       param.Url,
		Signature: signature,
		AppID:     param.AppId,
	}

	p.SuccessWithData(data)
}
