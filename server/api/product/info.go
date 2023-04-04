package product

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type productInfoReq struct {
	ID uint32 `json:"id"` // 商品ID
}

// Info 商品信息
// @Summary      商品信息
// @Description  商品信息
// @Accept       json
// @Produce      json
// @param        root  body  productInfoReq  true  "参数"
// @Tags         商品管理
// @Success      200  {object}  render.Response{data=models.Product}
// @Router       /api/product/info [Post]
func Info(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param productInfoReq
	p.Init(&param)

	mgr := models.ProductMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(info)

}
