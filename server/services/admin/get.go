/**
 * @Author: smono
 * @Description:
 * @File:  get
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:49
 */

package admin

import "chatmono/models"

func GetAdmin(userID uint32) models.Admin {
	mgr := models.AdminMgr(db())
	info, _ := mgr.GetFromUserID(userID)
	return info
}
