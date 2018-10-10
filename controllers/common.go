package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundView returns NotFound page
func NotFoundView(c *gin.Context) {
	c.HTML(http.StatusNotFound, "common/404.tmpl", nil)
}

// NoRoute redirects to get customers route
func NoRoute(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/customers")
}

// Ping returns ok if server is ok
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
