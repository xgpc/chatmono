/**
 * @Author: smono
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2022/9/27 16:38
 */

package fileServer

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/util/guzzle"
	"net/http"
)

type UploadController struct {
	Base *frame.Base
	Ctx  iris.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var uploadClient *guzzle.Client

func (this *UploadController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodPost, "/oss/isExist", "OssExist")
}

// PostConfig 获取上传文件配置信息
// @Summary      获取上传文件配置信息
// @Description  获取上传文件配置信息
// @Produce      json
// @Tags         上传文件
// @param        token  header    string  true  "登录用户token"
// @Success      200    {object}  Response
// @Router       /api/server/config [post]
func (this *UploadController) PostConfig() {
	res := map[string]interface{}{
		"ossPath":  "",
		"filePath": `https://cs.smono.com`,
	}

	this.Base.SuccessWithData(res)
}

type PostOssUpload struct {
	UploadType  string `valid:"required"json:"upload_type"` // 文件上传模式[oss, local]
	ProjectName string `valid:"required" json:"project_name"`
	FileType    string `valid:"required" json:"file_type"`
	FileSize    int64  `valid:"required" json:"file_size"`
	FileData    []byte `valid:"-" json:"file_data"` // 上传模式为local时, 需要填入文件数据
}

// PostOssUpload 上传
// @Summary      上传
// @Description  上传
// @Produce      json
// @Tags         上传文件
// @param        token  header    string         true  "登录用户token"
// @param        root   body      PostOssUpload  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/upload [post]
func (this *UploadController) PostOssUpload() {
	param := &PostOssUpload{}
	this.Base.Init(param)

	res := "Hand.Upload()"

	this.Base.SuccessWithData(res)
}

type PostOssConfirm struct {
	Key        string `valid:"required" json:"key"`
	ProductKey string `valid:"required" json:"product_key"`
}

// PostOssConfirm 上传确认
// @Summary      上传确认
// @Description  上传确认
// @Produce      json
// @Tags         上传文件
// @param        token  header    string          true  "登录用户token"
// @param        root   body      PostOssConfirm  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/confirm [post]
func (this *UploadController) PostOssConfirm() {
	param := &PostOssConfirm{}
	this.Base.Init(param)

	resp, err := uploadClient.RequestJSON(map[string]interface{}{
		"Key":        param.Key,
		"ProductKey": param.ProductKey,
	}).Post("/server/oss/confirm")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Base.SuccessWithData(res.Data)
}

type PostOssKey struct {
	Key string `valid:"required" json:"Key"`
}

// PostOssDelete 删除
// @Summary      删除
// @Description  删除
// @Produce      json
// @Tags         上传文件
// @param        token  header    string      true  "登录用户token"
// @param        root   body      PostOssKey  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/delete [post]
func (this *UploadController) PostOssDelete() {
	param := &PostOssKey{}
	this.Base.Init(param)

	resp, err := uploadClient.RequestJSON(map[string]interface{}{
		"Key":    param.Key,
		"UserID": this.Base.MyId(),
	}).Post("/server/oss/delete")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Base.SuccessWithData(res.Data)
}

// OssExist 查看是否存在
// @Summary      查看是否存在
// @Description  查看是否存在
// @Produce      json
// @Tags         上传文件
// @param        token  header    string      true  "登录用户token"
// @param        root   body      PostOssKey  true  "上传参数"
// @Success      200    {object}  Response
// @Router       /api/server/oss/isExist [post]
func (this *UploadController) OssExist() {
	param := &PostOssKey{}
	this.Base.Init(param)

	resp, err := uploadClient.RequestJSON(map[string]interface{}{
		"Key": param.Key,
	}).Post("/server/oss/isExist")
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := Response{}
	resp.JSON(&res)
	if res.Code != 0 {
		exce.ThrowSys(exce.CodeRequestError, res.Msg)
	}
	this.Base.SuccessWithData(res.Data)
}
