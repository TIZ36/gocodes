package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gomock/api"
	"gomock/common"
	"gomock/middleware"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = "./config/dev.yaml"

func main() {
	r := gin.New()

	r.Use(middleware.DefaultStructedLogger())
	r.Use(gin.Recovery())

	config := loadConfig(configFile)

	log.Info().Msg(fmt.Sprintf("config is %v", config))
	common.NewAppContext(config)

	defer common.AppCtx.DestroyAppCtx()

	r.GET("/chat", api.GetRouter)

	r.POST("/gomock/v1/api/user/basic/new", api.PostUserCreate)

	log.Print("prepare to listen on 8888")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	gin.DefaultWriter = log.Logger

	err := errors.New("mock err here")
	log.Error().Stack().Err(err).Msg("a info log")
	r.Run(":8888")

	fmt.Println("hhh")
}

func loadConfig(path string) common.Config {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Error().Msg("read config error")
		panic(err)
	}

	var config common.Config

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		log.Error().Msg("invalid yaml format")
		panic(err)
	}

	return config
}
