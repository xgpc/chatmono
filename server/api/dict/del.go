/**
 * @Author: smono
 * @Description:
 * @File:  del
 * @Version: 1.0.0
 * @Date: 2022/9/19 20:54
 */

package dict

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type dictDelReq struct {
	ID uint32 `json:"id"` // 字典ID
}

// PostDel 删除字典
// @Summary      删除字典
// @Description  删除字典
// @Accept       json
// @Produce      json
// @param        root  body  dictDelReq  true  "参数"
// @Tags         字典管理
// @Success      200  {object}  render.Response
// @Router       /api/dict/del [Post]
func PostDel(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param dictDelReq
	p.Init(&param)

	var info models.Dict

	// TODO: 权限判断
	err := p.DB().Model(&models.Dict{}).
		Where(models.DictColumns.ID, param.ID).Delete(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.Success()

}
