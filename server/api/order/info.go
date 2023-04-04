package order

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type orderInfoReq struct {
	ID uint32 `json:"id"` // 订单ID
}

// Info 订单详情
// @Summary      订单详情
// @Description  订单详情
// @Accept       json
// @Produce      json
// @param        root  body  orderInfoReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response{data=models.ProductOrder}
// @Router       /api/order/info [Post]
func Info(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param orderInfoReq
	p.Init(&param)

	var info models.ProductOrder

	err := p.DB().Debug().Model(&info).
		Preload("Product").
		Preload("Address").
		Where(models.ProductOrderColumns.ID, param.ID).
		First(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithData(info)
}
