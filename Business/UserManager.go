package Business

import (
	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetUser(userId int) User {
	// db katmanına gönderme

	return User{}
}
func AddUser(user User) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.UserAdded,
	}
}
func DeleteUser(userId int, user User) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.UserDeleted,
	}

}
func UpdateUser(userId int, user User) Result {
	// db katmanına gönderme

	return Result{
		Success: true,
		Message: Messages.UserUpdated,
	}
}
