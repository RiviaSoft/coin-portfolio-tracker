package Business

import (
	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetAllBuyOperations(userId int) []RecentBuyOperation {
	// db katmanına gönderme

	return []RecentBuyOperation{}
}
func AddBuyOperation(userId int, operation RecentBuyOperation) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.OperationAdded,
	}
}
func DeleteBuyOperation(userId int, operation RecentBuyOperation) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.OperationRemoved,
	}

}
func UpdateBuyOperation(userId int, operation RecentBuyOperation) Result {

	return Result{
		Success: true,
		Message: Messages.OperationUpdated,
	}
}
