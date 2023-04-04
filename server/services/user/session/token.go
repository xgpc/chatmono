/**
 * @Author: smono
 * @Description:
 * @File:  session
 * @Version: 1.0.0
 * @Date: 2022/9/23 0:16
 */

package session

import (
	"github.com/google/uuid"
	"github.com/xgpc/dsg/exce"
)

func CreatedToKen() string {
	uid, err := uuid.NewUUID()
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	return uid.String()
}
