package session

import (
	"chatmono/services/openAI"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type DelsessiondelReq struct {
	SessionKey string `json:"session_key"`
}

// Del 删除会话
// @Summary      删除会话
// @Description  删除会话
// @Accept       json
// @Produce      json
// @param        root  body  sessiondelReq  true  "参数"
// @Tags         会话管理
// @Success      200  {object}  render.Response
// @Router       /api/session/del [Post]
func Del(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param DelsessiondelReq
	p.Init(&param)

	err := openAI.DelSession(p.MyId(), param.SessionKey)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.Success()

}
