package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	sv "github.com/core-go/core"
	"github.com/core-go/core/cors"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/log/strings"
	"github.com/core-go/log/zap"
	"github.com/gorilla/mux"
	"net/http"

	"go-service/internal/app"
)

func main() {
	var conf app.Config
	err := config.Load(&conf, "configs/sql", "configs/config")
	if err != nil {
		panic(err)
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

	err = app.Route(r, context.Background(), conf)
	if err != nil {
		panic(err)
	}
	c := cors.New(conf.Allow)
	handler := c.Handler(r)
	fmt.Println(sv.ServerInfo(conf.Server))
	sv.StartServer(conf.Server, handler)
}

func MaskLog(name, s string) string {
	return strings.Mask(s, 1, 6, "x")
}
