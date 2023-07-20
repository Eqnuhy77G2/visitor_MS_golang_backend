package models
import "github.com/jinzhu/gorm"

type Visitor struct {

	gorm.Model
	//Id uint `gorm:"primary key"`
	VisitorName string `gorm:"varchar(20);not null" form:"visitor_name" json:"visitor_name"`
	Sex string `gorm:"enum('男','女');not null" form:"sex" json:"sex"`
	Phone string `gorm:"varchar(20);not null" form:"phone" json:"phone"`
	VisitId string `gorm:"varchar(20);not null" form:"visit_id" json:"visit_id"`
	Event string `gorm:"varchar(100);not null" form:"event" json:"event"`
	//VisitTime time.Time
}


type Admin struct {
	gorm.Model
	Username      string `gorm:"varchar(20);not null" form:"username" json:"username" `
	Password  string `gorm:"varchar(20);not null" form:"password" json:"password" `
}
