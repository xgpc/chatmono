/**
 * @Author: smono
 * @Description:
 * @File:  decrypt
 * @Version: 1.0.0
 * @Date: 2022/9/21 11:27
 */

package wxMini

import (
	"chatmono/services/wechat/wechatMini"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type wxMiniDecryptReq struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

// PostDecrypt 数据解密
// @Summary      数据解密
// @Description  数据解密
// @Accept       json
// @Produce      json
// @param        root  body  wxMiniDecryptReq  true  "参数"
// @Tags         小程序
// @Success      200  {object}  render.Response
// @Router       /api/wxMini/decrypt [Post]
func PostDecrypt(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param wxMiniDecryptReq
	p.Init(&param)

	res, err := wechatMini.Decrypt(param.SessionKey, param.EncryptedData, param.Iv)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(res)
}
