/**
 * @Author: smono
 * @Description:
 * @File:  up
 * @Version: 1.0.0
 * @Date: 2022/9/25 0:45
 */

package dict

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type upReq struct {
	models.Dict
}

// Up 修改字典
// @Summary      修改字典
// @Description  修改字典
// @Accept       json
// @Produce      json
// @param        root  body  upReq  true  "参数"
// @Tags         字典管理
// @Success      200  {object}  render.Response
// @Router       /api/dict/up [Post]
func Up(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param upReq
	p.Init(&param)

	err := p.DB().Model(&models.Dict{}).
		Where(models.DictColumns.ID, param.ID).
		Updates(&param.Dict).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(param.Dict)
}
