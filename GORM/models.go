package main

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model // This will add fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name  string `json:"name"`
	Author string `json:"author"`
	Price float64 `json:"price"`
} //in Gorm, the first letter of the field must be capitalized to be exported and accessible outside the package

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(&book)
	return result.Error
}

func getBookById(db *gorm.DB, id uint) (*Book, error) {
	var book Book
	result := db.First(&book, id) // Find the book by ID
	if result.Error != nil {
		return &Book{}, result.Error
	}
	return &book, nil
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Save(&book) // Save the updated book
	return result.Error
}

//id you use GORM model. then it will conduct a soft delete
// which means it will not delete the record from the database, but it will set the DeletedAt field to the current time.
// with this, you can still retrieve the record later if needed. but you can't query it directly unless you use Unscoped() method
func deleteBook(db *gorm.DB, id uint) error {
	result := db.Delete(&Book{}, id) // soft delete the book by ID
	//if you want to hard delete the book, you can use Unscoped() method
	// result := db.Unscoped().Delete(&Book{}, id) // hard delete the book by ID
	return result.Error
}

func getBookByName(db *gorm.DB, name string) (*Book ,error) { 
	var book Book
	result := db.Where("name = ?", name).First(&book) // Find the book by name
	if result.Error != nil {
		return &Book{}, result.Error
	}
	return &book, nil
}

func getAllBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	result := db.Find(&books) // Find all books
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
