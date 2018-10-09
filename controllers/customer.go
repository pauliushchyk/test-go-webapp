package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pauliushchyk/test-go-webapp/models"

	"github.com/pauliushchyk/test-go-webapp/shared"

	"github.com/gin-gonic/gin"
)

// GetCustomersView returns all customers view
func GetCustomersView(c *gin.Context) {
	page, e := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)

	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	by := c.Query("by")
	order := c.Query("order")
	search := c.Query("search")

	db := shared.GetConnection()

	var customers []models.Customer
	var count int

	if db.Where("LOWER(first_name) LIKE LOWER(?)", "%"+search+"%").Or("last_name LIKE ?", "%"+search+"%").Order(strings.TrimSpace(strings.Join([]string{by, order}, " "))).Limit(10).Offset((page-1)*10).Find(&customers).Limit(-1).Offset(-1).Count(&count).Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.HTML(http.StatusOK, "customers/view.tmpl", gin.H{
		"title":     "Customers",
		"customers": customers,
		"count":     count,
		"search":    search,
	})
}

// GetCustomerView returns customer view
func GetCustomerView(c *gin.Context) {
	id, e := strconv.ParseUint(c.Param("id"), 10, 32)

	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := shared.GetConnection()

	var customer models.Customer

	if db.First(&customer, id).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "customer/view.tmpl", gin.H{
		"title":    customer.FirstName + " " + customer.LastName,
		"customer": customer,
	})
}

// CreateCustomerView returns create customer view
func CreateCustomerView(c *gin.Context) {
	c.HTML(http.StatusOK, "customers/create.tmpl", gin.H{
		"title": "Create customer",
	})
}

// CreateCustomer creates customer
func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if e := c.Bind(&customer); e != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	db := shared.GetConnection()

	db.Create(&customer)

	c.Redirect(http.StatusMovedPermanently, "/customer/"+fmt.Sprint(customer.ID))
}

// UpdateCustomerView returns update customer view
func UpdateCustomerView(c *gin.Context) {
	id, e := strconv.ParseUint(c.Param("id"), 10, 32)

	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := shared.GetConnection()

	var customer models.Customer

	if db.First(&customer, id).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "customer/update.tmpl", gin.H{
		"title":    "Update " + customer.FirstName + " " + customer.LastName,
		"customer": customer,
	})
}

// UpdateCustomer updates customer
func UpdateCustomer(c *gin.Context) {
	id, e := strconv.ParseUint(c.Param("id"), 10, 32)

	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var model models.Customer

	if e := c.Bind(&model); e != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	db := shared.GetConnection()

	var customer models.Customer

	if db.First(&customer, id).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Model(customer).Updates(model)

	c.Redirect(http.StatusMovedPermanently, "/customer/"+fmt.Sprint(customer.ID))
}

// DeleteCustomerView returns delete customer view
func DeleteCustomerView(c *gin.Context) {
	id, e := strconv.ParseUint(c.Param("id"), 10, 32)

	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := shared.GetConnection()

	var customer models.Customer

	if db.First(&customer, id).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "customer/delete.tmpl", gin.H{
		"title":    "Delete " + customer.FirstName + " " + customer.LastName,
		"customer": customer,
	})
}

// DeleteCustomer deletes customer
func DeleteCustomer(c *gin.Context) {
	id, e := strconv.ParseUint(c.Param("id"), 10, 32)

	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := shared.GetConnection()

	var customer models.Customer

	if db.Delete(&customer, id).Error != nil {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/customers")
}
