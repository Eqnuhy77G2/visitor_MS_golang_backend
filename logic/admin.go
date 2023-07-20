package logic

import (
	"VisitorsManagementSystem/dao"
	"VisitorsManagementSystem/models"
)

func AllData() []models.Visitor {
	var visitor []models.Visitor
	dao.InitDB().Find(&visitor)
	return visitor
}

func DeleteId(id int64) bool{
	var visitor models.Visitor
	dao.InitDB().Where("id = ?",id).Delete(&visitor)
	return true

}

func Update(visitorName string,phone string,visitId string)  bool{
	var visitor models.Visitor
	dao.InitDB().Model(&visitor).Where("visitor_name = ?",visitorName).Updates(map[string]interface{}{"phone":phone,"visit_id":visitId})
	return true
}


//func UpdatePost(data models.Post, id int64) (int64, error) {
//	var model models.Post
//	row := db.First(&model, id)
//	if row.Error == nil {
//		result := db.Model(&model).Updates(&data)
//		return model.ID, result.Error
//	}
//	return 0, row.Error
//}


