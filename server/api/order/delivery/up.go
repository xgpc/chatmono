package delivery

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/util"
)

type deliveryUpReq struct {
	ID             uint32 `json:"id"`                                              // 订单ID
	DeliveryStatus int    `gorm:"column:delivery_status" json:"delivery_status"`   // 配送状态:[1未接单, 2已接单/已发货, 3已送达, 4已取消]
	DeliveryUserID uint32 `gorm:"column:delivery_user_id" json:"delivery_user_id"` // 配送人ID/发货人ID
	DeliveryStatAt int64  `gorm:"column:delivery_stat_at" json:"delivery_stat_at"` // 接单时间
	DeliveryEndAt  int64  `gorm:"column:delivery_end_at" json:"delivery_end_at"`   // 送达时间
	ExpressSn      string `gorm:"column:express_sn" json:"express_sn"`             // 快递单号
	DeliveryType   int    `gorm:"column:delivery_type" json:"delivery_type"`       // 配送类型:[1同城配送, 2快递配送]
}

// DeliveryUp 修改配送状态
// @Summary      修改配送状态
// @Description  修改配送状态(如果配送状态为已送到, 将触发分账, 并返回分账信息data['profitsharing'])
// @Accept       json
// @Produce      json
// @param        root  body  deliveryUpReq  true  "参数"
// @Tags         订单管理
// @Success      200  {object}  render.Response
// @Router       /api/order/delivery/up [Post]
func DeliveryUp(ctx iris.Context) {
	p := frame.NewBase(ctx)
	md := map[string]interface{}{}
	var param deliveryUpReq
	p.Init(&param)
	var info models.ProductOrder

	// 修改订单状态
	body := util.StructToMapByRef(param)
	err := p.DB().Model(&info).Where(models.ProductOrderColumns.ID, param.ID).Updates(body).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 获取订单详情
	err = p.DB().Model(&info).Where(models.ProductOrderColumns.ID, param.ID).Preload("Product").First(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	md["order"] = info

	p.SuccessWithData(md)
}
