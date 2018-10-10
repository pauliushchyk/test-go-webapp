package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/pauliushchyk/test-go-webapp/models"

	"github.com/pauliushchyk/test-go-webapp/shared"

	"github.com/stretchr/testify/assert"
)

var r = setUpRouter()

var db = shared.OpenTestConnection()

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestNotFoundRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/invalid-route", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func TestGetCustomerRoute(t *testing.T) {
	customer := &models.Customer{
		FirstName: "John",
		LastName:  "Doe",
		BirthDate: time.Now(),
		Gender:    "Male",
		Email:     "email@mail.com",
		Address:   "address",
	}

	db.Create(&customer)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customer/"+fmt.Sprint(customer.ID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	db.Delete(&customer, customer.ID)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/customer/"+fmt.Sprint(customer.ID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateCustomerRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customers/create", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCustomerRoute(t *testing.T) {
	customer := &models.Customer{
		FirstName: "John",
		LastName:  "Doe",
		BirthDate: time.Now(),
		Gender:    "Male",
		Email:     "email@mail.com",
		Address:   "address",
	}

	db.Create(&customer)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customer/"+fmt.Sprint(customer.ID)+"/update", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/customer/"+fmt.Sprint(customer.ID)+"/update", strings.NewReader("first_name=Jane&last_name=Doe&birth_date="+time.Now().Format("2006-01-02")+"&gender=female&email=jane@doe.com&address=Jane's address"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	var updatedCustomer models.Customer

	db.First(&updatedCustomer, customer.ID)

	assert.Equal(t, "Jane", updatedCustomer.FirstName)
	assert.Equal(t, "female", updatedCustomer.Gender)
	assert.Equal(t, "jane@doe.com", updatedCustomer.Email)
	assert.Equal(t, "Jane's address", updatedCustomer.Address)
}

func TestDeleteCustomerRoute(t *testing.T) {
	customer := &models.Customer{
		FirstName: "John",
		LastName:  "Doe",
		BirthDate: time.Now(),
		Gender:    "Male",
		Email:     "email@mail.com",
		Address:   "address",
	}

	db.Create(&customer)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customer/"+fmt.Sprint(customer.ID)+"/delete", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/customer/"+fmt.Sprint(customer.ID)+"/delete", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/customer/"+fmt.Sprint(customer.ID)+"/delete", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
