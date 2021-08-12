package web

import "C"
import (
	"github.com/gin-gonic/gin"
)

// Router returns a http router with endpoints set
func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/ducks/:id", getDuck)
	router.GET("/ducks", getDucks)
	router.POST("/ducks", addDuck)
	router.GET("/ducks-match/:id", getDuckMatch)

	return router
}
