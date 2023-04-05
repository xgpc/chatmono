package user

import (
	"chatmono/models"
	"chatmono/services/aes"
	"chatmono/services/user"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type userLogonReq struct {
	Mobile   string `json:"mobile"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	UserImg  string `json:"user_img"`
}

// Logon 注册
// @Summary      注册
// @Description  注册
// @Accept       json
// @Produce      json
// @param        root  body  userLogonReq  true  "参数"
// @Tags         用户管理
// @Success      200  {object}  render.Response
// @Router       /api/user/logon [Post]
func Logon(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param userLogonReq
	p.Init(&param)

	var userInfo models.User
	var num int64
	err2 := p.DB().Model(&userInfo).
		Where(models.UserColumns.MobileData, aes.EnCode(param.Mobile)).
		Count(&num).Error

	if err2 != nil {
		exce.ThrowSys(exce.CodeSysBusy, err2)
	}
	if num != 0 {
		exce.ThrowSys(exce.CodeRequestError, "该账号已注册")
	}

	if param.UserName == "" {
		param.UserName = param.Mobile
	}

	err := p.DB().Model(&userInfo).
		Where(models.UserColumns.MobileData, aes.EnCode(param.Mobile)).
		Attrs(models.User{
			UserImg:    param.UserImg,
			UserName:   param.UserName,
			MobileData: aes.EnCode(param.Mobile),
			Mobile:     aes.EnMobile(param.Mobile),
			Password:   aes.EnCode(param.Password),
		}).
		FirstOrCreate(&userInfo).Error
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	token := user.SetSession(userInfo.ID)

	userInfo.Mobile = param.Mobile
	userInfo.MobileData = nil

	res := map[string]interface{}{}
	res["token"] = token
	res["user"] = userInfo

	p.SuccessWithData(res)

}
