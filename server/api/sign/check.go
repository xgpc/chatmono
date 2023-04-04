package sign

import (
	"chatmono/services/signServer"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

type signCheckReq struct {
}

// Check 查询今日是否签到
// @Summary      查询今日是否签到
// @Description  查询今日是否签到
// @Accept       json
// @Produce      json
// @param        root  body  signCheckReq  true  "参数"
// @Tags         打卡签到
// @Success      200  {object}  render.Response
// @Router       /api/sign/check [Post]
func Check(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param signCheckReq
	p.Init(&param)

	res := signServer.SignUserCheckToday(p.MyId())
	//info, err := signServer.SignUserGetInfo(p.MyId())
	//if err != nil {
	//	exce.ThrowSys(exce.CodeRequestError, err.Error())
	//}

	p.SuccessWithData(map[string]interface{}{
		"is_sign": res,
		//"sign_info": info,
	})

}
