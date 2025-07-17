package main 

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

// Database connection details
const ( 
	host     = "localhost"
	port	 = 5433
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func authMiddleware(c *fiber.Ctx) error {
	// Middleware to check for JWT token in cookies
	cookie  := c.Cookies("jwt_token") // Get the JWT token from cookies with the name "jwt_token"
	//for tsting only 
	jwtSecret := []byte("secret_key") // Secret key for signing the JWT token

	//check if the token is valid 
	token , err:= jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil // Return the secret key for validation
	})

	if err != nil {
		log.Println("Error parsing JWT token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Check if the token is valid and extract claims
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		log.Println("JWT token is valid. Claims:", claims)
	} else {
		log.Println("Invalid JWT token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next() // Proceed to the next handler
}

func main() {

	// Connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
  newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold: time.Second, // Slow SQL threshold
      LogLevel:      logger.Info, // Log level , there are 4 levels: Silent, Error, Warn, Info
      Colorful:      true,        // Enable color
    },
  )
	// Open a connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // Use the new logger
	})

	if err != nil {
		panic("failed to connect to database: " + err.Error()) //panic use to stop the program if the connection fails
	}
	log.Println("Successfully connected to the database!")
	
	// Migrate the schema
	err = db.AutoMigrate(&Book{}, &User{}) // Automatically create the table based on the Book struct

	if err != nil {
		panic("failed to migrate database: " + err.Error()) //panic use to stop the program if the migration fails
	}
	// // !!! It WON'T delete unused columns. you need to manually implement it.
	// if err != nil {
	// 	log.Fatalf("Error migrating database: %v", err)
	// }

	// //create a new book
	// book := Book{Name: "Go Programming", Author: "John Doe", Price: 29.99}
	// if err = createBook(db, &book); err != nil {
	// 	log.Fatalf("Error creating book: %v", err)
	// } else {
	// 	log.Println("Book created successfully!")
	// }

	// //get a book by ID
	// bookID := 11 // Assuming the book ID is 1
	// retrievedBook, err := getBookById(db, uint(bookID))
	// if err != nil {
	// 	log.Fatalf("Error retrieving book: %v", err)
	// } else {
	// 	log.Printf("Retrieved Book: %+v\n", retrievedBook)
	// }

	// //update a book
	// bookToUpdate, _:= getBookById(db, uint(bookID))
	// bookToUpdate.Name = "Advanced Go Programming"
	// bookToUpdate.Author = "Jane Doe"
	// bookToUpdate.Price = 69.69
	// if err = updateBook(db, bookToUpdate); err != nil {
	// 	log.Fatalf("Error updating book: %v", err)
	// } else {
	// 	log.Println("Book updated successfully!")
	// }

	// //delete a book
	// if err = deleteBook(db, uint(bookID)); err != nil {
	// 	log.Fatalf("Error deleting book: %v", err)
	// } else {
	// 	log.Println("Book deleted successfully!")
	// }
	// log.Println("All operations completed successfully!")

	// //get a book by name
	// bookName := "Go Programming"
	// if b, err := getBookByName(db, bookName); err != nil {
	// 	log.Fatalf("Error retrieving book by name: %v", err)
	// } else {
	// 	log.Printf("Retrieved Book by Name: %+v\n", b)
	// }

	// //get all books
	// if books, err := getAllBooks(db); err != nil {
	// 	log.Fatalf("Error retrieving all books: %v", err)
	// } else {
	// 	log.Printf("Retrieved All Books: %+v\n", books)
	// }

	// ========== Fiber Setup ==========
	app := fiber.New()

	// ========== User Routes ==========
	app.Post("/users/register" , func (c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		if err := createUSer(db, user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User registered successfully",
			"user":    user,
		})
	})

	app.Post("/users/login", func (c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		token, err := loginUser(db, user)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}

		//set cookie with JWT token
		c.Cookie(&fiber.Cookie{
			Name:     "jwt_token",
			Value:    token,
			Expires:  time.Now().Add(72 * time.Hour), // Set expiration to 72 hours
			HTTPOnly: true, // Prevent JavaScript access to the cookie
		})

		return c.JSON(fiber.Map{
			"message": "Login successful",
		})
	})

	app.Use(authMiddleware) // Apply the authentication middleware to all routes

	// ========== Book Routes ==========
	//get all books
	app.Get("/books", func (c *fiber.Ctx) error {
		books , err := getAllBooks(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(books)
	})

	//get book by ID
	app.Get("/books/:id" ,func (c *fiber.Ctx) error {
		bid , err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid book ID",
			})
		}
		if book , err := getBookById(db, uint(bid)); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		} else {
			return c.JSON(book)
		}
	})

	//create a new book
	app.Post("/books", func (c *fiber.Ctx) error {
		var newBook Book
		if err := c.BodyParser(&newBook); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		if err := createBook(db, &newBook); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Book created successfully",
			"book":    newBook,
	})
	})

	//update a book`
	app.Put("/books/:id", func (c *fiber.Ctx) error {
		bid, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid book ID",
			})
		}
		var updatedBook Book
		if err := c.BodyParser(&updatedBook); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		updatedBook.ID = uint(bid) // Set the ID for the book to update
		if err := updateBook(db, &updatedBook); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "Book updated successfully",
			"book":    updatedBook,
		})
	})

	//delete a book
	app.Delete("/books/:id", func (c *fiber.Ctx) error {
		bid, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid book ID",
			})
		}
		if err := deleteBook(db, uint(bid)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "Book deleted successfully",
		})
	})

	app.Listen(":8080")

}
