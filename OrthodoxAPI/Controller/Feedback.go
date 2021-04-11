package Controller

import (
	"../DAO"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
   )


//GetUsers ... Get all users
func GetAllFeedbacks(c *gin.Context) {
	username := c.Params.ByName("username")
	var reportobj []DAO.Feedback
	
	err := DAO.GetAllFeedbacks(&reportobj, username)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reportobj)
	}
}

//CreateUser ... Create User
func CreateFeedback(c *gin.Context) {
	var reportobj DAO.Feedback
	// Turn reportobj into an json and return
	c.BindJSON(&reportobj)
	err := DAO.CreateFeedback(&reportobj)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reportobj)
	}
}