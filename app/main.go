package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"api.sianggota.com/api"
	"api.sianggota.com/config"
	"api.sianggota.com/database"
	"api.sianggota.com/database/migration"
	"api.sianggota.com/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//init config
	config := config.GetConfig()
	//connect database
	db := database.Connect(config.Database)
	if d, ok := db.DB(); ok != nil {
		if err := d.Ping(); err != nil {
			defer d.Close()
		}
	}
	e := echo.New()
	//set middleware
	middlewares.New(e)
	e.GET("/", func(c echo.Context) error {
		message := `
		<center>==================================================================================
		<p><h2>sianggota</h2>
		<p>Sistem Informasi Anggota
		<p>==================================================================================</center>`
		return c.HTML(200, message)
	})
	migration.Migrate()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"hello": "world"})
	})

	//
	// init api
	//

	api.New(*e)
	e.GET("/list-routes", func(c echo.Context) error {
		return c.JSON(200, e.Routes())
	})

	go func() {
		address := fmt.Sprintf("%s:%d", config.App.Host, config.App.Port)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
