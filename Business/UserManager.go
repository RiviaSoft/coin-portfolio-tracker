package Business

import (
	"log"

	Messages "github.com/msrexe/portfolio-tracker/Core/Message"
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/DataAccess/databaseOperations"
	Security "github.com/msrexe/portfolio-tracker/Security"
)

func GetUser(mail string) (User, error) {
	result, err := databaseOperations.GetByMail(mail)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	return result, nil
}

func AddUser(user UserModel) Result {
	hashedPass, _ := Security.HashPassword(user.Password)
	hashedUser := User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: hashedPass,
	}
	err := databaseOperations.AddUser(hashedUser)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	if !Security.VerifyMail(hashedUser.Email) {
		return Result{
			Success: false,
			Message: Messages.InvalidMail,
		}
	}
	return Result{
		Success: true,
		Message: Messages.UserAdded,
	}
}
func DeleteUser(user User) Result {
	err := databaseOperations.DeleteUser(user)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	return Result{
		Success: true,
		Message: Messages.UserDeleted,
	}

}
func UpdateUser(user UserModel) Result {
	hashedPass, _ := Security.HashPassword(user.Password)
	hashedUser := User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: hashedPass,
	}
	err := databaseOperations.UpdateUser(hashedUser)
	if err != nil {
		return Result{
			Success: false,
			Message: err.Error(),
		}
	}
	return Result{
		Success: true,
		Message: Messages.UserUpdated,
	}
}
