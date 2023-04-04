package order

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

type orderRefreshReq struct {
	ID uint32 `json:"id"` // 订单ID
}

// Refresh 刷新订单
// @Summary      刷新订单
// @Description  刷新订单
// @Accept       json
// @Produce      json
// @param        root  body  orderRefreshReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response
// @Router       /api/order/refresh [Post]
func Refresh(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param orderRefreshReq
	p.Init(&param)

}
