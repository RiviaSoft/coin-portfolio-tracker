package Business

import (
	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetUser(mail string) User {
	// db katmanına gönderme

	return User{}
}
func AddUser(user UserModel) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.UserAdded,
	}
}
func DeleteUser(userId int, user UserModel) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.UserDeleted,
	}

}
func UpdateUser(userId int, user UserModel) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.UserUpdated,
	}
}
