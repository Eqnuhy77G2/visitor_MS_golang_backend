package controllers

import (
	"VisitorsManagementSystem/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Select(c *gin.Context) {
	var form Search

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if logic.SelectVisitorName(form.VisitorName) == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "User does not exist",
		})
		return
	}

	if logic.SelectPhone(form.VisitorName) != form.Phone {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "Wrong phone number",
		})
		return
	}

	s := logic.SelectVisitor(form.VisitorName, form.Phone)

	var results []Visitors

	for i := 0; i <= len(s)-1; i++ {
		visistorName := s[i].VisitorName
		createdAt := s[i].CreatedAt
		event := s[i].Event
		var user Visitors
		user = Visitors{VisitorName: visistorName, CreatedAt: createdAt, Event: event}

		results = append(results, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"results": results,
		"message": "search successful",
	})

}

func Delete(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	reuslt := logic.DeleteId(id)
	result := logic.AllData()
	var users []Users

	for i := 0; i <= len(result)-1; i++ {
		id := result[i].ID
		visistorName := result[i].VisitorName
		sex := result[i].Sex
		createdAt := result[i].CreatedAt
		event := result[i].Event
		var user Users
		user = Users{ID: id, VisitorName: visistorName, Sex: sex, CreatedAt: createdAt, Event: event}

		users = append(users, user)
	}
	if reuslt == true {

		c.JSON(http.StatusOK, gin.H{
			"data":    users,
			"message": "successfully deleted",
		})
		return
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":    users,
			"message": "failed to delete",
		})
	}
}

func Update(c *gin.Context) {
	visitorName := c.PostForm("visitor_name")
	visitId := c.PostForm("visit_id")
	phone := c.PostForm("phone")

	if logic.SelectVisitorName(visitorName) == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "User does not exist",
		})
		return
	}

	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "wrong phone number",
		})
		return
	}

	if len(visitId) != 18 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "ID card is wrong",
		})
		return
	}

	reuslt := logic.Update(visitorName, phone, visitId)
	s := logic.SelectVisitor(visitorName, phone)
	var users []Visitors
	for i := 0; i <= len(s)-1; i++ {
		visistorName := s[i].VisitorName
		createdAt := s[i].CreatedAt
		event := s[i].Event
		var user Visitors
		user = Visitors{VisitorName: visistorName, CreatedAt: createdAt, Event: event}

		users = append(users, user)
	}
	if reuslt == true {
		c.JSON(http.StatusOK, gin.H{
			"data":    users,
			"message": "update completed",
		})
		return
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"results": users,
			"message": "update failed",
		})
	}
}

//VisitorName string `gorm:"varchar(20);not null" form:"visitor_name" json:"visitor_name"`
//Sex string `gorm:"enum('男','女');not null" form:"sex" json:"sex"`
//Phone string `gorm:"varchar(20);not null" form:"phone" json:"phone"`
//VisitId string `gorm:"varchar(20);not null" form:"visit_id" json:"visit_id"`
//Event string `gorm:"varchar(100);not null" form:"event" json:"event"`

type Visitors struct {
	VisitorName string    `json:"visitor_name"`
	CreatedAt   time.Time `json:"created_at"`
	Event       string    `json:"event"`
}

type Search struct {
	VisitorName string `form:"visitorName" json:"visitorName" uri:"visitorName" xml:"visitorName" binding:"required"`
	Phone       string `form:"phone" json:"phone" uri:"phone" xml:"phone" binding:"required"`
}
