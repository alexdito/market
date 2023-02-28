package providers

import (
	"application/internal/app/routes"
	"github.com/gin-gonic/gin"
)

func RouteBuild(r *gin.Engine, controllers *ControllerDI) {
	api := r.Group("/api")
	routes.Navigation(api, controllers.NavigationController)
	routes.Card(api, controllers.CardController)
}
