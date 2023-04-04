/**
 * @Author: smono
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2022/9/19 17:40
 */

package dict

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type dictQueryReq struct {
	Label    string `gorm:"column:label" json:"label"`         // 标签名称,模糊查询Label     string `gorm:"column:label" json:"label"`
	DictType string `gorm:"column:dict_type" json:"dict_type"` // 标签类型 all 查询所有
}

// PostQuery 字典查询
// @Summary      字典查询
// @Description  字典查询
// @Accept       json
// @Produce      json
// @param        root  body  dictQueryReq  true  "参数"
// @Tags         字典管理
// @Success      200  {object}  render.Response
// @Router       /api/dict/query [Post]
func PostQuery(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param dictQueryReq
	p.Init(&param)
	var list []models.Dict

	tx := p.DB().Model(&models.Dict{})

	if param.DictType != "all" {
		tx = tx.Where(models.DictColumns.DictType, param.DictType)
	}

	if param.Label != "" {
		tx = tx.Scopes(cond.Like(models.DictColumns.Label, param.Label))
	}

	err := tx.Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithList(list, len(list))
}
