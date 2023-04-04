package score

import (
	"chatmono/models"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type scoreTotalReq struct {
}

// Total 我的总积分
// @Summary      我的总积分
// @Description  我的总积分
// @Accept       json
// @Produce      json
// @param        root  body  scoreTotalReq  true  "参数"
// @Tags         积分管理
// @Success      200  {object}  render.Response
// @Router       /api/score/total [Post]
func Total(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param scoreTotalReq
	p.Init(&param)

	var num int64
	err := p.DB().Debug().Model(&models.Score{}).
		Where(models.ScoreColumns.UserID, p.MyId()).
		Select(fmt.Sprintf("SUM(%s)", models.ScoreColumns.ScoreNum)).Find(&num).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(map[string]interface{}{
		"total": num,
	})

}
