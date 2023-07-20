package logic

import (
	"VisitorsManagementSystem/dao"
	"VisitorsManagementSystem/models"
)

func SelectPassword(adminName string) string{
	var admin models.Admin
	dao.InitDB().Where("username = ?",adminName).Find(&admin)
	return admin.Password
}

func SelectAdmin(adminName string) string {
	var admin models.Admin
	dao.InitDB().Where("username = ?",adminName).Find(&admin)
	return admin.Username

}
