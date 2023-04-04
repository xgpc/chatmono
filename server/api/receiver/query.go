/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/

package receiver

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type receiverQueryReq struct {
}

// Query 分销商查询
// @Summary      分销商查询
// @Description  分销商查询
// @Accept       json
// @Produce      json
// @param        root  body  receiverQueryReq  true  "参数"
// @param        page               query    int     true  "页数"
// @param        page_size           query    int     true  "页展示条数"
// @Tags         分销管理
// @Success      200  {object}  render.Response
// @Router       /api/receiver/query [Post]
func Query(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverQueryReq
	p.Init(&param)

	var list []models.User
	var total int64

	var ids []uint32
	err := p.DB().Model(&models.WxUser{}).
		Where(models.WxUserColumns.IsProfitsharing, models.IsProfitsharingTrue).
		Count(&total).
		Scopes(cond.PageByQuery(ctx)).
		Select(models.WxUserColumns.UserID).Find(&ids).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	err = p.DB().Model(&models.User{}).Where(models.UserColumns.ID, ids).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithList(list, total)
}
