package models

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _AdminMgr struct {
	*_BaseMgr
}

// AdminMgr open func
func AdminMgr(db *gorm.DB) *_AdminMgr {
	if db == nil {
		panic(fmt.Errorf("AdminMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdminMgr{_BaseMgr: &_BaseMgr{DB: db.Table("admin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdminMgr) GetTableName() string {
	return "admin"
}

// Reset 重置gorm会话
func (obj *_AdminMgr) Reset() *_AdminMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AdminMgr) Get() (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_AdminMgr) Gets() (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AdminMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Admin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_AdminMgr) WithUserID(userID uint32) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithSuper super获取 超级管理员:[0 普通, 1 超管]
func (obj *_AdminMgr) WithSuper(super int) Option {
	return optionFunc(func(o *options) { o.query["super"] = super })
}

// WithRule rule获取 权限规则
func (obj *_AdminMgr) WithRule(rule datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["rule"] = rule })
}

// GetByOption 功能选项模式获取
func (obj *_AdminMgr) GetByOption(opts ...Option) (result Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AdminMgr) GetByOptions(opts ...Option) (results []*Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_AdminMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Admin, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUserID 通过user_id获取内容
func (obj *_AdminMgr) GetFromUserID(userID uint32) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`user_id` = ?", userID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AdminMgr) GetBatchFromUserID(userIDs []uint32) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSuper 通过super获取内容 超级管理员:[0 普通, 1 超管]
func (obj *_AdminMgr) GetFromSuper(super int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`super` = ?", super).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSuper 批量查找 超级管理员:[0 普通, 1 超管]
func (obj *_AdminMgr) GetBatchFromSuper(supers []int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`super` IN (?)", supers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRule 通过rule获取内容 权限规则
func (obj *_AdminMgr) GetFromRule(rule datatypes.JSON) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`rule` = ?", rule).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRule 批量查找 权限规则
func (obj *_AdminMgr) GetBatchFromRule(rules []datatypes.JSON) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`rule` IN (?)", rules).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UserID).Find(&results[i].User).Error; err != nil { // 用户表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AdminMgr) FetchByPrimaryKey(userID uint32) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`user_id` = ?", userID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByAdminUserIDUIndex primary or index 获取唯一内容
func (obj *_AdminMgr) FetchUniqueByAdminUserIDUIndex(userID uint32) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`user_id` = ?", userID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
