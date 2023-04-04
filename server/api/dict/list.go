/**
 * @Author: smono
 * @Description:
 * @File:  list
 * @Version: 1.0.0
 * @Date: 2022/9/28 10:43
 */

package dict

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type reqList struct {
	DictType string `json:"dict_type"` // 字典类型:[all 所有]
}

// List         字典列表
// @Summary      字典列表
// @Description  字典列表
// @Accept       json
// @Produce      json
// @param        root  body  reqList  true  "参数"
// @param        page               query    int     true  "页数"
// @param        page_size           query    int     true  "页展示条数"
// @Tags         字典管理
// @Success      200  {object}  render.Response
// @Router       /api/dict/list [Post]
func List(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param reqList
	p.Init(&param)
	var list []models.Dict
	var total int64
	tx := p.DB().Model(&models.Dict{})

	if param.DictType != "all" {
		tx = tx.Where(models.DictColumns.DictType, param.DictType)
	}

	err := tx.Count(&total).Scopes(cond.PageByQuery(ctx)).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithList(list, total)
}
