package DAO

import (
	"../Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
   )

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]Credentials) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
	 return err
	}
	return nil
   }

//CreateUser ... Insert New data
func CreateUser(user *Credentials) (err error) {
	err = Config.DB.Create(user).Error;
	if err != nil {
		return err
	}
	return nil
   }

//GetUserByUsername ... Fetch only one user by Username
func GetUserByUsername(user *Credentials, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
   }

//UpdateUser ... Update user   
func UpdateUser(user *Credentials, username string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
   }

//DeleteUser ... Delete user   
func DeleteUser(user *Credentials, username string) (err error) {
	Config.DB.Where("username = ?", username).Delete(user)
	return nil
   }