package session

import (
	"chatmono/services/openAI"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type sessionGetReq struct {
	SessionKey string `json:"session_key"`
}

// Get 会话数据
// @Summary      会话数据
// @Description  会话数据
// @Accept       json
// @Produce      json
// @param        root  body  sessionGetReq  true  "参数"
// @Tags         会话管理
// @Success      200  {object}  render.Response
// @Router       /api/session/get [Post]
func Get(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param sessionGetReq
	p.Init(&param)

	session, err := openAI.GetSession(p.MyId(), param.SessionKey)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(session)
}
