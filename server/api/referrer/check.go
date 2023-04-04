package referrer

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

type referrerCheckReq struct {
}

// Check 查询是否已经有推荐人
// @Summary      查询是否已经有推荐人
// @Description  查询是否已经有推荐人
// @Accept       json
// @Produce      json
// @param        root  body  referrerCheckReq  true  "参数"
// @Tags         推荐管理
// @Success      200  {object}  render.Response
// @Router       /api/referrer/check [Post]
func Check(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param referrerCheckReq
	p.Init(&param)

	var info models.ScoreReferrer

	p.DB().Model(&info).
		Where(models.ScoreReferrerColumns.UserID, p.MyId()).First(&info)
	p.SuccessWithData(info)
}
