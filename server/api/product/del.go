package product

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type productDelReq struct {
	ID uint32 `json:"id"` // 商品ID
}

// Del 商品删除
// @Summary      商品删除
// @Description  商品删除
// @Accept       json
// @Produce      json
// @param        root  body  productDelReq  true  "参数"
// @Tags         商品管理
// @Success      200  {object}  render.Response
// @Router       /api/product/del [Post]
func Del(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param productDelReq
	p.Init(&param)

	exce.ThrowSys(exce.CodeRequestError, "暂不提供")

}
