package Security

import (
	"encoding/json"
	"log"
	"time"

	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

var (
	MY_SIGNING_KEY  = []byte("mysupersecretkey")
	UID             = 12
	USERNAME        = "msrexe"
	PASSWORD        = "123456"
	EMAIL           = "mlheymen.ms@gmail.com"
	EXPIRATION_TIME = time.Now().Add(time.Hour * 2).Unix()
	isAdmin         = true
	//isUserExist          = false
	BROKEN_TOKEN_MESSAGE = "Token is broken"
)

func SetupMW(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", loginHandler)
	auth.Post("/register", registerHandler)

	//JWTware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: MY_SIGNING_KEY,
	}))

}

func registerHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.AddUser(user)
	if result.Success {
		return c.JSON(result)
	}
	return c.SendStatus(409) // already exists
}

func loginHandler(c *fiber.Ctx) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")
	log.Println(user, pass)

	// Throws Unauthorized error
	if user != USERNAME || pass != PASSWORD {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = UID
	claims["uname"] = USERNAME
	claims["email"] = EMAIL
	claims["admin"] = isAdmin
	claims["exp"] = EXPIRATION_TIME

	// Generate encoded token and send it as response.
	t, err := token.SignedString(MY_SIGNING_KEY)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func GetUserClaims(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// uname := claims["uname"].(string)
	// mail := claims["email"].(string)
	return claims
}
