package Controller

import (
	"../DAO"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
   )

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	//Creates the object Credentials in CredentialsModel file
	// This user will then contain all users
	var user []DAO.Credentials
	err := DAO.GetAllUsers(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)// User is returned here
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user DAO.Credentials
	c.BindJSON(&user)
	err := DAO.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByUsername ... Get the user by username   
func GetUserByUsername(c *gin.Context) {
	username := c.Params.ByName("username")
	var user DAO.Credentials
	err := DAO.GetUserByUsername(&user, username)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user DAO.Credentials
	username := c.Params.ByName("username")
	err := DAO.GetUserByUsername(&user, username)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = DAO.UpdateUser(&user, username)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user DAO.Credentials
	username := c.Params.ByName("username")
	err := DAO.DeleteUser(&user, username)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, gin.H{"username" + username: "is deleted"})
	}
}
