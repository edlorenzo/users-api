package main

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/edlorenzo/users-api/config"
	_ "github.com/edlorenzo/users-api/docs"
	"github.com/edlorenzo/users-api/handler"
	"github.com/edlorenzo/users-api/integration/redis"
	"github.com/edlorenzo/users-api/logs"
	"github.com/edlorenzo/users-api/router"
	"github.com/edlorenzo/users-api/store"
	"github.com/edlorenzo/users-api/utils"
)

// @description User List API
// @title User List API

// @BasePath /api

// @schemes http https
// @produce application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logs.InitLoggers()

	cfg, err := config.InitConfig()
	if err != nil {
		logs.Logs.Errorf("[Error] Init Config: %v", err)
	} else {
		logs.Logs.Infoln("[Info] Init Config: Success")
	}

	client, err := redis.InitRedis(cfg.RedisConfig)
	if err != nil {
		logs.Logs.Errorf("[Error] Init Redis: %v", err)
	} else {
		logs.Logs.Infoln("[Info] Init Redis: Success")
	}

	r := router.New()
	r.Get("/swagger/*", swagger.HandlerDefault)

	status, err := utils.HttpConnectionStatus(cfg.GithubConfig)
	if err != nil {
		logs.Logs.Errorf("[Error] Github API: %v", err)
	} else {
		logs.Logs.Infof("[Info] Github API Status: %t", status)
	}
	us := store.NewUserStore(client, cfg.GithubConfig.UsersURI)
	h := handler.NewHandler(us)
	h.Register(r)
	err = r.Listen(":5007")
	if err != nil {
		fmt.Printf("%v", err)
		logs.Logs.Errorf("listener error: %v \n", err)
	}
}
