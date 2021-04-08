package Security

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

var (
	MY_SIGNING_KEY  = []byte("mysupersecretkey")
	USERNAME        = "msrexe"
	PASSWORD        = "123456"
	EXPIRATION_TIME = time.Now().Add(time.Hour * 2).Unix()
	isAdmin         = true
	isUserExist     = false
)

func SetupMW(app *fiber.App) {

	//JWTware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: MY_SIGNING_KEY,
	}))

	app.Post("/login", loginHandler)
}

func register(c *fiber.Ctx) error {
	//register işlemleri
}

func loginHandler(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != USERNAME || pass != PASSWORD {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uname"] = USERNAME
	claims["admin"] = isAdmin
	claims["exp"] = EXPIRATION_TIME

	// Generate encoded token and send it as response.
	t, err := token.SignedString(MY_SIGNING_KEY)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
