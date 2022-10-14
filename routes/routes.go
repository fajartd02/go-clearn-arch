package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fajartd02/golang-devcamp/handler"
)

func SetupRoutes(db *gorm.DB, productHdl handler.ProductHandler, userHdl handler.UserHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// products
	r.GET("/products", productHdl.GetAll)
	r.GET("/products/:id", productHdl.GetSingle)
	r.PUT("/products/:id", productHdl.Update)
	r.POST("/products", productHdl.Create)
	r.DELETE("/products/:id", productHdl.Delete)

	// user
	r.POST("/user", userHdl.Create)
	r.GET("/users", userHdl.GetAll)

	return r
}
