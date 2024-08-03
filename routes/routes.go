package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hammasmart/reservation-e/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addresource", controllers.ResourceViewerAdmin())
	incomingRoutes.POST("/users/resourceview", controllers.SearchResource())
	incomingRoutes.POST("/users/search", controllers.SearchResourceByQuery())
}
