package router

import (
	"net/http"

	"github.com/asynccnu/library_service_v2/handler/library"
	"github.com/asynccnu/library_service_v2/handler/sd"
	"github.com/asynccnu/library_service_v2/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	api := g.Group("/api/lib/v2")
	// api.Use(middleware.AuthMiddleware())
	{
		api.GET("/search", library.Search)
		api.GET("/book/:id", library.GetBookInfo)
		api.GET("/stars", library.ListStars)
		api.GET("/my_books", library.ListBorrowedBooks)
		api.POST("/book/:id/renew", library.Renew)
		api.POST("/book/:id", library.Star)
		api.DELETE("/book/:id", library.Unstar)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
