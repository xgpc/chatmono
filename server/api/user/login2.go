package user

import (
	"chatmono/models"
	"chatmono/services/aes"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type userLogin2Req struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Login2 用户密码登录
// @Summary      用户密码登录
// @Description  用户密码登录
// @Accept       json
// @Produce      json
// @param        root  body  userLogin2Req  true  "参数"
// @Tags         用户管理
// @Success      200  {object}  render.Response
// @Router       /api/user/login2 [Post]
func Login2(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param userLogin2Req
	p.Init(&param)

	var userInfo models.User

	err := p.DB().Model(&userInfo).
		Where(models.UserColumns.UserName, param.UserName).
		Where(models.UserColumns.Password, aes.EnCode(param.Password)).
		First(&userInfo).Error
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	token := user.SetSession(userInfo.ID)

	// 去掉敏感信息
	userInfo.MobileData = nil

	res := map[string]interface{}{}
	res["token"] = token
	res["user"] = userInfo

	p.SuccessWithData(res)
}
