package Business

import (
	"log"

	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/DataAccess/databaseOperations"
)

func AddWalletOperation(walletOperation WalletOperation) Result {
	err := databaseOperations.AddWalletOperation(walletOperation)
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

func GetAllWalletOperation(walletId int) ([]WalletOperationDTO, error) {
	wallets, err := databaseOperations.GetAllWalletOperation(walletId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return wallets, nil
}

func DeleteWalletOperation(walletOperation WalletOperation) Result {
	err := databaseOperations.DeleteWalletOperation(walletOperation)
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
