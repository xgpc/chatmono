/**
 * @Author: smono
 * @Description:
 * @File:  up
 * @Version: 1.0.0
 * @Date: 2022/9/21 14:24
 */

package user

import (
	"chatmono/models"
	"chatmono/services/admin"
	"chatmono/services/aes"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type userUpReq struct {
	models.User
	Mobile     string `json:"mobile"`
	MobileData []byte `gorm:"column:mobile_data" json:"mobile_data"` // 自动生成 不允许修改
}

// PostUp 修改个人信息
// @Summary      修改个人信息
// @Description  修改个人信息
// @Accept       json
// @Produce      json
// @param        root  body  userUpReq  true  "参数"
// @Tags         用户管理
// @Success      200  {object}  render.Response
// @Router       /api/user/Up [Post]
func PostUp(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param userUpReq
	p.Init(&param)

	// 管理员权限
	if param.ID != p.MyId() {
		admin.CheckAdmin(p.MyId())
	}

	if param.Mobile != "" {
		param.User.MobileData = aes.EnCode(param.Mobile)
		param.User.Mobile = aes.EnMobile(param.Mobile)
	}

	err := p.DB().Model(&param.User).Where(models.UserColumns.ID, param.ID).Updates(&param.User).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.Success()
}
