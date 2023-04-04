/**
 * @Author: smono
 * @Description:
 * @File:  session
 * @Version: 1.0.0
 * @Date: 2022/9/24 16:16
 */

package user

import (
	"chatmono/services/user/session"
	"github.com/xgpc/dsg/exce"
	"strconv"
)

func SetSession(userID uint32) string {
	token := session.CreatedToKen()
	//idStr := strconv.Itoa(int(userID))
	session.Set(token, userID)
	return token
}

func GetSession(key string) uint32 {
	data := session.Get(key)
	idStr := data.(string)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	return uint32(id)
}

func ExpireToken(key string) {
	session.Expire(key)
}
