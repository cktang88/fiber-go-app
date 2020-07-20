package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"

	jwtware "github.com/gofiber/jwt"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	fmt.Printf("Using cors.")

	// Match any route
	app.Use(func(c *fiber.Ctx) {
		fmt.Println("🥇 First handler")
		c.Next()
	})

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) {
		fmt.Println("🥈 Second handler")
		c.Next()
	})

	// GET /api/register
	app.Get("/api/list", func(c *fiber.Ctx) {
		fmt.Println("🥉 Last handler")
		c.Send("Hello, World 👋!")
	})

	// Login route
	app.Post("/login", login)

	// Unauthenticated route
	app.Get("/", accessible)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	app.Listen(3000)
}

func login(c *fiber.Ctx) {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		c.SendStatus(fiber.StatusUnauthorized)
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) {
	c.Send("Accessible")
}

func restricted(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	c.Send("Welcome " + name)
}
