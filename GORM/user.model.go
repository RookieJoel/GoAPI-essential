package main

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type User struct { 
	gorm.Model // This will add fields ID, CreatedAt, UpdatedAt, DeletedAt
	Email   string `gorm:"unique;not null"` // Unique email field, not null
	Password string `gorm:"not null"` // Password field, not null
}

func createUSer (db *gorm.DB, user *User) error {
	// Hash the password before saving it to the database
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass) // Store the hashed password
	return db.Create(user).Error // Create the user in the database
}

// loginUser checks the user's credentials and returns a JWT token (string) if successful
func loginUser(db *gorm.DB, user *User) (string, error) {
	//get user form email
	selectedUser := new(User)
	res := db.Where("email = ?", user.Email).First(selectedUser)
	if res.Error != nil {
		return "", res.Error 
	}

	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password)); err != nil {
		return "", err // Passwords do not match
	}
	//pass => return jwt
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = selectedUser.ID // Set user ID in claims
	claims["email"] = selectedUser.Email // Set user email in claims
	claims["exp"] = jwt.TimeFunc().Add(time.Hour * 72).Unix() // Set expiration time (72 hours)

	// Sign the token with a secret key
	if tokenString, err := token.SignedString([]byte("secret_key")); err != nil {
		return "", err
	}else { 
		return tokenString, nil 
	}

}