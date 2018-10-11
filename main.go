package main

import (
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/pauliushchyk/test-go-webapp/config"
	"github.com/pauliushchyk/test-go-webapp/controllers"
	"github.com/pauliushchyk/test-go-webapp/models"
	"github.com/pauliushchyk/test-go-webapp/shared"
)

func main() {
	db := shared.OpenConnection()

	defer shared.CloseConnection()

	db.AutoMigrate(&models.Customer{})

	r := setUpRouter()

	log.Fatal(r.Run(config.Get("http.port")))
}

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true
	r.HTMLRender = createRender()
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")

	r.NoRoute(controllers.NoRoute)

	r.GET("/ping", controllers.Ping)

	csr := r.Group("/customers")
	{
		csr.GET("", controllers.GetCustomersView)

		csr.GET("/create", controllers.CreateCustomerView)
		csr.POST("/create", controllers.CreateCustomer)
	}

	cr := r.Group("/customer")
	{
		cr.GET("/:id", controllers.GetCustomerView)

		cr.GET("/:id/update", controllers.UpdateCustomerView)
		cr.POST("/:id/update", controllers.UpdateCustomer)

		cr.GET("/:id/delete", controllers.DeleteCustomerView)
		cr.POST("/:id/delete", controllers.DeleteCustomer)
	}

	return r
}

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("common/404.tmpl", "templates/layouts/index.tmpl", "templates/includes/common/404.tmpl")
	r.AddFromFiles("customer/delete.tmpl", "templates/layouts/index.tmpl", "templates/includes/customer/delete.tmpl")
	r.AddFromFiles("customer/update.tmpl", "templates/layouts/index.tmpl", "templates/includes/customer/update.tmpl")
	r.AddFromFiles("customer/view.tmpl", "templates/layouts/index.tmpl", "templates/includes/customer/view.tmpl")
	r.AddFromFiles("customers/create.tmpl", "templates/layouts/index.tmpl", "templates/includes/customers/create.tmpl")
	r.AddFromFiles("customers/view.tmpl", "templates/layouts/index.tmpl", "templates/includes/customers/view.tmpl")

	return r
}
