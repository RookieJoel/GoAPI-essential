package main

import (
	"github.com/gofiber/fiber/v2" //import fiber
	"strconv" //for converting string to int
)


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
	

