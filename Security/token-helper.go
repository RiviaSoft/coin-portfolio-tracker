package Security

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	. "github.com/msrexe/portfolio-tracker/Core/EnvVariables"
)

var (
	EXPIRATION_TIME = time.Now().Add(time.Hour * 60).Unix()
	//isUserExist          = false
)

func CreateToken(userId int, username, email string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = userId
	claims["uname"] = username
	claims["email"] = email
	//claims["admin"] = isAdmin
	claims["exp"] = EXPIRATION_TIME

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(GoDotEnvVariable("SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
