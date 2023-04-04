package product

import (
	"chatmono/models"
	"chatmono/pkg"
	"github.com/kataras/iris/v12"
	"github.com/shopspring/decimal"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

type productUpReq struct {
	ID            uint32  `json:"id"`                       // 商品ID
	ProductName   string  `json:"product_name,omitempty"`   // 商品名称
	Price         float64 `json:"price,omitempty"`          // 商品售价
	ProductType   int     `json:"product_type,omitempty"`   // 商品类型: [1销售 2积分兑换]
	ProductBody   string  `json:"product_body,omitempty"`   // 商品数据
	ProductStatus int     `json:"product_status,omitempty"` // 商品状态:[1上架, 2下架]
	Score         int     `json:"score,omitempty"`          // 商品积分(商品类型为2积分兑换时填入)
	RebateRate    float64 `json:"rebate_rate"`              // 佣金比例
}

// Up 修改商品
// @Summary      修改商品
// @Description  修改商品
// @Accept       json
// @Produce      json
// @param        root  body  productUpReq  true  "参数"
// @Tags         商品管理
// @Success      200  {object}  render.Response
// @Router       /api/product/up [Post]
func Up(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param productUpReq
	p.Init(&param)

	// TODO: 检查数据 商品积分为0时不会更新
	body := pkg.StructToMap(param)
	err := p.DB().Model(&models.Product{}).
		Where(models.ProductColumns.ID, param.ID).
		Updates(body).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 分销金不能大于 售价 * 30%
	mgr := models.ProductMgr(p.DB())
	info, err := mgr.GetFromID(param.ID)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 计算分销金额
	d2 := decimal.NewFromFloat(info.Price).Mul(decimal.NewFromFloat(0.3))
	RebateRate, _ := d2.Float64()

	if RebateRate < info.RebateRate {
		exce.ThrowSys(exce.CodeRequestError, "数据已修改, 但分销金大于(售价*30%), 请重新调整后重试, 否则分销会失败")
	}

	p.SuccessWithData(info)
}
