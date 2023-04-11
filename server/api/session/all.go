package session

import (
	"chatmono/services/openAI"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

type sessionAllReq struct {
}

// All 获取会话列表
// @Summary      获取会话列表
// @Description  获取会话列表
// @Accept       json
// @Produce      json
// @param        root  body  sessionAllReq  true  "参数"
// @Tags         会话管理
// @Success      200  {object}  render.Response
// @Router       /api/session/all [Post]
func All(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param sessionAllReq
	p.Init(&param)

	rsp := openAI.GetAll(p.MyId())

	p.SuccessWithData(rsp)
}
