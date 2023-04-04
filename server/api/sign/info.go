package sign

import (
	"chatmono/services/signServer"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type signInfoReq struct {
}

type signInfoRsp struct {
	*signServer.SignData
	ConsecutiveScore int `json:"consecutive_score""`
}

// Info 近7天打卡信息
// @Summary      近7天打卡信息
// @Description  近7天打卡信息
// @Accept       json
// @Produce      json
// @param        root  body  signInfoReq  true  "参数"
// @Tags         打卡签到
// @Success      200  {object}  render.Response{data=signInfoRsp}
// @Router       /api/sign/info [Post]
func Info(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param signInfoReq
	p.Init(&param)
	res, err := signServer.SignUserGetInfo(p.MyId())
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	var rsp signInfoRsp
	rsp.SignData = res
	if res.Consecutive7 {
		rsp.ConsecutiveScore = 7
	} else {
		rsp.ConsecutiveScore = res.ConsecutiveNum + 1
	}

	p.SuccessWithData(rsp)
}
