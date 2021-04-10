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
func AddBuyOperation(operation RecentBuyOperation) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.OperationAdded,
	}
}
func DeleteBuyOperation(operation RecentBuyOperation) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.OperationRemoved,
	}

}
func UpdateBuyOperation(operation RecentBuyOperation) Result {

	return Result{
		Success: true,
		Message: Messages.OperationUpdated,
	}
}
