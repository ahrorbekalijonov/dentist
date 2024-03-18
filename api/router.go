package api

import (
	_ "github.com/dentist/api/docs" // swag

	v1 "github.com/dentist/api/v1"
	"github.com/dentist/config"
	"github.com/dentist/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type RoutOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

//New...
// @Title           Dentist
// @Version         1.0
// @Description     Dentist-backend
func New(opts RoutOptions) *gin.Engine {
	router := gin.Default()
	
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Storage: opts.Storage,
		Cfg: opts.Cfg,
	})

	router.Use(gin.Recovery())

	v1 := router.Group("/v1")

	//client...
	v1.POST("/client", handlerV1.CreateClient)
	v1.GET("/client", handlerV1.GetClient)
	v1.PUT("/client", handlerV1.UpdateClient)
	v1.DELETE("/client", handlerV1.DeleteClient)
	v1.GET("/clients", handlerV1.GetAllClients)
	v1.GET("/count", handlerV1.GetAllClientsCount)
	v1.GET("/search", handlerV1.SearchClients)

	//appointment...
	v1.POST("/appointment", handlerV1.CreateAppointment)
	v1.GET("/appointment", handlerV1.GetAppointment)
	v1.PUT("/appointment", handlerV1.UpdateAppointment)
	v1.DELETE("/appointment", handlerV1.DeleteAppointment)
	v1.GET("/appointments", handlerV1.GetAllAppointments)
	v1.GET("/appointmentsdate", handlerV1.GetAppointmentsWithDate)
	v1.GET("/appointmentid", handlerV1.GetAppointmentWithClientId)
	v1.POST("/appointmentnew", handlerV1.CreateAppointmentWithClient)


	v1.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
