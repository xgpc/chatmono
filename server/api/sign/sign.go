package sign

import (
	"chatmono/models"
	"chatmono/services/signServer"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"time"
)

type signReq struct {
}

// Sign 打卡签到
// @Summary      打卡签到
// @Description  打卡签到(返回今日第几个签到) 7天以上7分, 否则1天1分 2天2分 3天3分
// @Accept       json
// @Produce      json
// @param        root  body  signReq  true  "参数"
// @Tags         打卡签到
// @Success      200  {object}  render.Response
// @Router       /api/sign/ [Post]
func Sign(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param signReq
	p.Init(&param)

	res, err := signServer.SignUser(p.MyId())
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 添加积分
	signInfo, err := signServer.SignUserGetInfo(p.MyId())
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	var Score int

	// 7天以上7分, 否则1天1分 2天2分 3天3分
	if signInfo.Consecutive7 {
		Score = 7
	} else {
		Score = signInfo.ConsecutiveNum + 1
	}

	// 添加积分
	var info models.Score

	t := time.Now()
	CreatedAt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	// 打卡签到情况第一个就是当天时间

	err = p.DB().Model(&info).
		Where(models.ScoreColumns.UserID, p.MyId()).
		Where(models.ScoreColumns.CreatedAt, CreatedAt).
		Attrs(models.Score{
			UserID:    p.MyId(),
			ScoreNum:  Score,
			ScoreType: models.ScoreTypeSign,
			CreatedAt: CreatedAt,
		}).
		FirstOrCreate(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(map[string]interface{}{
		"num":        res,
		"sign_info":  signInfo,
		"score_info": info,
	})

}

// 获取近7日内的打卡签到积分情况
