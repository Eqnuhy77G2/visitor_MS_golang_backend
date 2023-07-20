package logic

import (
	"VisitorsManagementSystem/dao"
	"VisitorsManagementSystem/models"
)

func SelectVisitorName(visitorName string) string {
	var visitor models.Visitor
	dao.InitDB().Where("visitor_name = ?",visitorName).Find(&visitor)
	return visitor.VisitorName
}

func FindVisitor(visitorName string) models.Visitor{
	var visitor models.Visitor
	dao.InitDB().Where("visitor_name = ?",visitorName).Find(&visitor)
	return visitor
}

func  SelectPhone(visitorName string) string {
	var visitor models.Visitor
	dao.InitDB().Where("visitor_name = ?",visitorName).Find(&visitor)
	return visitor.Phone
}

func SelectVisitor(visitorName string,phone string) []models.Visitor {
	var visitor []models.Visitor
	dao.InitDB().Where("visitor_name = ? AND phone= ?",visitorName,phone).Find(&visitor)
	return visitor
}



