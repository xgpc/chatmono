package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProductOrderMgr struct {
	*_BaseMgr
}

// ProductOrderMgr open func
func ProductOrderMgr(db *gorm.DB) *_ProductOrderMgr {
	if db == nil {
		panic(fmt.Errorf("ProductOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductOrderMgr) GetTableName() string {
	return "product_order"
}

// Reset 重置gorm会话
func (obj *_ProductOrderMgr) Reset() *_ProductOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductOrderMgr) Get() (result ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductOrderMgr) Gets() (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ProductOrderMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_ProductOrderMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithPrice price获取 该单商品总金额 (商品单价 * 购买数量)
func (obj *_ProductOrderMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithProductID product_id获取 商品ID
func (obj *_ProductOrderMgr) WithProductID(productID uint32) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithProductNum product_num获取 商品购买数量
func (obj *_ProductOrderMgr) WithProductNum(productNum int) Option {
	return optionFunc(func(o *options) { o.query["product_num"] = productNum })
}

// WithCreatedUserID created_user_id获取 下单人ID
func (obj *_ProductOrderMgr) WithCreatedUserID(createdUserID uint32) Option {
	return optionFunc(func(o *options) { o.query["created_user_id"] = createdUserID })
}

// WithPayStatus pay_status获取 支付状态:[1未支付, 2已取消或超时, 3支付完成, 4已退款]
func (obj *_ProductOrderMgr) WithPayStatus(payStatus int) Option {
	return optionFunc(func(o *options) { o.query["pay_status"] = payStatus })
}

// WithTransactionID transaction_id获取 微信支付单号
func (obj *_ProductOrderMgr) WithTransactionID(transactionID string) Option {
	return optionFunc(func(o *options) { o.query["transaction_id"] = transactionID })
}

// WithCreatedAt created_at获取 下单时间
func (obj *_ProductOrderMgr) WithCreatedAt(createdAt int64) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithDeliveryStatus delivery_status获取 配送状态:[1未接单, 2已接单/已发货, 3已送达, 4已取消]
func (obj *_ProductOrderMgr) WithDeliveryStatus(deliveryStatus int) Option {
	return optionFunc(func(o *options) { o.query["delivery_status"] = deliveryStatus })
}

// WithAddressID address_id获取 配送地址ID
func (obj *_ProductOrderMgr) WithAddressID(addressID uint32) Option {
	return optionFunc(func(o *options) { o.query["address_id"] = addressID })
}

// WithDeliveryUserID delivery_user_id获取 配送人ID/发货人ID
func (obj *_ProductOrderMgr) WithDeliveryUserID(deliveryUserID uint32) Option {
	return optionFunc(func(o *options) { o.query["delivery_user_id"] = deliveryUserID })
}

// WithDeliveryStatAt delivery_stat_at获取 接单时间
func (obj *_ProductOrderMgr) WithDeliveryStatAt(deliveryStatAt int64) Option {
	return optionFunc(func(o *options) { o.query["delivery_stat_at"] = deliveryStatAt })
}

// WithDeliveryEndAt delivery_end_at获取 送达时间
func (obj *_ProductOrderMgr) WithDeliveryEndAt(deliveryEndAt int64) Option {
	return optionFunc(func(o *options) { o.query["delivery_end_at"] = deliveryEndAt })
}

// WithOpenID open_id获取
func (obj *_ProductOrderMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithExpressSn express_sn获取 快递单号
func (obj *_ProductOrderMgr) WithExpressSn(expressSn string) Option {
	return optionFunc(func(o *options) { o.query["express_sn"] = expressSn })
}

// WithDeliveryType delivery_type获取 配送类型:[1同城配送, 2快递配送]
func (obj *_ProductOrderMgr) WithDeliveryType(deliveryType int) Option {
	return optionFunc(func(o *options) { o.query["delivery_type"] = deliveryType })
}

// WithRefundStatus refund_status获取 退款状态:[1 正常, 2申请退款, 3驳回退款, 4同意退款]
func (obj *_ProductOrderMgr) WithRefundStatus(refundStatus int) Option {
	return optionFunc(func(o *options) { o.query["refund_status"] = refundStatus })
}

// WithRefundReason refund_reason获取 退款原因
func (obj *_ProductOrderMgr) WithRefundReason(refundReason string) Option {
	return optionFunc(func(o *options) { o.query["refund_reason"] = refundReason })
}

// WithRefundExamineUser refund_examine_user获取 退款操作人ID
func (obj *_ProductOrderMgr) WithRefundExamineUser(refundExamineUser uint32) Option {
	return optionFunc(func(o *options) { o.query["refund_examine_user"] = refundExamineUser })
}

// WithRefundDetail refund_detail获取 退款详情
func (obj *_ProductOrderMgr) WithRefundDetail(refundDetail string) Option {
	return optionFunc(func(o *options) { o.query["refund_detail"] = refundDetail })
}

// WithReferrerID referrer_id获取 推荐人ID
func (obj *_ProductOrderMgr) WithReferrerID(referrerID uint32) Option {
	return optionFunc(func(o *options) { o.query["referrer_id"] = referrerID })
}

// WithProfitsharingStatus profitsharing_status获取 分账状态:[1 未分账, 2 已分帐]
func (obj *_ProductOrderMgr) WithProfitsharingStatus(profitsharingStatus int) Option {
	return optionFunc(func(o *options) { o.query["profitsharing_status"] = profitsharingStatus })
}

// WithProfitsharingMoney profitsharing_money获取 分账金额(商品分账金 * 购买数量)
func (obj *_ProductOrderMgr) WithProfitsharingMoney(profitsharingMoney float64) Option {
	return optionFunc(func(o *options) { o.query["profitsharing_money"] = profitsharingMoney })
}

// WithOrderType order_type获取 订单类型:[1在线支付, 2积分支付]
func (obj *_ProductOrderMgr) WithOrderType(orderType int) Option {
	return optionFunc(func(o *options) { o.query["order_type"] = orderType })
}

// GetByOption 功能选项模式获取
func (obj *_ProductOrderMgr) GetByOption(opts ...Option) (result ProductOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductOrderMgr) GetByOptions(opts ...Option) (results []*ProductOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ProductOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ProductOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_ProductOrderMgr) GetFromID(id uint32) (result ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ProductOrderMgr) GetBatchFromID(ids []uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_ProductOrderMgr) GetFromOrderSn(orderSn string) (result ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`order_sn` = ?", orderSn).First(&result).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_ProductOrderMgr) GetBatchFromOrderSn(orderSns []string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 该单商品总金额 (商品单价 * 购买数量)
func (obj *_ProductOrderMgr) GetFromPrice(price float64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 该单商品总金额 (商品单价 * 购买数量)
func (obj *_ProductOrderMgr) GetBatchFromPrice(prices []float64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromProductID 通过product_id获取内容 商品ID
func (obj *_ProductOrderMgr) GetFromProductID(productID uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`product_id` = ?", productID).Find(&results).Error

	return
}

// GetBatchFromProductID 批量查找 商品ID
func (obj *_ProductOrderMgr) GetBatchFromProductID(productIDs []uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`product_id` IN (?)", productIDs).Find(&results).Error

	return
}

// GetFromProductNum 通过product_num获取内容 商品购买数量
func (obj *_ProductOrderMgr) GetFromProductNum(productNum int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`product_num` = ?", productNum).Find(&results).Error

	return
}

// GetBatchFromProductNum 批量查找 商品购买数量
func (obj *_ProductOrderMgr) GetBatchFromProductNum(productNums []int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`product_num` IN (?)", productNums).Find(&results).Error

	return
}

// GetFromCreatedUserID 通过created_user_id获取内容 下单人ID
func (obj *_ProductOrderMgr) GetFromCreatedUserID(createdUserID uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`created_user_id` = ?", createdUserID).Find(&results).Error

	return
}

// GetBatchFromCreatedUserID 批量查找 下单人ID
func (obj *_ProductOrderMgr) GetBatchFromCreatedUserID(createdUserIDs []uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`created_user_id` IN (?)", createdUserIDs).Find(&results).Error

	return
}

// GetFromPayStatus 通过pay_status获取内容 支付状态:[1未支付, 2已取消或超时, 3支付完成, 4已退款]
func (obj *_ProductOrderMgr) GetFromPayStatus(payStatus int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`pay_status` = ?", payStatus).Find(&results).Error

	return
}

// GetBatchFromPayStatus 批量查找 支付状态:[1未支付, 2已取消或超时, 3支付完成, 4已退款]
func (obj *_ProductOrderMgr) GetBatchFromPayStatus(payStatuss []int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`pay_status` IN (?)", payStatuss).Find(&results).Error

	return
}

// GetFromTransactionID 通过transaction_id获取内容 微信支付单号
func (obj *_ProductOrderMgr) GetFromTransactionID(transactionID string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`transaction_id` = ?", transactionID).Find(&results).Error

	return
}

// GetBatchFromTransactionID 批量查找 微信支付单号
func (obj *_ProductOrderMgr) GetBatchFromTransactionID(transactionIDs []string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`transaction_id` IN (?)", transactionIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 下单时间
func (obj *_ProductOrderMgr) GetFromCreatedAt(createdAt int64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 下单时间
func (obj *_ProductOrderMgr) GetBatchFromCreatedAt(createdAts []int64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromDeliveryStatus 通过delivery_status获取内容 配送状态:[1未接单, 2已接单/已发货, 3已送达, 4已取消]
func (obj *_ProductOrderMgr) GetFromDeliveryStatus(deliveryStatus int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_status` = ?", deliveryStatus).Find(&results).Error

	return
}

// GetBatchFromDeliveryStatus 批量查找 配送状态:[1未接单, 2已接单/已发货, 3已送达, 4已取消]
func (obj *_ProductOrderMgr) GetBatchFromDeliveryStatus(deliveryStatuss []int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_status` IN (?)", deliveryStatuss).Find(&results).Error

	return
}

// GetFromAddressID 通过address_id获取内容 配送地址ID
func (obj *_ProductOrderMgr) GetFromAddressID(addressID uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`address_id` = ?", addressID).Find(&results).Error

	return
}

// GetBatchFromAddressID 批量查找 配送地址ID
func (obj *_ProductOrderMgr) GetBatchFromAddressID(addressIDs []uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`address_id` IN (?)", addressIDs).Find(&results).Error

	return
}

// GetFromDeliveryUserID 通过delivery_user_id获取内容 配送人ID/发货人ID
func (obj *_ProductOrderMgr) GetFromDeliveryUserID(deliveryUserID uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_user_id` = ?", deliveryUserID).Find(&results).Error

	return
}

// GetBatchFromDeliveryUserID 批量查找 配送人ID/发货人ID
func (obj *_ProductOrderMgr) GetBatchFromDeliveryUserID(deliveryUserIDs []uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_user_id` IN (?)", deliveryUserIDs).Find(&results).Error

	return
}

// GetFromDeliveryStatAt 通过delivery_stat_at获取内容 接单时间
func (obj *_ProductOrderMgr) GetFromDeliveryStatAt(deliveryStatAt int64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_stat_at` = ?", deliveryStatAt).Find(&results).Error

	return
}

// GetBatchFromDeliveryStatAt 批量查找 接单时间
func (obj *_ProductOrderMgr) GetBatchFromDeliveryStatAt(deliveryStatAts []int64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_stat_at` IN (?)", deliveryStatAts).Find(&results).Error

	return
}

// GetFromDeliveryEndAt 通过delivery_end_at获取内容 送达时间
func (obj *_ProductOrderMgr) GetFromDeliveryEndAt(deliveryEndAt int64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_end_at` = ?", deliveryEndAt).Find(&results).Error

	return
}

// GetBatchFromDeliveryEndAt 批量查找 送达时间
func (obj *_ProductOrderMgr) GetBatchFromDeliveryEndAt(deliveryEndAts []int64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_end_at` IN (?)", deliveryEndAts).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容
func (obj *_ProductOrderMgr) GetFromOpenID(openID string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`open_id` = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量查找
func (obj *_ProductOrderMgr) GetBatchFromOpenID(openIDs []string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`open_id` IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromExpressSn 通过express_sn获取内容 快递单号
func (obj *_ProductOrderMgr) GetFromExpressSn(expressSn string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`express_sn` = ?", expressSn).Find(&results).Error

	return
}

// GetBatchFromExpressSn 批量查找 快递单号
func (obj *_ProductOrderMgr) GetBatchFromExpressSn(expressSns []string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`express_sn` IN (?)", expressSns).Find(&results).Error

	return
}

// GetFromDeliveryType 通过delivery_type获取内容 配送类型:[1同城配送, 2快递配送]
func (obj *_ProductOrderMgr) GetFromDeliveryType(deliveryType int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_type` = ?", deliveryType).Find(&results).Error

	return
}

// GetBatchFromDeliveryType 批量查找 配送类型:[1同城配送, 2快递配送]
func (obj *_ProductOrderMgr) GetBatchFromDeliveryType(deliveryTypes []int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`delivery_type` IN (?)", deliveryTypes).Find(&results).Error

	return
}

// GetFromRefundStatus 通过refund_status获取内容 退款状态:[1 正常, 2申请退款, 3驳回退款, 4同意退款]
func (obj *_ProductOrderMgr) GetFromRefundStatus(refundStatus int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_status` = ?", refundStatus).Find(&results).Error

	return
}

// GetBatchFromRefundStatus 批量查找 退款状态:[1 正常, 2申请退款, 3驳回退款, 4同意退款]
func (obj *_ProductOrderMgr) GetBatchFromRefundStatus(refundStatuss []int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_status` IN (?)", refundStatuss).Find(&results).Error

	return
}

// GetFromRefundReason 通过refund_reason获取内容 退款原因
func (obj *_ProductOrderMgr) GetFromRefundReason(refundReason string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_reason` = ?", refundReason).Find(&results).Error

	return
}

// GetBatchFromRefundReason 批量查找 退款原因
func (obj *_ProductOrderMgr) GetBatchFromRefundReason(refundReasons []string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_reason` IN (?)", refundReasons).Find(&results).Error

	return
}

// GetFromRefundExamineUser 通过refund_examine_user获取内容 退款操作人ID
func (obj *_ProductOrderMgr) GetFromRefundExamineUser(refundExamineUser uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_examine_user` = ?", refundExamineUser).Find(&results).Error

	return
}

// GetBatchFromRefundExamineUser 批量查找 退款操作人ID
func (obj *_ProductOrderMgr) GetBatchFromRefundExamineUser(refundExamineUsers []uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_examine_user` IN (?)", refundExamineUsers).Find(&results).Error

	return
}

// GetFromRefundDetail 通过refund_detail获取内容 退款详情
func (obj *_ProductOrderMgr) GetFromRefundDetail(refundDetail string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_detail` = ?", refundDetail).Find(&results).Error

	return
}

// GetBatchFromRefundDetail 批量查找 退款详情
func (obj *_ProductOrderMgr) GetBatchFromRefundDetail(refundDetails []string) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`refund_detail` IN (?)", refundDetails).Find(&results).Error

	return
}

// GetFromReferrerID 通过referrer_id获取内容 推荐人ID
func (obj *_ProductOrderMgr) GetFromReferrerID(referrerID uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`referrer_id` = ?", referrerID).Find(&results).Error

	return
}

// GetBatchFromReferrerID 批量查找 推荐人ID
func (obj *_ProductOrderMgr) GetBatchFromReferrerID(referrerIDs []uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`referrer_id` IN (?)", referrerIDs).Find(&results).Error

	return
}

// GetFromProfitsharingStatus 通过profitsharing_status获取内容 分账状态:[1 未分账, 2 已分帐]
func (obj *_ProductOrderMgr) GetFromProfitsharingStatus(profitsharingStatus int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`profitsharing_status` = ?", profitsharingStatus).Find(&results).Error

	return
}

// GetBatchFromProfitsharingStatus 批量查找 分账状态:[1 未分账, 2 已分帐]
func (obj *_ProductOrderMgr) GetBatchFromProfitsharingStatus(profitsharingStatuss []int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`profitsharing_status` IN (?)", profitsharingStatuss).Find(&results).Error

	return
}

// GetFromProfitsharingMoney 通过profitsharing_money获取内容 分账金额(商品分账金 * 购买数量)
func (obj *_ProductOrderMgr) GetFromProfitsharingMoney(profitsharingMoney float64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`profitsharing_money` = ?", profitsharingMoney).Find(&results).Error

	return
}

// GetBatchFromProfitsharingMoney 批量查找 分账金额(商品分账金 * 购买数量)
func (obj *_ProductOrderMgr) GetBatchFromProfitsharingMoney(profitsharingMoneys []float64) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`profitsharing_money` IN (?)", profitsharingMoneys).Find(&results).Error

	return
}

// GetFromOrderType 通过order_type获取内容 订单类型:[1在线支付, 2积分支付]
func (obj *_ProductOrderMgr) GetFromOrderType(orderType int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`order_type` = ?", orderType).Find(&results).Error

	return
}

// GetBatchFromOrderType 批量查找 订单类型:[1在线支付, 2积分支付]
func (obj *_ProductOrderMgr) GetBatchFromOrderType(orderTypes []int) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`order_type` IN (?)", orderTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductOrderMgr) FetchByPrimaryKey(id uint32) (result ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByID primary or index 获取唯一内容
func (obj *_ProductOrderMgr) FetchUniqueByID(id uint32) (result ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByOrderSn primary or index 获取唯一内容
func (obj *_ProductOrderMgr) FetchUniqueByOrderSn(orderSn string) (result ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`order_sn` = ?", orderSn).First(&result).Error

	return
}

// FetchIndexByProductOrderProductIDFk  获取多个内容
func (obj *_ProductOrderMgr) FetchIndexByProductOrderProductIDFk(productID uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`product_id` = ?", productID).Find(&results).Error

	return
}

// FetchIndexByProductOrderAddressIDFk  获取多个内容
func (obj *_ProductOrderMgr) FetchIndexByProductOrderAddressIDFk(addressID uint32) (results []*ProductOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductOrder{}).Where("`address_id` = ?", addressID).Find(&results).Error

	return
}
