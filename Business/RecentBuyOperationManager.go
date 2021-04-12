package Business

import (
	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/DataAccess/databaseOperations"
)

func GetAllBuyOperations(userId int) ([]RecentBuyOperation, error) {
	result, err := databaseOperations.GetAllRecentBuyOperation(userId)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func AddBuyOperation(operation RecentBuyOperation) Result {
	recentBuyOperations, err := databaseOperations.GetAllRecentBuyOperation(operation.UserId)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	recordIsExist := false
	for _, value := range recentBuyOperations {
		if value.CoinSymbol == operation.CoinSymbol {
			recordIsExist = true
			totalCoinAmount := value.CoinAmount + operation.CoinAmount
			value.BuyCost = ((value.BuyCost * value.CoinAmount) + (operation.BuyCost * operation.CoinAmount)) / totalCoinAmount
			value.CoinAmount = totalCoinAmount
			if err := databaseOperations.UpdateRecentBuyOperation(value); err != nil {
				return Result{
					Success: false,
					Message: err.Error(),
				}
			}
		}
	}

	if !recordIsExist {
		if err := databaseOperations.AddRecentBuyOperation(operation); err != nil {
			return Result{
				Success: false,
				Message: err.Error(),
			}
		}
	}
	return Result{
		Success: true,
		Message: Messages.OperationAdded,
	}
}

func DeleteBuyOperation(operation RecentBuyOperation) Result {
	if err := databaseOperations.DeleteRecentBuyOperation(operation); err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}

	return Result{
		Success: true,
		Message: Messages.OperationRemoved,
	}

}
func UpdateBuyOperation(operation RecentBuyOperation) Result {
	if err := databaseOperations.UpdateRecentBuyOperation(operation); err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}

	return Result{
		Success: true,
		Message: Messages.OperationUpdated,
	}
}
