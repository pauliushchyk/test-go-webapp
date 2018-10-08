package main

import (
	"log"
	"path/filepath"
	"strings"

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

	r.NoRoute(controllers.NotFoundView)

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

	includes, e := filepath.Glob("templates/includes/**/*.tmpl")

	if e != nil {
		panic(e.Error())
	}

	for _, include := range includes {
		r.AddFromFiles(strings.Replace(include, "templates/includes/", "", -1), "templates/layouts/index.tmpl", include)
	}

	return r
}
