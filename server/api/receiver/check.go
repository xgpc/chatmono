package receiver

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type receiverCheckReq struct {
}

// Check 查询是否开通分销
// @Summary      查询是否开通分销
// @Description  查询是否开通分销
// @Accept       json
// @Produce      json
// @param        root  body  receiverCheckReq  true  "参数"
// @Tags         分销管理
// @Success      200  {object}  render.Response{data=models.WxUser}
// @Router       /api/receiver/check [Post]
func Check(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param receiverCheckReq
	p.Init(&param)

	var info models.WxUser
	err := p.DB().Model(&info).Where(models.WxUserColumns.UserID, p.MyId()).Find(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(info)
}
