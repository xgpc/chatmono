package models

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _DictMgr struct {
	*_BaseMgr
}

// DictMgr open func
func DictMgr(db *gorm.DB) *_DictMgr {
	if db == nil {
		panic(fmt.Errorf("DictMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictMgr) GetTableName() string {
	return "dict"
}

// Reset 重置gorm会话
func (obj *_DictMgr) Reset() *_DictMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictMgr) Get() (result Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictMgr) Gets() (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Dict{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithLabel label获取
func (obj *_DictMgr) WithLabel(label string) Option {
	return optionFunc(func(o *options) { o.query["label"] = label })
}

// WithDictType dict_type获取
func (obj *_DictMgr) WithDictType(dictType string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithValueData value_data获取
func (obj *_DictMgr) WithValueData(valueData datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["value_data"] = valueData })
}

// WithCreatedUserID created_user_id获取
func (obj *_DictMgr) WithCreatedUserID(createdUserID uint32) Option {
	return optionFunc(func(o *options) { o.query["created_user_id"] = createdUserID })
}

// WithParentID parent_id获取
func (obj *_DictMgr) WithParentID(parentID uint32) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// GetByOption 功能选项模式获取
func (obj *_DictMgr) GetByOption(opts ...Option) (result Dict, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictMgr) GetByOptions(opts ...Option) (results []*Dict, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_DictMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Dict, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Dict{}).Where(options.query)
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
func (obj *_DictMgr) GetFromID(id uint32) (result Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictMgr) GetBatchFromID(ids []uint32) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromLabel 通过label获取内容
func (obj *_DictMgr) GetFromLabel(label string) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`label` = ?", label).Find(&results).Error

	return
}

// GetBatchFromLabel 批量查找
func (obj *_DictMgr) GetBatchFromLabel(labels []string) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`label` IN (?)", labels).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容
func (obj *_DictMgr) GetFromDictType(dictType string) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`dict_type` = ?", dictType).Find(&results).Error

	return
}

// GetBatchFromDictType 批量查找
func (obj *_DictMgr) GetBatchFromDictType(dictTypes []string) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`dict_type` IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromValueData 通过value_data获取内容
func (obj *_DictMgr) GetFromValueData(valueData datatypes.JSON) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`value_data` = ?", valueData).Find(&results).Error

	return
}

// GetBatchFromValueData 批量查找
func (obj *_DictMgr) GetBatchFromValueData(valueDatas []datatypes.JSON) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`value_data` IN (?)", valueDatas).Find(&results).Error

	return
}

// GetFromCreatedUserID 通过created_user_id获取内容
func (obj *_DictMgr) GetFromCreatedUserID(createdUserID uint32) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`created_user_id` = ?", createdUserID).Find(&results).Error

	return
}

// GetBatchFromCreatedUserID 批量查找
func (obj *_DictMgr) GetBatchFromCreatedUserID(createdUserIDs []uint32) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`created_user_id` IN (?)", createdUserIDs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_DictMgr) GetFromParentID(parentID uint32) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找
func (obj *_DictMgr) GetBatchFromParentID(parentIDs []uint32) (results []*Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictMgr) FetchByPrimaryKey(id uint32) (result Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByID primary or index 获取唯一内容
func (obj *_DictMgr) FetchUniqueByID(id uint32) (result Dict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Dict{}).Where("`id` = ?", id).First(&result).Error

	return
}
