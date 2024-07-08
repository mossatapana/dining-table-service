package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mossatapana/dining-table-service/handlers"
	"github.com/mossatapana/dining-table-service/services"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:8080", "127.0.0.1:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{"*"},
	}))

	tableService := services.NewTableService()
	tableHandler := handlers.NewTableHandlers(tableService)

	v1 := e.Group("/api/v1")
	v1Table := v1.Group("/table")
	v1Table.POST("/initial", tableHandler.Initial)
	v1Table.POST("/reserve", tableHandler.Reserve)
	v1Table.POST("/cancel", tableHandler.Cancel)

	e.Logger.Fatal(e.Start(":8080"))
}
