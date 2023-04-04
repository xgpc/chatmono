package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SessionTemplatesMgr struct {
	*_BaseMgr
}

// SessionTemplatesMgr open func
func SessionTemplatesMgr(db *gorm.DB) *_SessionTemplatesMgr {
	if db == nil {
		panic(fmt.Errorf("SessionTemplatesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SessionTemplatesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("session_templates"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SessionTemplatesMgr) GetTableName() string {
	return "session_templates"
}

// Reset 重置gorm会话
func (obj *_SessionTemplatesMgr) Reset() *_SessionTemplatesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SessionTemplatesMgr) Get() (result SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SessionTemplatesMgr) Gets() (results []*SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SessionTemplatesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SessionTemplatesMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取
func (obj *_SessionTemplatesMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithBody body获取
func (obj *_SessionTemplatesMgr) WithBody(body string) Option {
	return optionFunc(func(o *options) { o.query["body"] = body })
}

// GetByOption 功能选项模式获取
func (obj *_SessionTemplatesMgr) GetByOption(opts ...Option) (result SessionTemplates, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SessionTemplatesMgr) GetByOptions(opts ...Option) (results []*SessionTemplates, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SessionTemplatesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SessionTemplates, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where(options.query)
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
func (obj *_SessionTemplatesMgr) GetFromID(id uint32) (result SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SessionTemplatesMgr) GetBatchFromID(ids []uint32) (results []*SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_SessionTemplatesMgr) GetFromTitle(title string) (results []*SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找
func (obj *_SessionTemplatesMgr) GetBatchFromTitle(titles []string) (results []*SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromBody 通过body获取内容
func (obj *_SessionTemplatesMgr) GetFromBody(body string) (results []*SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`body` = ?", body).Find(&results).Error

	return
}

// GetBatchFromBody 批量查找
func (obj *_SessionTemplatesMgr) GetBatchFromBody(bodys []string) (results []*SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`body` IN (?)", bodys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SessionTemplatesMgr) FetchByPrimaryKey(id uint32) (result SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByID primary or index 获取唯一内容
func (obj *_SessionTemplatesMgr) FetchUniqueByID(id uint32) (result SessionTemplates, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SessionTemplates{}).Where("`id` = ?", id).First(&result).Error

	return
}
