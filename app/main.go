package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		message := `
		<center>==================================================================================
		<p><h2>sianggota</h2>
		<p>Sistem Informasi Anggota
		<p>==================================================================================</center>`
		return c.HTML(200, message)
	})
	go func() {
		address := fmt.Sprintf("localhost:%d", 3000)

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
