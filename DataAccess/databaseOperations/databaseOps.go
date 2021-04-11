package databaseOperations

import (
	"log"

	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

//RECENT BUY OPERATION
func AddRecentBuyOperation(operation RecentBuyOperation) error {
	db, err := ConnectDB()
	db.Create(&operation)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetAllRecentBuyOperation(userId int) []RecentBuyOperation {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
	}
	var operations []RecentBuyOperation
	result := db.Where("userid = ?", userId).Find(&operations)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return operations
}

func DeleteRecentBuyOperation(operation RecentBuyOperation) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	db.Delete(&operation)
	return nil
}

func UpdateRecentBuyOperation(operation RecentBuyOperation) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	db.Where("id = ?", operation.Id).Updates(&operation)
	return nil
}

//***********************************************************************************
//ARCHIVED OPERATION
func AddArchivedOperation(operation ArchivedOperation) error {
	db, err := ConnectDB()
	db.Create(&operation)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetAllArchivedOperation(userId int) []ArchivedOperation {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
	}
	var operations []ArchivedOperation
	result := db.Where("userid = ?", userId).Find(&operations)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return operations
}

func DeleteArchivedOperation(operation ArchivedOperation) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	db.Delete(&operation)
	return nil
}

//***********************************************************************************
//USER
func AddUser(user User) error {
	db, err := ConnectDB()
	db.Create(&user)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetByMail(mail string) User {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
	}
	var user User
	result := db.Where("email = ?", mail).First(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user
}

func DeleteUser(user User) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	db.Delete(&user)
	return nil
}

func UpdateUser(user User) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	db.Where("id = ?", user.Id).Updates(&user)
	return nil
}
