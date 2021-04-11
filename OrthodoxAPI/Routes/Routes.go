package Routes

import (
	"../Controller"
	"github.com/gin-gonic/gin"
   )

//SetupRouter ... Configure routes
// Make Controller object
// That makes queries to the database'\

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controller.GetUsers)
		grp1.POST("user", Controller.CreateUser)
		grp1.GET("user/:username", Controller.GetUserByUsername)
		grp1.PUT("user/:username", Controller.UpdateUser)
		grp1.DELETE("user/:username", Controller.DeleteUser)
		
	}
	grp2 := r.Group("/user-api-feedback")
	{
		grp2.GET("user/:username", Controller.GetAllFeedbacks)
		grp2.POST("user", Controller.CreateFeedback)
	}
	return r
}