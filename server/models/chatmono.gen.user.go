package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserName user_name获取
func (obj *_UserMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithMobileData mobile_data获取
func (obj *_UserMgr) WithMobileData(mobileData []byte) Option {
	return optionFunc(func(o *options) { o.query["mobile_data"] = mobileData })
}

// WithUserImg user_img获取 用户头像
func (obj *_UserMgr) WithUserImg(userImg string) Option {
	return optionFunc(func(o *options) { o.query["user_img"] = userImg })
}

// WithAreaCode area_code获取
func (obj *_UserMgr) WithAreaCode(areaCode string) Option {
	return optionFunc(func(o *options) { o.query["area_code"] = areaCode })
}

// WithCity city获取
func (obj *_UserMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithMobile mobile获取
func (obj *_UserMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithPassword password获取 密码
func (obj *_UserMgr) WithPassword(password []byte) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]User, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query)
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
func (obj *_UserMgr) GetFromID(id uint32) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UserMgr) GetBatchFromID(ids []uint32) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_UserMgr) GetFromUserName(userName string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`user_name` = ?", userName).First(&result).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_UserMgr) GetBatchFromUserName(userNames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromMobileData 通过mobile_data获取内容
func (obj *_UserMgr) GetFromMobileData(mobileData []byte) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`mobile_data` = ?", mobileData).Find(&results).Error

	return
}

// GetBatchFromMobileData 批量查找
func (obj *_UserMgr) GetBatchFromMobileData(mobileDatas [][]byte) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`mobile_data` IN (?)", mobileDatas).Find(&results).Error

	return
}

// GetFromUserImg 通过user_img获取内容 用户头像
func (obj *_UserMgr) GetFromUserImg(userImg string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`user_img` = ?", userImg).Find(&results).Error

	return
}

// GetBatchFromUserImg 批量查找 用户头像
func (obj *_UserMgr) GetBatchFromUserImg(userImgs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`user_img` IN (?)", userImgs).Find(&results).Error

	return
}

// GetFromAreaCode 通过area_code获取内容
func (obj *_UserMgr) GetFromAreaCode(areaCode string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`area_code` = ?", areaCode).Find(&results).Error

	return
}

// GetBatchFromAreaCode 批量查找
func (obj *_UserMgr) GetBatchFromAreaCode(areaCodes []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`area_code` IN (?)", areaCodes).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容
func (obj *_UserMgr) GetFromCity(city string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找
func (obj *_UserMgr) GetBatchFromCity(citys []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容
func (obj *_UserMgr) GetFromMobile(mobile string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找
func (obj *_UserMgr) GetBatchFromMobile(mobiles []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_UserMgr) GetFromPassword(password []byte) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_UserMgr) GetBatchFromPassword(passwords [][]byte) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id uint32) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUserIDUIndex primary or index 获取唯一内容
func (obj *_UserMgr) FetchUniqueByUserIDUIndex(id uint32) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUserName primary or index 获取唯一内容
func (obj *_UserMgr) FetchUniqueByUserName(userName string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`user_name` = ?", userName).First(&result).Error

	return
}
