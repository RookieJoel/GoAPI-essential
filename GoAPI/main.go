package main

import (
	"fmt"
	"log"
	// "net/http"
	"github.com/gofiber/fiber/v2" //import fiber
	"strconv" //for converting string to int
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
	// fmt.Println("Starting Go API server...")
	
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

	//or you can use a separate function for the handler
	app.Get("/books", getBooks) //using a separate function for the handler
	app.Get("/books/:id", getBookByID) 

	//create a new book
	app.Post("/books", createBook)

	//update a book
	app.Put("books/:id", updateBook) 

	//delete a book
	app.Delete("/books/:id",deleteBook )
	
	
	
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books) //returning books as JSON
}

func getBookByID(c *fiber.Ctx) error {
	id := c.Params("id") //getting the id (string) from the URL
	
	// Convert the id to an integer
	bookID, err := strconv.Atoi(id) //converting string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) //returning 400 if conversion fails
	}

	// Loop through the books to find the one with the matching ID
	// _ = index and book variables
	for _, book := range books {
		if book.ID == bookID {
			return c.JSON(book) //returning the book as JSON if found
		}
	}
	
	return c.Status(fiber.StatusNotFound).SendString("Book not found") //returning 404 if book not found
}

func createBook(c *fiber.Ctx) error {
	
	newBook := new(Book) // Create a new Book instance to hold the incoming data, this is like getting pointer to a Book struct

	// Parse the JSON body into the newBook variable
	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) //returning 400 if parsing fails
	}
	// Assign an ID to the new book (incremental)
	newBook.ID = len(books) + 1

	// Append the new book to the books slice
	books = append(books, *newBook) //append only accepts a value, so we need to send only the value of newBook, not the pointer

	return c.Status(fiber.StatusCreated).JSON(newBook) //returning 201 and the new book as JSON
}

func updateBook(c *fiber.Ctx) error {
	id := c.Params("id") //getting the id (string) from the URL
	
	// Convert the id to an integer
	bookID, err := strconv.Atoi(id) //converting string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) //returning 400 if conversion fails
	}

	bookUpdate := new(Book) // Create a new Book instance to hold the incoming data for update
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) //returning 400 if parsing fails
	}

	// Find the book by ID
	for index, book := range books {
		if book.ID == bookID {
			// Update the book's details
			books[index].Title = bookUpdate.Title
			books[index].Author = bookUpdate.Author
			return c.JSON(books[index]) //returning the updated book as JSON
		}
	}
	
	return c.Status(fiber.StatusNotFound).SendString("Book not found") //returning 404 if book not found
}

func deleteBook(c *fiber.Ctx) error {
	id:= c.Params("id");
	// Convert the id to an integer
	bookId , err := strconv.Atoi(id) 
	if err != nil { 
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) //returning 400 if conversion fails
	}

	// Loop through the books to find the one with the matching ID
	for idx , book := range  books {
		if book.ID == bookId{
			// Remove the book from the slice by appending the parts before and after it
			books = append(books[:idx], books[idx+1:]...)
			return c.SendString("Book deleted successfully") //returning success message
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Book not found") //returning 404 if book not found
}
	