// package routes

// import (
// 	"example/controllers"
// 	"example/services"

// 	"github.com/gin-gonic/gin"
// )

// func SetupRouter() *gin.Engine {
// 	r := gin.Default()

// 	UserRoute(r)

// 	return r
// }

// func UserRoute(r *gin.Engine) {
// 	adminService := services.LoginService()
// 	adminController := controllers.NewLoginController()(adminService)

// 	admin := r.Group("/admin")
// 	{
// 		admin.POST("/login", adminController)
// 	}
// }
