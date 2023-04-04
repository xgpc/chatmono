package openAI

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

type openAINewReq struct {
}

// New 创建会话
// @Summary      创建会话
// @Description  创建会话
// @Accept       json
// @Produce      json
// @param        root  body  openAINewReq  true  "参数"
// @Tags         chat
// @Success      200  {object}  render.Response
// @Router       /api/openAI/new [Post]
func New(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param openAINewReq
	p.Init(&param)

	// TODO: 后期实现

}
