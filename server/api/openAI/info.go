package openAI

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type openAIInfoReq struct {
	SessionKey string `json:"session_key"`
}

// Info 对话信息
// @Summary      对话信息
// @Description  对话信息
// @Accept       json
// @Produce      json
// @param        root  body  openAIInfoReq  true  "参数"
// @Tags         openAI
// @Success      200  {object}  render.Response
// @Router       /api/openAI/info [Post]
func Info(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param openAIInfoReq
	p.Init(&param)

	result, err := p.Redis().Get(context.Background(), param.SessionKey).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	//var rsp openAI.RspOpenAI
	//
	//util.json

	p.SuccessWithData(result)
}
