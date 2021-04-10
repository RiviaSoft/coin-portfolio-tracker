package Business

import (
	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"

	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetAllArchivedOperations(userId int) []ArchivedOperation {
	// db katmanına gönderme
	return []ArchivedOperation{}
}
func AddArchivedOperation(operation ArchivedOperation) Result {
	// db katmanına gönderme
	return Result{
		Success: true,
		Message: Messages.OperationArchived,
	}
}
func DeleteArchivedOperation(operation ArchivedOperation) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.OperationRemoved,
	}
}
