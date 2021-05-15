package Business

import (
	"log"

	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/DataAccess/databaseOperations"
)

func AddWallet(wallet Wallet) Result {
	err := databaseOperations.AddWallet(wallet)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	return Result{
		Success: true,
		Message: Messages.WalletAdded,
	}
}

func GetWalletByUserId(userId int) ([]Wallet, error) {
	wallets, err := databaseOperations.GetAllWallet(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return wallets, nil
}
func GetWalletById(userId int, walletId int) (Wallet, error) {
	wallet, err := databaseOperations.GetWalletById(userId, walletId)
	if err != nil {
		log.Println(err.Error())
		return Wallet{}, err
	}
	return wallet, nil
}

func DeleteWallet(wallet Wallet) Result {
	err := databaseOperations.DeleteWallet(wallet)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	err = databaseOperations.DeleteAllWalletOperations(wallet.Id)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	return Result{
		Success: true,
		Message: Messages.WalletDeleted,
	}
}
