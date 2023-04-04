/**
 * @Author: smono
 * @Description:
 * @File:  isSuper
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:58
 */

package admin

import "github.com/xgpc/dsg/exce"

func IsSuper(userID uint32) bool {
	info := GetAdmin(userID)
	if info.Super == 1 {
		return true
	}
	return false
}

func CheckSuper(userID uint32) {
	if !IsSuper(userID) {
		exce.ThrowSys(exce.CodeUserNoAuth, "您不是超级管理员")
	}
}
