package databaseOperations

import (
	"log"
	"strconv"

	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

//RECENT BUY OPERATION
func AddRecentBuyOperation(operation RecentBuyOperation) error {
	db, err := ConnectDB()
	conn, _ := db.DB()
	defer conn.Close()
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
	conn, _ := db.DB()
	defer conn.Close()
	var operations []RecentBuyOperation
	result := db.Where("user_id = ?", strconv.Itoa(userId)).Find(&operations)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return operations, nil
}

func GetRecentBuyOperationById(userId int, id int) (RecentBuyOperation, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return RecentBuyOperation{}, err
	}
	conn, _ := db.DB()
	defer conn.Close()
	var operation RecentBuyOperation
	result := db.Table("recent_buy_operations").Where("user_id = ? AND id = ?", strconv.Itoa(userId), strconv.Itoa(id)).Find(&operation)
	if result.Error != nil {
		log.Println(result.Error)
		return RecentBuyOperation{}, err
	}
	return operation, nil
}

func DeleteRecentBuyOperation(operation RecentBuyOperation) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	conn, _ := db.DB()
	defer conn.Close()
	db.Delete(&operation)
	return nil
}

func UpdateRecentBuyOperation(operation RecentBuyOperation) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	conn, _ := db.DB()
	defer conn.Close()
	db.Select("id = ?", operation.Id).Updates(&operation)
	return nil
}

//***********************************************************************************
//ARCHIVED OPERATION
func AddArchivedOperation(operation ArchivedOperation) error {
	db, err := ConnectDB()
	conn, _ := db.DB()
	defer conn.Close()
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
	conn, _ := db.DB()
	defer conn.Close()
	var operations []ArchivedOperation
	result := db.Where("user_id = ?", strconv.Itoa(userId)).Find(&operations)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return operations, nil
}

func GetArchivedOperationById(userId int, id int) (ArchivedOperation, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return ArchivedOperation{}, err
	}
	conn, _ := db.DB()
	defer conn.Close()
	var operation ArchivedOperation
	result := db.Table("archived_operations").Where("user_id = ? AND id = ?", strconv.Itoa(userId), strconv.Itoa(id)).Find(&operation)
	if result.Error != nil {
		log.Println(result.Error)
		return ArchivedOperation{}, err
	}
	return operation, nil
}

func DeleteArchivedOperation(operation ArchivedOperation) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	conn, _ := db.DB()
	defer conn.Close()
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
	conn, _ := db.DB()
	defer conn.Close()
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
	conn, _ := db.DB()
	defer conn.Close()
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
	conn, _ := db.DB()
	defer conn.Close()
	db.Delete(&user)
	return nil
}

func UpdateUser(user User) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	conn, _ := db.DB()
	defer conn.Close()
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
	conn, _ := db.DB()
	defer conn.Close()
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
	conn, _ := db.DB()
	defer conn.Close()
	var wallets []Wallet
	result := db.Where("user_id = ?", strconv.Itoa(userId)).Find(&wallets)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return wallets, nil
}

func GetWalletById(userId int, id int) (Wallet, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return Wallet{}, err
	}
	conn, _ := db.DB()
	defer conn.Close()
	var wallet Wallet
	result := db.Where("user_id = ? AND id = ?", strconv.Itoa(userId), strconv.Itoa(id)).Find(&wallet)
	if result.Error != nil {
		log.Println(result.Error)
		return Wallet{}, err
	}
	return wallet, nil
}

func DeleteWallet(wallet Wallet) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	conn, _ := db.DB()
	defer conn.Close()
	db.Delete(&wallet)
	return nil
}

//***********************************************************************************
//WALLET OPERATIONS
func AddWalletOperation(walletOperation WalletOperation) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	conn, _ := db.DB()
	defer conn.Close()
	result := db.Create(&walletOperation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllWalletOperation(walletId int) ([]WalletOperation, error) {
	db, err := ConnectDB()
	conn, _ := db.DB()
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var walletOperations []WalletOperation
	result := db.Where("wallet_id = ?", strconv.Itoa(walletId)).Find(&walletOperations)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, err
	}
	return walletOperations, nil
}

func DeleteWalletOperation(walletOperation WalletOperation) error {
	db, err := ConnectDB()
	conn, _ := db.DB()
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	db.Delete(&walletOperation)
	return nil
}

func DeleteAllWalletOperations(walletId int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	conn, _ := db.DB()
	defer conn.Close()
	result := db.Table("wallet_operations").Where("wallet_id = ?", strconv.Itoa(walletId))
	if result.Error != nil {
		log.Println(result.Error)
		return err
	}
	db.Delete(&result)
	return nil
}
