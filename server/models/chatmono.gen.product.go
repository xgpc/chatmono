package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProductMgr struct {
	*_BaseMgr
}

// ProductMgr open func
func ProductMgr(db *gorm.DB) *_ProductMgr {
	if db == nil {
		panic(fmt.Errorf("ProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductMgr) GetTableName() string {
	return "product"
}

// Reset 重置gorm会话
func (obj *_ProductMgr) Reset() *_ProductMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductMgr) Get() (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductMgr) Gets() (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Product{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 商品ID
func (obj *_ProductMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithProductName product_name获取 商品名称
func (obj *_ProductMgr) WithProductName(productName string) Option {
	return optionFunc(func(o *options) { o.query["product_name"] = productName })
}

// WithPrice price获取 商品售价
func (obj *_ProductMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithProductType product_type获取 商品类型: [1销售 2积分]
func (obj *_ProductMgr) WithProductType(productType int) Option {
	return optionFunc(func(o *options) { o.query["product_type"] = productType })
}

// WithProductBody product_body获取 商品数据
func (obj *_ProductMgr) WithProductBody(productBody string) Option {
	return optionFunc(func(o *options) { o.query["product_body"] = productBody })
}

// WithProductStatus product_status获取 商品状态:[1上架, 2下架]
func (obj *_ProductMgr) WithProductStatus(productStatus int) Option {
	return optionFunc(func(o *options) { o.query["product_status"] = productStatus })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ProductMgr) WithCreatedAt(createdAt int64) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 修改时间
func (obj *_ProductMgr) WithUpdatedAt(updatedAt int64) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithCreatedUserID created_user_id获取 创建人
func (obj *_ProductMgr) WithCreatedUserID(createdUserID uint32) Option {
	return optionFunc(func(o *options) { o.query["created_user_id"] = createdUserID })
}

// WithScore score获取 商品积分
func (obj *_ProductMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithRebateRate rebate_rate获取 佣金(每份商品可得佣金)
func (obj *_ProductMgr) WithRebateRate(rebateRate float64) Option {
	return optionFunc(func(o *options) { o.query["rebate_rate"] = rebateRate })
}

// GetByOption 功能选项模式获取
func (obj *_ProductMgr) GetByOption(opts ...Option) (result Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductMgr) GetByOptions(opts ...Option) (results []*Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ProductMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Product, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 商品ID
func (obj *_ProductMgr) GetFromID(id uint32) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 商品ID
func (obj *_ProductMgr) GetBatchFromID(ids []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromProductName 通过product_name获取内容 商品名称
func (obj *_ProductMgr) GetFromProductName(productName string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_name` = ?", productName).Find(&results).Error

	return
}

// GetBatchFromProductName 批量查找 商品名称
func (obj *_ProductMgr) GetBatchFromProductName(productNames []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_name` IN (?)", productNames).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品售价
func (obj *_ProductMgr) GetFromPrice(price float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品售价
func (obj *_ProductMgr) GetBatchFromPrice(prices []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromProductType 通过product_type获取内容 商品类型: [1销售 2积分]
func (obj *_ProductMgr) GetFromProductType(productType int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_type` = ?", productType).Find(&results).Error

	return
}

// GetBatchFromProductType 批量查找 商品类型: [1销售 2积分]
func (obj *_ProductMgr) GetBatchFromProductType(productTypes []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_type` IN (?)", productTypes).Find(&results).Error

	return
}

// GetFromProductBody 通过product_body获取内容 商品数据
func (obj *_ProductMgr) GetFromProductBody(productBody string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_body` = ?", productBody).Find(&results).Error

	return
}

// GetBatchFromProductBody 批量查找 商品数据
func (obj *_ProductMgr) GetBatchFromProductBody(productBodys []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_body` IN (?)", productBodys).Find(&results).Error

	return
}

// GetFromProductStatus 通过product_status获取内容 商品状态:[1上架, 2下架]
func (obj *_ProductMgr) GetFromProductStatus(productStatus int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_status` = ?", productStatus).Find(&results).Error

	return
}

// GetBatchFromProductStatus 批量查找 商品状态:[1上架, 2下架]
func (obj *_ProductMgr) GetBatchFromProductStatus(productStatuss []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`product_status` IN (?)", productStatuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ProductMgr) GetFromCreatedAt(createdAt int64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_ProductMgr) GetBatchFromCreatedAt(createdAts []int64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 修改时间
func (obj *_ProductMgr) GetFromUpdatedAt(updatedAt int64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 修改时间
func (obj *_ProductMgr) GetBatchFromUpdatedAt(updatedAts []int64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromCreatedUserID 通过created_user_id获取内容 创建人
func (obj *_ProductMgr) GetFromCreatedUserID(createdUserID uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`created_user_id` = ?", createdUserID).Find(&results).Error

	return
}

// GetBatchFromCreatedUserID 批量查找 创建人
func (obj *_ProductMgr) GetBatchFromCreatedUserID(createdUserIDs []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`created_user_id` IN (?)", createdUserIDs).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 商品积分
func (obj *_ProductMgr) GetFromScore(score int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 商品积分
func (obj *_ProductMgr) GetBatchFromScore(scores []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromRebateRate 通过rebate_rate获取内容 佣金(每份商品可得佣金)
func (obj *_ProductMgr) GetFromRebateRate(rebateRate float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`rebate_rate` = ?", rebateRate).Find(&results).Error

	return
}

// GetBatchFromRebateRate 批量查找 佣金(每份商品可得佣金)
func (obj *_ProductMgr) GetBatchFromRebateRate(rebateRates []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`rebate_rate` IN (?)", rebateRates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductMgr) FetchByPrimaryKey(id uint32) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByID primary or index 获取唯一内容
func (obj *_ProductMgr) FetchUniqueByID(id uint32) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`id` = ?", id).First(&result).Error

	return
}
