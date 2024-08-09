package main

import (
	"context"
	"fmt"
	"github.com/core-go/core/config"
	"github.com/core-go/core/cors"
	svr "github.com/core-go/core/server"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/log/strings"
	"github.com/core-go/log/zap"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	err := config.Load(&cfg, "configs/sql", "configs/config")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	log.Initialize(cfg.Log)
	r.Use(func(handler http.Handler) http.Handler {
		return mid.BuildContextWithMask(handler, MaskLog)
	})
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(cfg.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	err = app.Route(r, context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	c := cors.New(cfg.Allow)
	handler := c.Handler(r)
	fmt.Println(svr.ServerInfo(cfg.Server))
	svr.StartServer(cfg.Server, handler)
}

func MaskLog(name, s string) string {
	return strings.Mask(s, 1, 6, "x")
}
