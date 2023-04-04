/**
 * @Author: smono
 * @Description:
 * @File:  requert_test.go
 * @Version: 1.0.0
 * @Date: 2022/9/28 10:23
 */

package idCard

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	//idInfo := GetIDCardInfo("尚富荣", "530423199205270031")
	//if idInfo.RespCode != "0004" {
	//    t.Fatal(idInfo.RespCode)
	//}

	idInfo := GetIDCardInfo("郑淑芸", "430721199510164007")
	if idInfo.RespCode != "0000" {
		t.Fatal(idInfo.RespCode)
	}
	fmt.Println(idInfo)
}

func TestName1(t *testing.T) {

	url := "https://idenauthen.market.alicloudapi.com/idenAuthentication"
	method := "POST"

	payload := strings.NewReader("idNo=530423199205270030&name=%E5%B0%9A%E5%AF%8C%E8%8D%A3")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "APPCODE cace38b77f72401289d716039468fa4f")
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	//req.Header.Add("Accept",	"*/*")
	//req.Header.Add("Host",	"idenauthen.market.alicloudapi.com")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmp := string(body)

	fmt.Println(tmp)

}
