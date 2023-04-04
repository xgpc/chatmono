package product

import (
	"chatmono/models"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/models/cond"
)

type productQueryReq struct {
	ProductName   string `gorm:"column:product_name" json:"product_name"`     // 商品名称
	ProductType   int    `gorm:"column:product_type" json:"product_type"`     // 商品类型: [0全部, 1销售 2积分]
	ProductStatus int    `gorm:"column:product_status" json:"product_status"` // 商品状态:[0全部, 1上架, 2下架]
}

// Query 商品查询
// @Summary      商品查询
// @Description  商品查询
// @Accept       json
// @Produce      json
// @param        root  body  productQueryReq  true  "参数"
// @param        page               query    int     true  "页数"
// @param        page_size           query    int     true  "页展示条数"
// @Tags            商品管理
// @Success      200  {object}  render.Response
// @Router       /api/product/query [Post]
func Query(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param productQueryReq
	p.Init(&param)
	var list []models.Product
	var total int64
	tx := p.DB().Model(&models.Product{})

	if param.ProductName != "" {
		tx = tx.Scopes(cond.Like(models.ProductColumns.ProductName, param.ProductName))
	}
	if param.ProductType != 0 {
		tx = tx.Where(models.ProductColumns.ProductType, param.ProductType)
	}
	if param.ProductStatus != 0 {
		tx = tx.Where(models.ProductColumns.ProductStatus, param.ProductStatus)
	}

	err := tx.Count(&total).Scopes(cond.PageByQuery(ctx)).Find(&list).Error
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	p.SuccessWithList(list, total)
}
