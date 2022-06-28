package routes

import (
	"resume/controllers"
	"resume/middlewares"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("views/*.html")

	r.GET("/", controllers.IndexGetHandler)
	r.POST("/", controllers.ContactPostHandler)

	public := r.Group("/api")
	{
		public.GET("/v1", controllers.PingGetHandler)
		public.GET("/v1/about", controllers.AboutGetHandler)
		public.GET("/v1/experience", controllers.ExperiencesGetHandler)
		public.GET("/v1/skill", controllers.SkillGetHandler)
		public.POST("/v1/demo/register", controllers.RegisterPostHandler)
		public.POST("/v1/demo/login", controllers.LoginPostHandler)
		public.POST("/v1/contact", controllers.ContactPostHandlerApi)
		// CoinGecko
		public.GET("/v1/demo/thirdparty/coingecko/coins/markets", controllers.GeckoGetHandler)
		// Bitkub
		public.GET("/v1/demo/thirdparty/bitkub/market/ticker", controllers.BKTickerGetHandler)
	}

	private := r.Group("/api")
	{
		private.Use(middlewares.JwtAuthMiddleware())
		private.GET("/v1/demo/user/profile", controllers.ProfileGetHandler)
		private.PATCH("/v1/demo/user/profile", controllers.PasswordPatchHandler)
		private.DELETE("/v1/demo/user/logout", controllers.LogoutDelHandler)

		private.GET("/v1/demo/account", controllers.AccountGetHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
