package controllers

import (
	"VisitorsManagementSystem/logic"
	"VisitorsManagementSystem/middleware"
	"VisitorsManagementSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)


//登录
func Login(ctx *gin.Context) {
	//获取参数
	var form login

	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
/*
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
*/
	//数据验证
	if len(form.Username) <= 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "name cannot be empty",
		})
		return
	}
	//验证是否有该用户
	if form.Username != logic.SelectAdmin(form.Username) {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":    422,
			"message": "User does not exist",
		})
		return
	}

	if len(form.Password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "The password cannot be less than 6 characters",
		})
		return
	}


	//判断密码是否正确
	if form.Password != logic.SelectPassword(form.Username) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "wrong password",
		})
		return
	}


	user := models.Admin{
		Username: form.Username,
		Password: form.Password,
	}

	token ,err := middleware.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"code":"500",
			"msg":"System exception",
		})

		log.Panicf("token generate error:%v",err)
		return
	}


	//返回结果
	result := logic.AllData()

	var users []Users

	for i:=0;i<=len(result)-1;i++{
		id := result[i].ID
		visistorName := result[i].VisitorName
		sex := result[i].Sex
		createdAt := result[i].CreatedAt
		event := result[i].Event
		var user Users
		user = Users{ID: id,VisitorName: visistorName,Sex: sex,CreatedAt: createdAt,Event: event}

		users = append(users,user)
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code":"200",
		"data":users,
		"msg":gin.H{"token":token},
		"message":"login successful",
	})

}

type login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username    string `form:"username" json:"username" uri:"username" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}


type Users struct {
	ID uint `json:"ID"`
	VisitorName string `json:"VisitorName"`
	Sex string `json:"Sex"`
	CreatedAt time.Time `json:"CreatedAt"`
	Event string `json:"Event"`
}