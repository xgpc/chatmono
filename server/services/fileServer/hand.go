/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/27 13:39
 */

package fileServer

import (
	"chatmono/services/fileServer/fileBase"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/util"
	"os"
	"path"
	"strings"
)

var Hand *fileBase.FileBase

type Config struct {
	UploadRoot string `yaml:"upload_root"` // 存放文件跟路劲
	UploadDir  string `yaml:"upload_dir"`  // 跟路劲下的文件夹    存放地址 = root + dir    访问地址 = dir + "xxx"
}

func Init(conf Config) {
	Hand = fileBase.NewFileBase(conf.UploadRoot, conf.UploadDir)
}

type uploadReq struct {
}

// upload   文件上传
// @Summary 文件上传
// @Description 文件上传
// @Accept       json
// @Produce      json
// @param        root  body  uploadReq  true  "参数"
// @Tags    文件管理
// @Success      200  {object}  render.Response
// @Router       /fileServer/upload [Post]
func upload(ctx iris.Context) {
	p := frame.NewBase(ctx)

	// 文件上传没有json 数据, 读取反而会导致错误
	//var param uploadReq
	//p.Init(&param)

	// TODO: 设置文件最大上传小大
	// ctx.SetMaxRequestBodySize(maxSize)
	// OR
	// app.Use(iris.LimitRequestBodySize(maxSize))
	// OR
	// OR iris.WithPostMaxMemory(maxSize)

	// single file
	f, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	defer f.Close()

	// fileDir + userID + ymd + file + .xxx
	// 如果存在
	// fileDir + userID + ymd + file + (1) + .xxx

	// Upload the file to specific destination.

	filePath := path.Join(Hand.UploadDir,
		p.MyIdToString(), util.TimeYmd_Now2())
	dest := path.Join(Hand.UploadRoot, filePath, fileHeader.Filename)

	if Hand.IsObjectExist(dest) {
		dest = upFileName(dest)
	}

	// 判读目录是否存在
	destPath := path.Join(Hand.UploadRoot, filePath)
	if !Hand.IsObjectExist(destPath) {
		os.MkdirAll(destPath, os.ModePerm)
	}

	_, err = ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	res := map[string]interface{}{}
	res["filePath"] = path.Join(filePath, fileHeader.Filename)
	p.SuccessWithData(res)

}

func upFileName(fileName string) string {
	tag := strings.LastIndex(fileName, ".")
	newFileName := fileName[0:tag] + "(1)" + fileName[tag:]
	if Hand.IsObjectExist(newFileName) {
		newFileName = upFileName(newFileName)
	}

	return newFileName
}

type reqExist struct {
	FilePath string `json:"file_path"`
}

// Exist        查询文件是否存在
// @Summary     查询文件是否存在
// @Description 查询文件是否存在
// @Accept       json
// @Produce      json
// @param        root  body  reqExist  true  "参数"
// @Tags    文件管理
// @Success      200  {object}  render.Response
// @Router       /fileServer/exist [Post]
func Exist(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param reqExist
	p.Init(&param)

	res := map[string]interface{}{}

	res["is_exist"] = Hand.IsObjectExist(param.FilePath)

	p.SuccessWithData(res)
}

type reqMove struct {
	OldFilePath string `json:"old_file_path"`
	NewFilePath string `json:"new_file_path"`
}

// Move         文件移动
// @Summary     文件移动
// @Description 文件移动
// @Accept       json
// @Produce      json
// @param        root  body  reqMove  true  "参数"
// @Tags    文件管理
// @Success      200  {object}  render.Response
// @Router       /fileServer/move [Post]
func Move(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param reqMove
	p.Init(&param)

	err := Hand.Move(param.OldFilePath, param.NewFilePath)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.Success()
}

type reqDel struct {
	FilePath string `json:"file_path"`
}

// Del          文件删除
// @Summary     文件删除
// @Description 文件删除
// @Accept       json
// @Produce      json
// @param        root  body  reqDel  true  "参数"
// @Tags    文件管理
// @Success      200  {object}  render.Response
// @Router       /fileServer/del [Post]
func Del(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param reqDel
	p.Init(&param)

	err := Hand.Delete(param.FilePath)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	p.Success()
}

type queryReq struct {
	Path string `json:"path"` // 文件夹路劲
}

// Query 查看目录
// @Summary      查看目录
// @Description  查看目录
// @Accept       json
// @Produce      json
// @param        root  body  queryReq  true  "参数"
// @Tags         文件管理
// @Success      200  {object}  render.Response
// @Router       /fileServer/query [Post]
func Query(ctx iris.Context) {
	p := frame.NewBase(ctx)
	var param queryReq
	p.Init(&param)

	resList := Hand.Query(param.Path)

	p.SuccessWithList(resList, len(resList))
}
