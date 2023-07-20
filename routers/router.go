package routers

import (
	"VisitorsManagementSystem/controllers"
	"VisitorsManagementSystem/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.POST("/login",controllers.Login)
	V1 := router.Group("/main")
	{
		V1.POST("/search",controllers.Select)

		V1.GET("/delete/:id",controllers.Delete)
		V1.GET("/index",controllers.AllVisitors)
		V1.POST("/update",controllers.Update)
	}

	router.POST("/register",controllers.AddVisitor)

	return router
}
