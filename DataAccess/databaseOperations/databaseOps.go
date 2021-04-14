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

//***********************************************************************************
//WALLET
func AddWallet(wallet Wallet) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	result := db.Create(&wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllWallet(userId int) ([]Wallet, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var wallets []Wallet
	result := db.Where("user_id = ?", strconv.Itoa(userId)).Find(&wallets)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return wallets, nil
}

func GetWalletByName(userId int, walletName string) ([]Wallet, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var wallets []Wallet
	result := db.Where("user_id = ? AND name = ?", strconv.Itoa(userId), walletName).Find(&wallets)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return wallets, nil
}

func DeleteWallet(wallet Wallet) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	db.Delete(&wallet)
	return nil
}

func DeleteAllWallet(userId int, walletName string) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	result := db.Where("user_id = ? AND name = ?", strconv.Itoa(userId), walletName)
	if result.Error != nil {
		log.Println(result.Error)
		return err
	}
	db.Delete(&result)
	return nil
}
