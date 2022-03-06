package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/log/strings"
	"github.com/core-go/log/zap"
	sv "github.com/core-go/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"

	"go-service/internal/app"
)

func main() {
	conf := app.Config{}
	er1 := config.Load(&conf, "configs/sql", "configs/config")
	if er1 != nil {
		panic(er1)
	}

	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(func(handler http.Handler) http.Handler {
		return mid.BuildContextWithMask(handler, MaskLog)
	})
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	er2 := app.Route(r, context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}
	/*
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	*/
	handler := cors.AllowAll().Handler(r)
	fmt.Println(sv.ServerInfo(conf.Server))
	server := sv.CreateServer(conf.Server, handler)
	if er3 := server.ListenAndServe(); er3 != nil {
		fmt.Println(er3.Error())
	}
}

func MaskLog(name, s string) string {
	return strings.Mask(s, 1, 6, "x")
}
