/**
 * @Author: smono
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2022/9/27 13:39
 */

package fileBase

import (
	"github.com/xgpc/dsg/exce"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type FileBase struct {
	UploadRoot string `yaml:"upload_root"` // 存放文件跟路劲
	UploadDir  string `yaml:"upload_dir"`  // 跟路劲下的文件夹    存放地址 = root + dir    访问地址 = dir + "xxx"
}

func NewFileBase(UploadRoot, UploadDir string) *FileBase {
	return &FileBase{
		UploadRoot: UploadRoot,
		UploadDir:  UploadDir,
	}
}

func (p *FileBase) Move(objectName string, destObjectName string) error {

	oldPath := path.Join(p.UploadRoot, objectName)
	newPath := path.Join(p.UploadRoot, destObjectName)

	return os.Rename(oldPath, newPath)
}

func (p *FileBase) Delete(objectName string) error {

	pathName := path.Join(p.UploadRoot, objectName)
	return os.Remove(pathName)
}

func (p *FileBase) IsObjectExist(objectName string) bool {
	pathName := path.Join(p.UploadRoot, objectName)
	return Exists(pathName)

}

type DirInfo struct {
	Name      string
	IsDir     bool
	Path      string
	CreatedAt int64
	Size      int64
}

func (p *FileBase) Query(FielPath string) []DirInfo {
	if -1 != strings.Index(FielPath, "..") {
		exce.ThrowSys(exce.CodeRequestError, "访问目录存在 ..  请检查后重试")
	}

	// 获取文件目录列表
	f, err := ioutil.ReadDir(path.Join(p.UploadRoot, FielPath))
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	var resList []DirInfo

	for _, v := range f {

		resList = append(resList, DirInfo{
			Name:      v.Name(),
			IsDir:     v.IsDir(),
			Path:      path.Join(FielPath, v.Name()),
			CreatedAt: v.ModTime().Unix(),
			Size:      v.Size(),
		})
	}

	return resList
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {

	_, err := os.Stat(path) //os.Stat获取文件信息

	if err != nil {

		if os.IsExist(err) {

			return true

		}

		return false

	}

	return true

}

// 判断所给路径是否为文件夹

func IsDir(path string) bool {

	s, err := os.Stat(path)

	if err != nil {

		return false

	}

	return s.IsDir()

}

// IsFile  判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
