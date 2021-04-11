package main

import (
	"../OrthodoxAPI/Config"
	"../OrthodoxAPI/DAO"
	"../OrthodoxAPI/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
   )

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL())
		
	if err != nil {
		fmt.Println("Status:", err)
	}
	
	defer Config.DB.Close()
	
	Config.DB.AutoMigrate(&DAO.Credentials{})
	Config.DB.AutoMigrate(&DAO.Feedback{})


	r := Routes.SetupRouter()
	r.Run()
   }