package databaseOperations

import (
	"log"
	"strconv"

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

func GetAllRecentBuyOperation(userId int) ([]RecentBuyOperation, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var operations []RecentBuyOperation
	result := db.Where("user_id = ?", strconv.Itoa(userId)).Find(&operations)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return operations, nil
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
	db.Select("id = ?", operation.Id).Updates(&operation)
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

func GetAllArchivedOperation(userId int) ([]ArchivedOperation, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var operations []ArchivedOperation
	result := db.Where("user_id = ?", strconv.Itoa(userId)).Find(&operations)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return operations, nil
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
	if err != nil {
		log.Println(err.Error())
		return err
	}
	result := db.Create(&user)
	if result.Error != nil {
		//log.Println(result.Error)
		return result.Error
	}
	return nil
}

func GetByMail(mail string) (User, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	var user User
	result := db.Where("email = ?", mail).Find(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user, nil
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
	db.Select("id = ?", user.Id).Updates(&user)
	return nil
}
