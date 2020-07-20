package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	dbr "github.com/gocraft/dbr/v2"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
	_ "github.com/mattn/go-sqlite3"

	jwtware "github.com/gofiber/jwt"
)

func main() {

	// create a connection (e.g. "postgres", "mysql", or "sqlite3")
	conn, err := dbr.Open("sqlite3", "./test.sqlite", nil)
	if err != nil {
		fmt.Println("Error connecting: ", err)
	}
	conn.SetMaxOpenConns(10)

	// create a session for each business unit of execution (e.g. a web request or goworkers job)
	sess := conn.NewSession(nil)

	// create a tx from sessions
	sess.Begin()

	app := fiber.New()

	app.Use(cors.New())
	fmt.Printf("Using cors.")

	// 3 requests per 10 seconds max
	cfg := limiter.Config{
		Timeout: 10,
		Max:     3,
	}

	app.Use(limiter.New(cfg))

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
