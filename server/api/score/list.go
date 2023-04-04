package score

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type scoreListReq struct {
	ScoreType int `gorm:"column:score_type" json:"score_type"` // 积分类型:[0全部, 1打卡签到, 2分享推荐, 3购物(一般为负数)]
}

// List 我的积分
// @Summary      我的积分
// @Description  我的积分
// @Accept       json
// @Produce      json
// @param        root  body  scoreListReq  true  "参数"
// @Tags         积分管理
// @Success      200  {object}  render.Response
// @Router       /api/score/list [Post]
func List(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param scoreListReq
	p.Init(&param)
	var list []models.Score
	var total int64

	tx := p.DB().Model(&models.Score{}).
		Where(models.ScoreColumns.UserID, p.MyId())
	if param.ScoreType != 0 {
		tx = tx.Where(models.ScoreColumns.ScoreType, param.ScoreType)
	}
	err := tx.Count(&total).Scopes(cond.PageByQuery(ctx)).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithList(list, total)
}
