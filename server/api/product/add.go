package product

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"time"
)

type productAddReq struct {
	ProductName   string  `json:"product_name"`                          // 商品名称
	Price         float64 `json:"price"`                                 // 商品售价
	ProductType   int     `json:"product_type"`                          // 商品类型: [1销售 2积分兑换]
	ProductBody   string  `json:"product_body"`                          // 商品数据
	ProductStatus int     `json:"product_status"`                        // 商品状态:[1上架, 2下架]
	Score         int     `json:"score"`                                 // 商品积分(商品类型为2积分兑换时填入)
	RebateRate    float64 `gorm:"column:rebate_rate" json:"rebate_rate"` // 佣金比例
}

// Add 添加商品
// @Summary      添加商品
// @Description  添加商品(管理员权限)
// @Accept       json
// @Produce      json
// @param        root  body  productAddReq  true  "参数"
// @Tags         商品管理
// @Success      200  {object}  render.Response{data=models.Product}
// @Router       /api/product/add [Post]
func Add(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param productAddReq
	p.Init(&param)

	// 积分商品没有 价格
	if param.ProductType == 1 && param.Price == 0 {
		exce.ThrowSys(exce.CodeRequestError, "售卖类型的商品, 价格不能为空")
	}

	var info models.Product
	info.ProductName = param.ProductName
	info.Price = param.Price
	info.ProductType = param.ProductType
	info.ProductBody = param.ProductBody
	info.ProductStatus = param.ProductStatus
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = time.Now().Unix()
	info.CreatedUserID = p.MyId()
	info.Score = param.Score
	info.RebateRate = param.RebateRate

	err := p.DB().Model(&info).Create(&info).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.SuccessWithData(info)
}
