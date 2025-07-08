package main

import (
	"fmt"
	"log"
	"time"

	// "net/http"
	"os"

	"github.com/gofiber/fiber/v2" //import fiber
	"github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv" //import godotenv for loading environment variables
)

//this is like using pure http package
//w = res and r = req
// func helloHandler(w http.ResponseWriter, r *http.Request) {

// 	if r.URL.Path != "/greet" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}

// 	// Check if the request method is GET
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	//send back a response
// 	fmt.Fprintf(w, "Hello, World! You've reached the Go API server.")

// }

// Book struct to hold book data
type Book struct {
	// response data will be in JSON format as indicated by the tags
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Sample book data (in-memory)
var books []Book

func main() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	// http.HandleFunc("/greet", helloHandler)
	
	// port := ":8080"
	// fmt.Printf("Server is running on port %s\n", port)
	// if err := http.ListenAndServe(port, nil); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }


	// Using Fiber framework
	// this is like using express in Node.js
	//code down below is auto error handled, no need to check for errors like in pure http package
	app := fiber.New() //this is like app = express()
	app.Get("/greet", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! You've reached the Go API server.")
		});
		
		//CRUD operations
		
		books = append(books ,Book{
			ID:     1,
			Title:  "The Go Programming Language",
			Author: "Alan Donovan",
		},
		Book{
			ID:     2,
			Title:  "Learning Go",
			Author: "Jon Bodner",
		},
		Book{
			ID:     3,
			Title:  "Go in Action",
			Author: "William Kennedy",
		},
	)
	
	//read all books
	// app.Get("/books", func(c *fiber.Ctx) error {
	// 	return c.JSON(books) //returning books as JSON
	// });

	//login
	app.Post("/login", login)

	//Middleware for JWT authentication
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")), //get JWT secret from environment
	}))

	//or you can use a separate function for the handler
	app.Get("/books", getBooks) //using a separate function for the handler
	app.Get("/books/:id", getBookByID) 

	//create a new book
	app.Post("/books", createBook)

	//update a book
	app.Put("books/:id", updateBook) 

	//delete a book
	app.Delete("/books/:id",deleteBook )
	
	//get environment variable
	app.Get("/env", getEnv)

	
	
	
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(c *fiber.Ctx) error {
	// Get the SECRET environment variable
	secret := os.Getenv("SECRET")
	if secret == "" {
		return c.Status(fiber.StatusInternalServerError).SendString("SECRET not set")
	}
	return c.JSON(fiber.Map{
		"SECRET" : secret,
	})
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//dummy user for login
var memberUser = User{
	Username: "admin",
	Password: "password",
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Username != memberUser.Username && user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      jwt.TimeFunc().Add(24 * time.Hour).Unix(), // token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not generate token")
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"username": user.Username,
		"token": tokenString,
	})
}