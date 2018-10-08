package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundHandler returns NotFound page
func NotFoundHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "common/404.tmpl", nil)
}
