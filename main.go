package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/log/strings"
	"github.com/core-go/log/zap"
	sv "github.com/core-go/service"
	"github.com/core-go/service/cors"
	"github.com/gorilla/mux"
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
	c := cors.New(conf.Allow)
	handler := c.Handler(r)
	fmt.Println(sv.ServerInfo(conf.Server))
	sv.StartServer(conf.Server, handler)
}

func MaskLog(name, s string) string {
	return strings.Mask(s, 1, 6, "x")
}
