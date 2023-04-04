/**
 * @Author: smono
 * @Description:
 * @File:  get
 * @Version: 1.0.0
 * @Date: 2022/9/30 16:25
 */

package user

import (
	"chatmono/models"
	"chatmono/services/aes"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type getReq struct {
	ID uint32 `json:"id"` //user id
}

// Get 获取用户详细信息
// @Summary      获取用户详细信息
// @Description  获取用户详细信息
// @Accept       json
// @Produce      json
// @param        root  body  getReq  true  "参数"
// @Tags         用户管理
// @Success      200  {object}  render.Response
// @Router       /api/user/get [Post]
func Get(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param getReq
	p.Init(&param)

	mgr := models.UserMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	info.Mobile = aes.DeCode(info.MobileData)
	info.MobileData = nil

	p.SuccessWithData(info)

}
