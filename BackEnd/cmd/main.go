package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func init() {
	e := echo.New()
	s := http.Server{
		Addr:        ":8080",
		Handler:     e,
		ReadTimeout: 30 * time.Second,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
