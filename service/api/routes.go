package api

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/service/api/handlers"
	"github.com/Confialink/wallet-kyc/service/api/middlewares"
	"github.com/Confialink/wallet-pkg-errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) initRoutes() {

	mwCors := middlewares.Cors(server.Configs.Cors)
	mwAuth := middlewares.Auth()
	mwAdminOrRoot := middlewares.AdminOrRoot()

	apiGroup := server.Gin.Group("/kyc", mwCors)
	{
		apiGroup.Use(
			gin.Recovery(),
			gin.Logger(),
			errors.ErrorHandler(connection.Logger),
		)

		apiGroup.GET("/health-check", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		privateGroup := apiGroup.Group("/private", mwAuth)
		{
			v1Group := privateGroup.Group("/v1")
			{
				tierController := handlers.NewTierController()
				requirementController := handlers.NewRequirementController()
				requestController := handlers.NewRequestController()
				tierSettingsController := handlers.NewTierSettingsController()

				// Section user routes
				tiers := v1Group.Group("/tiers")
				{
					tiers.GET("", tierController.ListForUser)
					tiers.GET("/current", tierController.Current)
				}

				tier := v1Group.Group("/tier")
				{
					tier.GET("/:tierId", tierController.Get)
				}

				requirements := v1Group.Group("/requirement/:requirementId")
				{
					requirements.PUT("", requirementController.Update)
				}

				requests := v1Group.Group("/requests")
				{
					requests.POST("", requestController.Create)
				}

				// Section admin routes
				admin := v1Group.Group("/admin", mwAdminOrRoot)
				{
					requests := admin.Group("/requests")
					{
						requests.GET("", requestController.List)
					}

					requirement := admin.Group("/requirement/:requirementId/user/:userId")
					{
						requirement.PUT("", requirementController.AdminUpdate)
						requirement.PUT("/update-status", requirementController.UpdateStatus)
					}

					request := admin.Group("/request/:requestId")
					{
						request.PUT("/update-status", requestController.UpdateStatus)
					}

					tiers := admin.Group("/tiers")
					{
						tiers.GET("/user/:userId", tierController.ListForAdmin)
					}

					countries := admin.Group("/countries")
					{
						countries.GET("", tierSettingsController.ListCountries)
					}

					country := admin.Group("/country/:country_code")
					{
						country.GET("/tiers", tierSettingsController.ListTiers)
					}

					tier := admin.Group("/tier/:tierId")
					{
						tier.GET("", tierSettingsController.GetTier)
						tier.PUT("", tierSettingsController.UpdateTier)
					}
				}

			}
		}
	}

	// Handle OPTIONS request
	server.Gin.OPTIONS("/*cors", mwCors, func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Handle NOT FOUND request
	server.Gin.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
