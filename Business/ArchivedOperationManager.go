package Business

import (
	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	"github.com/msrexe/portfolio-tracker/DataAccess/databaseOperations"

	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetAllArchivedOperations(userId int) ([]ArchivedOperation, error) {
	result, err := databaseOperations.GetAllArchivedOperation(userId)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func AddArchivedOperation(operation ArchivedOperation) Result {
	recentBuyOperations, err := databaseOperations.GetAllRecentBuyOperation(operation.UserId)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	for _, value := range recentBuyOperations {
		if value.CoinSymbol == operation.CoinSymbol {
			if value.CoinAmount > operation.CoinAmount {
				value.CoinAmount = value.CoinAmount - operation.CoinAmount
				if err := databaseOperations.UpdateRecentBuyOperation(value); err != nil {
					return Result{
						Success: false,
						Message: err.Error(),
					}
				}
				if err := databaseOperations.AddArchivedOperation(operation); err != nil {
					return Result{
						Success: false,
						Message: err.Error(),
					}
				}
			} else if value.CoinAmount == operation.CoinAmount {
				if err := databaseOperations.DeleteRecentBuyOperation(value); err != nil {
					return Result{
						Success: false,
						Message: err.Error(),
					}
				}
				if err := databaseOperations.AddArchivedOperation(operation); err != nil {
					return Result{
						Success: false,
						Message: err.Error(),
					}
				}
			}
		}
	}
	return Result{
		Success: true,
		Message: Messages.OperationArchived,
	}
}
func DeleteArchivedOperation(operation ArchivedOperation) Result {
	if err := databaseOperations.DeleteArchivedOperation(operation); err != nil {
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
