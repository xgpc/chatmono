package session

import (
	"chatmono/services/openAI"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

type sessionSetReq struct {
	SessionKey  string             `json:"session_key"`  // 会话key
	SessionData openAI.SessionData `json:"session_data"` //会话数据
}

// Set 保存会话
// @Summary      保存会话
// @Description  保存会话
// @Accept       json
// @Produce      json
// @param        root  body  sessionSetReq  true  "参数"
// @Tags         会话管理
// @Success      200  {object}  render.Response
// @Router       /api/session/set [Post]
func Set(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param sessionSetReq
	p.Init(&param)

	openAI.SetSession(p.MyId(), param.SessionKey, &param.SessionData)

	p.Success()

}
