package main

import (
	"VisitorsManagementSystem/dao"
	"VisitorsManagementSystem/routers"
)

func main() {
	dao.InitDB()

	router := routers.Router()
	router.Run(":8080")
}
