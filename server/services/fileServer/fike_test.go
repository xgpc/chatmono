/**
 * @Author: smono
 * @Description:
 * @File:  fike_test.go
 * @Version: 1.0.0
 * @Date: 2022/9/28 22:45
 */

package fileServer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMove(t *testing.T) {
	err := os.Rename("tempfile/1/1.txt", "tempfile/2/1.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDel(t *testing.T) {
	err := os.Remove("tempfile/2/1.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFile(t *testing.T) {
	fileName := "temp.file/2/1.txt"

	fmt.Println(strings.LastIndex(fileName, "."))

	fmt.Println(fileName[0:13] + "(1)" + fileName[13:])
}

func TestDir(t *testing.T) {
	// 获取文件目录列表
	f, err := ioutil.ReadDir("../")
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range f {
		fmt.Println(v.Name(), "- isDir", v.IsDir())
	}

}

func TestUpDir(t *testing.T) {
	fmt.Println(strings.Index("../asdas", ".."))
	fmt.Println(strings.Index("./asdas", ".."))
}
