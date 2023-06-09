package stop

import (
	"transit-server/services/auth"

	"github.com/gin-gonic/gin"
)

func RouteStop(router *gin.RouterGroup) {
	stopGroup := router.Group("/stop")

	stopGroup.POST("/", HandleAddStop)
	stopGroup.GET("/", HandleGetStop)
	stopGroup.DELETE("/", HandleDeleteStop)
	stopGroup.GET("/nearest", auth.ValidateToken(), HandleGetNearestStop)
}
