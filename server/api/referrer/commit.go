package referrer

import (
	"chatmono/config"
	"chatmono/models"
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"gorm.io/gorm"
	"time"
)

type referrerCommitReq struct {
	ReferrerUserID uint32 `json:"referrer_user_id"` // 推荐人ID
}

// Commit 提交成为被推荐人
// @Summary      提交成为被推荐人
// @Description  提交成为被推荐人
// @Accept       json
// @Produce      json
// @param        root  body  referrerCommitReq  true  "参数"
// @Tags         推荐管理
// @Success      200  {object}  render.Response
// @Router       /api/referrer/commit [Post]
func Commit(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param referrerCommitReq
	p.Init(&param)

	var info models.ScoreReferrer
	var Score models.Score
	err := p.DB().Transaction(func(tx *gorm.DB) error {

		// 查询数据, 没有就创建
		err2 := tx.Model(&info).
			Where(models.ScoreReferrerColumns.UserID, p.MyId()).
			Attrs(models.ScoreReferrer{
				UserID:      p.MyId(),
				RecommendAt: time.Now().Unix(),
			}).
			FirstOrCreate(&info).Error
		if err2 != nil {
			return err2
		}

		//保存
		if info.ReferrerUserID != 0 {
			return errors.New("已经拥有推荐人, 请勿重复提交")
		}

		info.ReferrerUserID = param.ReferrerUserID

		err := tx.Where(models.ScoreReferrerColumns.UserID, p.MyId()).Save(&info).Error
		if err != nil {
			return err
		}
		// 添加积分 推荐人 获得积分

		Score.ScoreType = models.ScoreTypeShare
		Score.CreatedAt = time.Now().Unix()
		Score.UserID = param.ReferrerUserID
		Score.ScoreNum = config.Config.ReferrerScoreNum

		return tx.Create(&Score).Error
	})
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(map[string]interface{}{
		"info":  info,
		"score": Score,
	})
}
