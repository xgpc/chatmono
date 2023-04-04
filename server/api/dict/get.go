/**
 * @Author: smono
 * @Description:
 * @File:  get
 * @Version: 1.0.0
 * @Date: 2022/10/2 11:09
 */

package dict

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type GetReq struct {
	ID uint32 `json:"id"` // 字典ID
}

// Get 获取字典信息
// @Summary      获取字典信息
// @Description  获取字典信息
// @Accept       json
// @Produce      json
// @param        root  body  GetReq  true  "参数"
// @Tags         字典管理
// @Success      200  {object}  render.Response
// @Router       /api/dict/get [Post]
func Get(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param GetReq
	p.Init(&param)

	mgr := models.DictMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(info)

}
