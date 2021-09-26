package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/fauzanmh/olp-user/config"
	_handler "github.com/fauzanmh/olp-user/handler"
	_mysqlRepo "github.com/fauzanmh/olp-user/repository/mysql"
	_usecaseCourse "github.com/fauzanmh/olp-user/usecase/course"
	_usecaseStatistic "github.com/fauzanmh/olp-user/usecase/statistic"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	appInit "github.com/fauzanmh/olp-user/init"
	appMiddleware "github.com/fauzanmh/olp-user/middleware"
	_ "github.com/spf13/viper/remote"
	echoSwagger "github.com/swaggo/echo-swagger"
	log "go.uber.org/zap"
)

var cfg *appInit.Config

func init() {
	// Start pre-requisite app dependencies
	cfg = appInit.StartAppInit()
}

func main() {
	// echo
	e := echo.New()

	// mutex

	// timeout
	timeoutContext := time.Duration(cfg.Context.Timeout) * time.Second

	// init database
	mysqlDB, err := appInit.ConnectToMysqlServer(cfg)
	if err != nil {
		log.S().Fatal(err)
	}

	// * repository
	mysqlRepo := _mysqlRepo.NewRepository(mysqlDB)

	// * usecase
	// course usecase
	courseUsecase := _usecaseCourse.NewCourseUseCase(cfg, mysqlRepo)
	// statistic usecase
	statisticUsecase := _usecaseStatistic.NewStatisticUseCase(cfg, mysqlRepo)

	// Middleware
	e.Use(appMiddleware.EchoCORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(appMiddleware.DumpRequestResponse))
	config.SetEchoErrorDefault(e)
	// End of middleware

	// Grouping Routes
	routerAPI := e.Group("/api")
	// swagger route
	routerAPI.GET("/swagger/*", echoSwagger.WrapHandler)
	// course routes
	_handler.NewCourseHandler(routerAPI, courseUsecase)
	// statistic routes
	_handler.NewStatisticHandler(routerAPI, statisticUsecase)

	go runHTTPHandler(e, cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func runHTTPHandler(e *echo.Echo, cfg *appInit.Config) {
	if err := e.Start(cfg.API.HTTP.Port); err != nil {
		fmt.Println("shutting down the server")
	}
}
