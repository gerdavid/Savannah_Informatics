package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// Customer model
type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// Order model
type Order struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	Item       string `json:"item"`
	Amount     int    `json:"amount"`
	Time       string `json:"time"`
}

// Database connection
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=savannah password=davinski sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

// Handler to create a new customer
func createCustomer(c echo.Context) error {
	customer := new(Customer)
	if err := c.Bind(customer); err != nil {
		return err
	}
	_, err := db.Exec("INSERT INTO customers (name, code) VALUES ($1, $2)", customer.Name, customer.Code)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, customer)
}

// Handler to create a new order
func createOrder(c echo.Context) error {
	order := new(Order)
	if err := c.Bind(order); err != nil {
		return err
	}
	_, err := db.Exec("INSERT INTO orders (customer_id, item, amount, time) VALUES ($1, $2, $3, $4)", order.CustomerID, order.Item, order.Amount, order.Time)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, order)
}

// Middleware function for API key authentication
func KeyAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the API key from the request header
		apiKey := c.Request().Header.Get("X-API-Key")

		// Define your API key here
		validAPIKey := "your_api_key"

		// Validate the API key
		if apiKey != validAPIKey {
			// API key is invalid
			return echo.ErrUnauthorized
		}

		// API key is valid, proceed to the next handler
		return next(c)
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
	// 	// Implement your key authentication logic here
	// 	return key == os.Getenv("API_KEY"), nil
	// }))

	// Routes
	e.POST("/customers", createCustomer)
	e.POST("/orders", createOrder)

	// Add KeyAuth middleware
	e.Use(KeyAuthMiddleware)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
