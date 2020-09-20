package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/mattn/go-sqlite3"

	jwtware "github.com/gofiber/jwt/v2"
)

func main() {

	// initialize resources
	dbInit()
	go redisInit()

	app := fiber.New()

	app.Use(cors.New())
	fmt.Printf("Using cors.")

	// 3 requests per 10 seconds max
	cfg := limiter.Config{
		Duration: 10 * time.Second,
		Max:      3,
	}

	app.Use(limiter.New(cfg))

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		return c.Next()
	})

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("🥈 Second handler")
		return c.Next()
	})

	// GET /api/register
	app.Get("/api/list", func(c *fiber.Ctx) error {
		fmt.Println("🥉 Last handler")
		return c.Send([]byte("Hello, World 👋!"))
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

	app.Listen(":3000")
}

func accessible(c *fiber.Ctx) error {
	return c.Send([]byte("Accessible"))
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.Send([]byte("Welcome " + name))
}
