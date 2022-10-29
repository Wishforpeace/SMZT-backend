package router

import (
	"SMZT/handler"
	"SMZT/handler/jottings"
	"SMZT/handler/sd"
	"SMZT/handler/user"
	"SMZT/pkg/errno"
	"SMZT/router/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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
		handler.SendError(c, errno.ErrIncorrectAPIRoute, nil, "", "")
	})

	// swagger API doc
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Jotting
	jotting := g.Group("api/v1/jotting")
	{
		jotting.GET("", jottings.GetJotting)
		jotting.POST("", jottings.PostJotting)
	}
	// user 模块
	userRouter := g.Group("api/v1/user")
	{
		userRouter.POST("/login", user.Login)

	}
	// The health check Fandlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
