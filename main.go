package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gomock/api"
	"gomock/middleware"
	"gomock/resource"
	"os"
)

func main() {
	r := gin.New()

	r.Use(middleware.DefaultStructedLogger())
	r.Use(gin.Recovery())

	resource.NewAppContext()

	defer resource.AppCtx.DestroyAppCtx()

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
