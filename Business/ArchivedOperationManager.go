package Business

import (
	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"

	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetAllArchivedOperations(userId int) []ArchivedOperation {
	// db katmanına gönderme
	return
}
func AddArchivedOperation(userId int, operation ArchivedOperation) Result {
	// db katmanına gönderme
	return Result{
		Success: true,
		Message: Messages.OperationArchived,
	}
}
func DeleteArchivedOperation(userId int, operation ArchivedOperation) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.OperationRemoved,
	}
}
