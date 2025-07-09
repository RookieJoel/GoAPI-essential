package main 

import (
	"fmt"
	"log"
	// "github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

// Database connection details
const ( 
	host     = "localhost"
	port	 = 5433
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

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
	err = db.AutoMigrate(&Book{}) // Automatically create the table based on the Book struct
	// !!! It WON'T delete unused columns. you need to manually implement it.
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	//create a new book
	book := Book{Name: "Go Programming", Author: "John Doe", Price: 29.99}
	if err = createBook(db, &book); err != nil {
		log.Fatalf("Error creating book: %v", err)
	} else {
		log.Println("Book created successfully!")
	}

	//get a book by ID
	bookID := 5 // Assuming the book ID is 1
	retrievedBook, err := getBookById(db, uint(bookID))
	if err != nil {
		log.Fatalf("Error retrieving book: %v", err)
	} else {
		log.Printf("Retrieved Book: %+v\n", retrievedBook)
	}

	//update a book
	bookToUpdate, _:= getBookById(db, uint(bookID))
	bookToUpdate.Name = "Advanced Go Programming"
	bookToUpdate.Author = "Jane Doe"
	bookToUpdate.Price = 69.69
	if err = updateBook(db, bookToUpdate); err != nil {
		log.Fatalf("Error updating book: %v", err)
	} else {
		log.Println("Book updated successfully!")
	}

	//delete a book
	if err = deleteBook(db, uint(bookID)); err != nil {
		log.Fatalf("Error deleting book: %v", err)
	} else {
		log.Println("Book deleted successfully!")
	}
	log.Println("All operations completed successfully!")


}
