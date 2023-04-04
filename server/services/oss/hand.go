/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/28 19:07
 */

package oss

var Hand *AliOss

func Init(conf Config) {
	Hand = NewAliOss(conf)
}
