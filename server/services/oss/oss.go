// Package oss
// @Author:        asus
// @Description:   $
// @File:          oss
// @Data:          2021/12/2114:09
//
package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/util"
	"hash"
	"io"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"`
}

type AliOss struct {
	conf   Config
	bucket *oss.Bucket
}

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
	expireTime  = 60
	uploadDir   = "tmpFiles/"
)

func NewAliOss(conf Config) *AliOss {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	// endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	// accessKeyId := "<yourAccessKeyId>"
	// accessKeySecret := "<yourAccessKeySecret>"
	// bucketName := "<yourBucketName>"

	// 创建OSSClient实例。
	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	// 创建存储空间。

	b, err := client.Bucket(conf.BucketName)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	var info AliOss
	info.bucket = b
	info.conf = conf
	return &info
}

type StorageFileRes struct {
	Name     string
	Url      string
	FormData StorageOssFormData
}

type StorageOssFormData struct {
	Key                 string `json:"key"`
	Policy              string `json:"policy"`
	OSSAccessKeyId      string `json:"OSSAccessKeyId"`
	Signature           string `json:"signature"`
	SuccessActionStatus string `json:"success_action_status"`
}

type CfgData struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

func (this *AliOss) getGmtIso8601(expireEnd int64) string {
	return time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
}

func (this *AliOss) getDownloadUrl() string {
	return "https://" + this.conf.BucketName + "." + this.conf.Endpoint
}

func (this *AliOss) getPolicyToken(proName, fileName string, mid uint32) StorageFileRes {
	uploadDir := "Files/" + proName
	userID := strconv.FormatUint(uint64(mid), 10)
	now := time.Now().Unix()
	expireEnd := now + expireTime
	var tokenExpire = this.getGmtIso8601(expireEnd)

	var cfgData CfgData
	cfgData.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadDir)
	cfgData.Conditions = append(cfgData.Conditions, condition)

	result, _ := json.Marshal(cfgData)
	deByte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(this.conf.AccessKeySecret))
	_, _ = io.WriteString(h, deByte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return StorageFileRes{
		Name: "file",
		Url:  this.getDownloadUrl(),
		FormData: StorageOssFormData{
			Key:                 uploadDir + "/" + time.Now().Format("2006_01_02") + "/" + userID + "/" + fileName,
			Policy:              deByte,
			OSSAccessKeyId:      this.conf.AccessKeyID,
			Signature:           signedStr,
			SuccessActionStatus: "200",
		},
	}
}

// Upload 上传文件
func (this *AliOss) Upload(ext string, fileSize int64, proName string, mid uint32) map[string]interface{} {

	fileExt := strings.ToLower(ext)

	id := util.Uuid()
	fileName := id + "." + fileExt
	return map[string]interface{}{
		"project": proName,
		"name":    fileName,
		"ext":     fileExt,
		"size":    fileSize,
		"id":      id,
		"oss":     this.getPolicyToken(proName, fileName, mid),
	}
}

// Move 移动文件
func (this *AliOss) Move(objectName, destObjectName string) *oss.CopyObjectResult {
	res, err := this.bucket.CopyObject(objectName, destObjectName)
	if err != nil {

		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	return &res
}

// Delete 删除文件
func (this *AliOss) Delete(objectName string) {
	err := this.bucket.DeleteObject(objectName)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
}

// IsObjectExist 判断文件是否存在。
func (this *AliOss) IsObjectExist(objectName string) bool {
	isExist, _ := this.bucket.IsObjectExist(objectName)
	return isExist
}
