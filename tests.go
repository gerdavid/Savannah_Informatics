package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreateCustomer(t *testing.T) {

	// Test with valid input
	c := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"John Doe","phone":"+1234567890"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c.NewContext(req, rec)
	status := http.StatusOK
	if err := createCustomer(c.AcquireContext()); err != nil || status != http.StatusOK {
		t.Errorf("createCustomer() error = %v, status = %v, wantErr %v, status %v", err, status, nil, http.StatusOK)
	}

	// Test with missing required field
	req = httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"phone":"+1234567890"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = echo.New()
	c.NewContext(req, rec)
	status = http.StatusBadRequest
	if err := createCustomer(c.AcquireContext()); err == nil || status != http.StatusBadRequest {
		t.Errorf("createCustomer() error = %v, status = %v, wantErr %v, status %v", err, status, true, http.StatusBadRequest)
	}

	// Test with invalid phone number
	req = httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"John Doe","phone":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = echo.New()
	c.NewContext(req, rec)

	if err := createCustomer(c.AcquireContext()); err == nil || status != http.StatusBadRequest {
		t.Errorf("createCustomer() error = %v, status = %v, wantErr %v, status %v", err, status, true, http.StatusBadRequest)
	}
}

func TestCreateOrder(t *testing.T) {

	// Test with valid input
	c := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"customerID":"1","items":[{"productID":"1","quantity":2}]}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c.NewContext(req, rec)
	status := http.StatusOK
	if err := createOrder(c.AcquireContext()); err != nil || status != http.StatusOK {
		t.Errorf("createOrder() error = %v, status = %v, wantErr %v, status %v", err, status, nil, http.StatusOK)
	}

	// Test with invalid customer ID
	req = httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"customerID":"invalid","items":[{"productID":"1","quantity":2}]}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = echo.New()
	c.NewContext(req, rec)

	if  err := createOrder(c.AcquireContext()); err == nil || status != http.StatusBadRequest {
		t.Errorf("createOrder() error = %v, status = %v, wantErr %v, status %v", err, status, true, http.StatusBadRequest)
	}

	// Test with missing required field
	req = httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"customerID":"1"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = echo.New()
	c.NewContext(req, rec)

	if  err := createOrder(c.AcquireContext()); err == nil || status != http.StatusBadRequest {
		t.Errorf("createOrder() error = %v, status = %v, wantErr %v, status %v", err, status, true, http.StatusBadRequest)
	}
}
