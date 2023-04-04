/**
 * @Author: smono
 * @Description:
 * @File:  oss_test.go
 * @Version: 1.0.0
 * @Date: 2022/9/28 17:14
 */

package oss

import (
	"github.com/xgpc/dsg/util/conf"
	"testing"
)

func TestOss(t *testing.T) {
	conf2 := Config{}
	conf.LoadConf(&conf2)

	//aliOss := NewAliOss(conf2)

	//aliOss.Upload("")

}
