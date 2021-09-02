// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"adolesce/internal/conf"
	"adolesce/internal/pkg/app"
	"adolesce/internal/server"
	"adolesce/internal/service"
	"adolesce/pkg/log"
)

// Injectors from wire.go:

func initTimer(configs *conf.Configs, logger log.Logger) (*app.App, func(), error) {
	exampleTimerHandler := service.NewExampleTimerService(logger)
	cron, cleanup, err := server.NewTimerServer(logger, exampleTimerHandler)
	if err != nil {
		return nil, nil, err
	}
	appApp, cleanup2, err := newTimerServer(logger, cron)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return appApp, func() {
		cleanup2()
		cleanup()
	}, nil
}
