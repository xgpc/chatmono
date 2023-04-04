package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _WxUserMgr struct {
	*_BaseMgr
}

// WxUserMgr open func
func WxUserMgr(db *gorm.DB) *_WxUserMgr {
	if db == nil {
		panic(fmt.Errorf("WxUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxUserMgr) GetTableName() string {
	return "wx_user"
}

// Reset 重置gorm会话
func (obj *_WxUserMgr) Reset() *_WxUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxUserMgr) Get() (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).First(&result).Error
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
func (obj *_WxUserMgr) Gets() (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Find(&results).Error
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
func (obj *_WxUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxUserMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_WxUserMgr) WithUserID(userID uint32) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithWxAppid wx_appid获取
func (obj *_WxUserMgr) WithWxAppid(wxAppid string) Option {
	return optionFunc(func(o *options) { o.query["wx_appid"] = wxAppid })
}

// WithOpenID open_id获取
func (obj *_WxUserMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithIsProfitsharing is_profitsharing获取 是否开通分账:[1否, 2是]
func (obj *_WxUserMgr) WithIsProfitsharing(isProfitsharing int) Option {
	return optionFunc(func(o *options) { o.query["is_profitsharing"] = isProfitsharing })
}

// GetByOption 功能选项模式获取
func (obj *_WxUserMgr) GetByOption(opts ...Option) (result WxUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where(options.query).First(&result).Error
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
func (obj *_WxUserMgr) GetByOptions(opts ...Option) (results []*WxUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where(options.query).Find(&results).Error
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
func (obj *_WxUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WxUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_WxUserMgr) GetFromID(id uint32) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_WxUserMgr) GetBatchFromID(ids []uint32) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`id` IN (?)", ids).Find(&results).Error
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

// GetFromUserID 通过user_id获取内容
func (obj *_WxUserMgr) GetFromUserID(userID uint32) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`user_id` = ?", userID).First(&result).Error
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
func (obj *_WxUserMgr) GetBatchFromUserID(userIDs []uint32) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
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

// GetFromWxAppid 通过wx_appid获取内容
func (obj *_WxUserMgr) GetFromWxAppid(wxAppid string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`wx_appid` = ?", wxAppid).Find(&results).Error
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

// GetBatchFromWxAppid 批量查找
func (obj *_WxUserMgr) GetBatchFromWxAppid(wxAppids []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`wx_appid` IN (?)", wxAppids).Find(&results).Error
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

// GetFromOpenID 通过open_id获取内容
func (obj *_WxUserMgr) GetFromOpenID(openID string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`open_id` = ?", openID).Find(&results).Error
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

// GetBatchFromOpenID 批量查找
func (obj *_WxUserMgr) GetBatchFromOpenID(openIDs []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`open_id` IN (?)", openIDs).Find(&results).Error
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

// GetFromIsProfitsharing 通过is_profitsharing获取内容 是否开通分账:[1否, 2是]
func (obj *_WxUserMgr) GetFromIsProfitsharing(isProfitsharing int) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`is_profitsharing` = ?", isProfitsharing).Find(&results).Error
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

// GetBatchFromIsProfitsharing 批量查找 是否开通分账:[1否, 2是]
func (obj *_WxUserMgr) GetBatchFromIsProfitsharing(isProfitsharings []int) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`is_profitsharing` IN (?)", isProfitsharings).Find(&results).Error
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
func (obj *_WxUserMgr) FetchByPrimaryKey(id uint32) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByWxUserIDUIndex primary or index 获取唯一内容
func (obj *_WxUserMgr) FetchUniqueByWxUserIDUIndex(id uint32) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByUserID primary or index 获取唯一内容
func (obj *_WxUserMgr) FetchUniqueByUserID(userID uint32) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`user_id` = ?", userID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.UserID).Find(&result.User).Error; err != nil { // 用户表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
