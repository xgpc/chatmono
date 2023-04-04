/**
 * @Author: smono
 * @Description:
 * @File:  requert
 * @Version: 1.0.0
 * @Date: 2022/9/28 10:19
 */

package idCard

import (
	"fmt"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetIDCardInfo(name, idCard string) *Res {

	url := "https://idenauthen.market.alicloudapi.com/idenAuthentication"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(
		"idNo=%s"+
			"&name=%s", idCard, name))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy)
	}
	req.Header.Add("Authorization", "APPCODE cace38b77f72401289d716039468fa4f")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	res, err := client.Do(req)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy)
	}

	var info Res

	util.JsonDecode(body, &info)

	return &info
}

type Res struct {
	Name        string `json:"name"`
	IdNo        string `json:"idNo"`
	RespMessage string `json:"respMessage"`
	RespCode    string `json:"respCode"`
	Province    string `json:"province"`
	City        string `json:"city"`
	County      string `json:"county"`
	Birthday    string `json:"birthday"`
	Sex         string `json:"sex"`
	Age         string `json:"age"`
}

var genderDict map[string]int = map[string]int{
	"M": 1,
	"F": 2,
}

func (p *Res) Gender() int {
	return genderDict[p.Sex]
}
