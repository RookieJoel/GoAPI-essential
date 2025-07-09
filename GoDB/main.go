package main

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

const (
	host = "localhost"
	port = 5433
	dbname = "mydatabase"
	user = "myuser"
	password = "mypassword"
)

var db *sql.DB

func main() { 
	// Connection string
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  // Open a connection
  sdb, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    log.Fatal(err)
  }
  db = sdb
  defer db.Close()
  // Check the connection
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Successfully connected!")

  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("Welcome to the GoDB API!")
  }) 

  //get products using Fiber
  app.Get("/products/:id", getProductsHandler)

  //get all products using Fiber
  app.Get("/products", getAllProductsGHandler)

  //create a new product using Fiber
  app.Post("/products", createProductsHandler)

  // update a product using Fiber
  app.Put("/products/:id", updateProductsHandler)

  // delete a product using Fiber
  app.Delete("/products/:id", deleteProductsHandler)

  app.Listen(":8080")

//   //create a product
//   err = createProduct(&Product{
// 	Name:  "Sample Product",
// 	Price: 100,
//   })
//   if err != nil {
// 	log.Fatalf("Error creating product: %v", err)
//   } else {
// 	log.Println("Product created successfully")
//   }

//   //get a product by ID
//   productID := 1 // Assuming we want to get the product with ID 1
//   product, err := getProductByID(productID)
//   if err != nil {
// 	log.Fatalf("Error getting product: %v", err)
//   } else {
// 	log.Printf("Product retrieved: %+v", product)
//   }

//   //get all products
//   products, err := getAllProducts()
//   if err != nil {
// 	log.Fatalf("Error getting all products: %v", err)
//   } else {
// 	log.Println("All products retrieved:")
// 	for _, p := range products {
// 	  log.Printf("Product ID: %d, Name: %s, Price: %d", p.ID, p.Name, p.Price)
// 	}
//   }

//   //update a product
//   err = updateProduct(productID, &Product{
// 	Name:  "Updated Product",
// 	Price: 150,
//   })
//   if err != nil {
// 	log.Fatalf("Error updating product: %v", err)
//   } else {
// 	log.Println("Product updated successfully")
//   }

//   //delete a product
//   err = deleteProduct(1)
//   if err != nil {
// 	log.Fatalf("Error deleting product: %v", err)
//   } else {
// 	log.Println("Product deleted successfully")
//   }

//=================================================================

}

func getAllProductsGHandler(c *fiber.Ctx) error {
	products, err := getAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(products)
}

func getProductsHandler(c *fiber.Ctx) error { 
	//get product by ID from the URL parameter
	pid , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product , err := getProductByID(pid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.JSON(product)
}

func createProductsHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := createProduct(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(p)
}

func updateProductsHandler(c *fiber.Ctx) error {
	pid , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	p := new(Product)
	// Parse the request body into the Product struct
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := updateProduct(pid, p); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(p)
}

func deleteProductsHandler(c *fiber.Ctx) error {
	pid , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := deleteProduct(pid); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
}