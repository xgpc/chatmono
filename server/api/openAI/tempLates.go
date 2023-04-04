package openAI

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type openAITemplatesReq struct {
}

// Templates 获取当前所有模板
// @Summary      获取当前所有模板
// @Description  获取当前所有模板
// @Accept       json
// @Produce      json
// @param        root  body  openAITemplatesReq  true  "参数"
// @Tags         openAI
// @Success      200  {object}  render.Response
// @Router       /api/openAI/templates/list [Post]
func Templates(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param openAITemplatesReq
	p.Init(&param)

	var list []models.SessionTemplates
	var total int64

	err := p.DB().Model(&models.SessionTemplates{}).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	total = int64(len(list))
	p.SuccessWithList(list, total)
}
