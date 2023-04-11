package openAI

import (
	"chatmono/services/openAI"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type OpenAISendReq struct {
	SessionKey string `json:"session_key"` // 当前会话的KEY

	Messages []openAI.Message `json:"messages"`
}

// Send 发送消息
// @Summary      发送消息
// @Description  发送消息
// @Accept       json
// @Produce      json
// @param        root  body  openAISendReq  true  "参数"
// @Tags         openAI
// @Success      200  {object}  render.Response
// @Router       /api/openAI/send [Post]
func Send(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param OpenAISendReq

	p.Init(&param)

	if len(param.Messages) > 10 {
		exce.ThrowSys(exce.CodeRequestError, "当前仅支持5次上下文对话")
	}

	req := &openAI.ReqOpenAI{
		Model:       "gpt-3.5-turbo",
		Messages:    param.Messages,
		Temperature: 0.7,
	}

	rsp, err := openAI.SendMessage(req)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(rsp)
}
