package service

import "sense100/model"


/*
根据id检测用户是否存在
*/
func CheckAdminExist(id int64) bool {
	var admin model.SsAdmin
	err := db.Find(&admin, id).Error
	if err!=nil || admin.ID == 0 {
		return false
	}
	return true
}
