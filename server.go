package main

import (
	"example/golang-gin-poc/controller"
	"example/golang-gin-poc/dto"
	"example/golang-gin-poc/middlewares"
	"example/golang-gin-poc/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	// setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	/*
		server.Use(gin.Recovery(), middlewares.Logger(),
			middlewares.BasicAuth(), gindump.Dump())
	*/
	server.Use(gin.Recovery(), middlewares.Logger())

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, &dto.JWT{
				Token: token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, &dto.Response{
				Message: "Not Authorized",
			})
		}
	})

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, &dto.Response{
					Message: err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, &dto.Response{
					Message: "Success!",
				})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, &dto.Response{
					Message: err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, &dto.Response{
					Message: "Success!",
				})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, &dto.Response{
					Message: err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, &dto.Response{
					Message: "Success!",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
