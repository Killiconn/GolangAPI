package DAO

import (
	"../Config"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
   )


// type Result struct {
	// Name string
	// Age  string
// }

//GetAllUsers Fetch all user data
func GetAllFeedbacks(reportobj *[]Feedback, email string) (err error) {
	// if err = Config.DB.Where("Username = ?", email).Find(&reportobj).Error; err != nil {
	// 	return err
	// }
	// return nil

	// var result Result

	Config.DB.Raw("SELECT Username, Feedback FROM feedbacks WHERE Username = ?", email).Scan(&reportobj)
	// Iterate through, then print

	// fmt.Println(reportobj.Username)
	// fmt.Println(reportobj.Feedback)
	
	// fmt.Println("%v", result)
	return nil


	// for results.Next() {
	// 	var tag tag

	// 	err = results.Scan(&reportobj.Username, &reportobj.Feedback)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	log.Printf(tag.Name)
	// }
   }


//CreateUser ... Insert New data
func CreateFeedback(reportobj *Feedback) (err error) {
	err = Config.DB.Create(reportobj).Error;
	if err != nil {
		return err
	}
	return nil
   }
