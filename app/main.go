package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"api.sianggota.com/config"
	"api.sianggota.com/database"
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
	e.GET("/", func(c echo.Context) error {
		message := `
		<center>==================================================================================
		<p><h2>sianggota</h2>
		<p>Sistem Informasi Anggota
		<p>==================================================================================</center>`
		return c.HTML(200, message)
	})
	type greeting struct {
		ID    string `gorm:"default:random_string(8)"` // db func
		Hello string
	}
	db.AutoMigrate(&greeting{})
	e.GET("/ping", func(c echo.Context) error {
		greet := greeting{
			Hello: "world",
		}
		result := db.Create(&greet)
		if result.Error != nil {
			return c.String(500, result.Error.Error())
		}

		return c.JSON(200, greet)
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
