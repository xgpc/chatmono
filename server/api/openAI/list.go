package openAI

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

type openAIListReq struct {
}

// List 会话列表
// @Summary      会话列表
// @Description  会话列表
// @Accept       json
// @Produce      json
// @param        root  body  openAIListReq  true  "参数"
// @Tags         openAI
// @Success      200  {object}  render.Response
// @Router       /api/openAI/list [Post]
func List(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param openAIListReq
	p.Init(&param)

}
