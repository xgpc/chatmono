/**
 * @Author: smono
 * @Description:
 * @File:  isAdmin
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:49
 */

package admin

import "github.com/xgpc/dsg/exce"

func IsAdmin(userID uint32) bool {
	info := GetAdmin(userID)
	if info.UserID != 0 {
		return true
	}
	return false
}

func CheckAdmin(userID uint32) {
	if !IsAdmin(userID) {
		exce.ThrowSys(exce.CodeUserNoAuth, "您不是管理员, 无权操作")
	}
}
