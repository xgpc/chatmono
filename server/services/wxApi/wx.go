package wxApi

import (
	"bytes"
	"chatmono/services/wxApi/baseCache"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/xgpc/util"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type wx struct {
	baseCache.Cache
	appID  string
	secret string
}

var Handel wx

func (this *wx) InitTag(db *gorm.DB, conn *redis.Client, tag, appID, secret string) {
	this.Cache.InitTag(db, conn, tag)
	this.appID = appID
}

type RestoKen struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

func (this *wx) getToken() string {
	url := "http://user.dousougou.com/api/wechat/access/token"
	method := "POST"

	payload := strings.NewReader(`{"AppId":"wx236d7a807086d1d2"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	tokenData := RestoKen{}
	err = util.JsonDecode(body, &tokenData)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return tokenData.Data
}

type Template struct {
	Touser     string `json:"touser"`
	TemplateId string `json:"template_id"`
	Url        string `json:"url,omitempty"`
	Data       map[string]struct {
		Value string `json:"value"`
		Color string `json:"color,omitempty"`
	} `json:"data"`
}

// 发送微信模板消息
func (this *wx) sendTemplate(template Template) error {
	requestBody, err := json.Marshal(template)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/template/send", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	params := request.URL.Query()
	params.Add("access_token", this.getToken())
	request.URL.RawQuery = params.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	//body, err := ioutil.ReadAll(response.Body)
	return nil
}

// WaterNotice 水站待接单通知
func (this *wx) WaterNotice(openIDList []string,
	product, num, orderSn,
	userName, Mobile,
	orderStatus, orderTime, orderPrice, addr, Url string) error {

	//{{first.DATA}}
	//客户姓名：{{keyword1.DATA}}
	//客户电话：{{keyword2.DATA}}
	//订单状态：{{keyword3.DATA}}
	//下单时间：{{keyword4.DATA}}
	//订单总价：{{keyword5.DATA}}
	//{{remark.DATA}}

	//商家接单通知
	//客户姓名：张三
	//客户电话：13388990099
	//订单状态：已付款
	//下单时间：2020年11月9日 12:23:47
	//订单总价：12
	//请及时接单

	if len(openIDList) == 0 {
		return nil
	}

	data := Template{
		TemplateId: "wIYzx-r07BoGrJFB6zBPfy6uuL4jyjimQsLTyrQujw0",
		Url:        Url,
		Data: map[string]struct {
			Value string `json:"value"`
			Color string `json:"color,omitempty"`
		}{
			"first": {
				Value: fmt.Sprintf("商品:%s 数量:%s 订单编号:%s", product, num, orderSn),
				Color: "#173177",
			},
			"keyword1": {
				Value: userName,
				Color: "#173177",
			},
			"keyword2": {
				Value: Mobile,
				Color: "#173177",
			},
			"keyword3": {
				Value: orderStatus,
				Color: "#173177",
			},
			"keyword4": {
				Value: orderTime,
				Color: "#173177",
			},
			"keyword5": {
				Value: orderPrice,
				Color: "#173177",
			},
			"remark": {
				Value: "地址:" + addr,
				Color: "#173177",
			},
		},
	}

	for _, v := range openIDList {
		data.Touser = v
		Handel.sendTemplate(data)
	}

	return nil
}

// 接单通知(配送水后, 通知用户)

// UserNotice 水站待接单通知
func (this *wx) UserNotice(openID,
	userName, Mobile, orderTime, remark, Url string) error {

	//{{first.DATA}}
	//姓名：{{keyword1.DATA}}
	//电话：{{keyword2.DATA}}
	//时间：{{keyword3.DATA}}
	//{{remark.DATA}}

	//你好，售后人员已接单，稍后会电话联系你。
	//姓名：张三
	//电话：13800000000
	//时间：2016-04-20 10:00:00
	//感谢你的使用。

	data := Template{
		TemplateId: "m186TiCsxRBlWyOMolOEWmj-25WVdqi96wS7jE4X1eY",
		Touser:     openID,
		Url:        Url,
		Data: map[string]struct {
			Value string `json:"value"`
			Color string `json:"color,omitempty"`
		}{
			"first": {
				Value: "水站已接单, 以下是配送员信息",
				Color: "#173177",
			},
			"keyword1": {
				Value: userName,
				Color: "#173177",
			},
			"keyword2": {
				Value: Mobile,
				Color: "#173177",
			},
			"keyword3": {
				Value: orderTime,
				Color: "#173177",
			},
			"remark": {
				Value: remark,
				Color: "#173177",
			},
		},
	}

	Handel.sendTemplate(data)

	return nil
}
