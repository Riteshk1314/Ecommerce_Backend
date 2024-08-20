package routes

import (
	"github.com/gin-gonic/gin",
	"../controllers"

)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp)
	incomingRoutes.POST("/users/loging", controllers.Login)
	incomingRoutes.POST("/admin/addproduct", controllers.AddProduct)
	incomingRoutes.GET("/users/productview", controllers.ViewProduct)
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery)
}