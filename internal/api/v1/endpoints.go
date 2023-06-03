package v1

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tailwarden/komiser/handlers"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
)

func Endpoints(ctx context.Context, telemetry bool, analytics utils.Analytics, db *bun.DB, cfg models.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Default())

	api := handlers.NewApiHandler(ctx, telemetry, analytics, db, cfg)

	router.POST("/resources/search", api.FilterResourcesHandler)
	router.POST("/resources/tags", api.BulkUpdateTagsHandler)
	router.POST("/resources/:id/tags", api.UpdateTagsHandler)
	router.GET("/resources/export-csv", api.DownloadInventoryCSV)
	router.GET("/resources/export-csv/:viewId", api.DownloadInventoryCSVForView)

	router.GET("/views", api.ListViewsHandler)
	router.POST("/views", api.NewViewHandler)
	router.PUT("/views/:id", api.UpdateViewHandler)
	router.DELETE("/views/:id", api.DeleteViewHandler)
	router.POST("/views/:id/resources/hide", api.HideResourcesFromViewHandler)
	router.POST("/views/:id/resources/unhide", api.UnhideResourcesFromViewHandler)
	router.GET("/views/:id/hidden/resources", api.ListHiddenResourcesHandler)
	router.GET("/views/:id/alerts", api.ListViewAlertsHandler)

	router.GET("/regions", api.ListRegionsHandler)
	router.GET("/providers", api.ListProvidersHandler)
	router.GET("/services", api.ListServicesHandler)
	router.GET("/accounts", api.ListAccountsHandler)
	router.GET("/stats", api.StatsHandler)
	router.POST("/stats/search", api.FilterStatsHandler)

	router.GET("/global/stats", api.DashboardStatsHandler)
	router.POST("/global/resources", api.ResourcesBreakdownStatsHandler)
	router.GET("/global/locations", api.LocationBreakdownStatsHandler)
	router.POST("/costs/explorer", api.CostBreakdownHandler)

	router.GET("/slack", api.IsSlackEnabledHandler)
	router.POST("/alerts", api.NewAlertHandler)
	router.PUT("/alerts/:id", api.UpdateAlertHandler)
	router.DELETE("/alerts/:id", api.DeleteAlertHandler)
	router.POST("/alerts/test", api.TestEndpointHandler)

	router.GET("/telemetry", api.TelemetryHandler)

	router.NoRoute(gin.WrapH(http.FileServer(assetFS())))

	return router
}
