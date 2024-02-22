package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AndroidStudyOpenSource/africastalking-go/sms"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

const (
	username = "enter yours"                                                          //Your Africa's Talking Username
	apiKey   = "enter yours" //Production or Sandbox API Key
	env      = "Sandbox"                                                          // Choose either Sandbox or Production
)

// Customer model
type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
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
	db, err = sql.Open("postgres", "user=yours dbname=yours password=yours sslmode=disable")
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
	_, err := db.Exec("INSERT INTO customers (name, phone) VALUES ($1, $2)", customer.Name, customer.Phone)
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
	getPhoneByCustomerID(db, order.CustomerID)
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

// Function to fetch the phone value from the customers table based on customer_id
func getPhoneByCustomerID(db *sql.DB, customerID int) (string, error) {
	// Query to fetch the phone value from the customers table based on customer_id
	query := "SELECT phone FROM customers WHERE id = $1"

	// Execute the query
	row := db.QueryRow(query, customerID)

	// Initialize a variable to store the phone value
	var phone string

	// Scan the result into the phone variable
	err := row.Scan(&phone)
	if err != nil {
		// Handle the error
		log.Println("Error fetching phone:", err)
		return "", err
	}
	log.Println(phone)
	sendSms(phone)
	return phone, nil
}

func sendSms(phone string) {
	//Call the Gateway, and pass the constants here!
	smsService := sms.NewService(username, apiKey, env)

	//Send SMS - REPLACE Recipient and Message with REAL Values
	_, err := smsService.Send(phone, "order received", "") //Leave blank, "", if you don't have one
	if err != nil {
		log.Println(err)
	}
	log.Println("SMS Sent")
	//fmt.Println(smsResponse)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	

	// Routes
	e.POST("/customers", createCustomer)
	e.POST("/orders", createOrder)

	// Add KeyAuth middleware
	e.Use(KeyAuthMiddleware)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
