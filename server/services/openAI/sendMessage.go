package openAI

import (
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/util"
	"io/ioutil"
	"net/http"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ReqOpenAI struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

func (p *ReqOpenAI) String() string {
	marshal, err := util.JsonEncode(p)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	return string(marshal)
}

func (p *ReqOpenAI) Push(msg Message) {
	p.Messages = append(p.Messages, msg)
}

type RspOpenAI map[string]interface{}

func SendMessage(reqOpenAI *ReqOpenAI) (RspOpenAI, error) {
	url := "https://api.openai.com/v1/chat/completions"
	method := "POST"

	payload := strings.NewReader(reqOpenAI.String())

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+_conf.OpenaiApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rsp RspOpenAI

	err = util.JsonDecode(body, &rsp)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
