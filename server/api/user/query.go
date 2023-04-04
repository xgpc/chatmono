/**
 * @Author: smono
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2022/9/25 1:01
 */

package user

import (
	"chatmono/models"
	"chatmono/services/aes"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type QueryReq struct {
	Mobile   string `json:"mobile"`
	UserName string `json:"user_name"`
}

// Query 查询用户
// @Summary      查询用户
// @Description  查询用户
// @Accept       json
// @Produce      json
// @param        root  body  QueryReq  true  "参数"
// @Tags         用户管理
// @Success      200  {object}  render.Response
// @Router       /api/user/Query [Post]
func Query(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param QueryReq
	p.Init(&param)

	tx := p.DB().Model(&models.User{})

	if param.UserName == "" && param.Mobile == "" {
		exce.ThrowSys(exce.CodeRequestError, "必须填写一个参数")
	}

	if param.Mobile != "" {
		tx = tx.Where(models.UserColumns.MobileData, aes.EnCode(param.Mobile))
	}

	if param.UserName != "" {
		tx = tx.Scopes(cond.Like(models.UserColumns.UserName, param.UserName))
	}

	var list []models.User
	err := tx.Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithList(list, len(list))
}
