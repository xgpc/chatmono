/**
 * @Author: smono
 * @Description:
 * @File:  created
 * @Version: 1.0.0
 * @Date: 2022/9/19 20:52
 */

package dict

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type dictCreatedReq struct {
	ID uint32 `json:"id"` // 系统自动创建
	models.Dict
}

// PostCreated 创建字典
// @Summary      创建字典
// @Description  创建字典
// @Accept       json
// @Produce      json
// @param        root  body  dictCreatedReq  true  "参数"
// @Tags         字典管理
// @Success      200  {object}  render.Response
// @Router       /api/dict/created [Post]
func PostCreated(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param dictCreatedReq
	p.Init(&param)

	// TODO: 权限判断

	//param.CreatedUserID = p.MyId()
	err := p.DB().Model(&models.Dict{}).Create(&param.Dict).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(param.Dict)
}
