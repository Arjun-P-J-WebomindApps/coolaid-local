package app

import (
	"github.com/gin-gonic/gin"
	config "github.com/webomindapps-dev/coolaid-backend/config"
	"github.com/webomindapps-dev/coolaid-backend/internal/api/graphql"
	routeH "github.com/webomindapps-dev/coolaid-backend/internal/api/http"
)

func mapUrls(graphqlServer *graphql.Server, httpServer *routeH.Server) {
	router.GET("/healthz", httpServer.Handler.RouteHandler.HealthZ)
	router.GET("/search/healthz", httpServer.Handler.RouteHandler.SearchHealthZ)

	cldV1API := router.Group("/api/v1")
	mapPublicApiRoutes(cldV1API, graphqlServer, httpServer)

	// Playground (dev only)
	if config.App.Name != "production" {
		internalAPI := router.Group("/api/internal")
		mapInternalApiRoutes(internalAPI, graphqlServer)
	}
}

// Public Facing API
// PRODUCTION
func mapPublicApiRoutes(apiRoute *gin.RouterGroup, graphqlServer *graphql.Server, httpServer *routeH.Server) {

	//Graphql Endpoint
	apiRoute.POST("/query",
		// controllers.AuthContextMiddleware(),
		graphql.GinHandler(graphqlServer),
	)

	apiRoute.GET("/search", httpServer.Handler.SearchHandler.Suggestions)

	// apiRoute.GET("/download/csv", csv.DownloadHandler)      // DB-backed download
	// apiRoute.GET("/download/csv/sample", csv.SampleHandler) // Schema-only sample
	// apiRoute.POST("/upload/csv", csv.UploadHandler)         // Upload + validate

}

// Internal Functions
// DEVELOPMENT ONLY
func mapInternalApiRoutes(apiRoute *gin.RouterGroup, _ *graphql.Server) {

	apiRoute.GET("/playground", graphql.PlaygroundHandler("/api/v1/query"))
}
