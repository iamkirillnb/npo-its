package main

import (
	"flag"
	"github.com/iamkirillnb/Packages/pkg/logger"
	"npo-its/internal"
	"npo-its/internal/handlers"
	"npo-its/internal/repos"
	"npo-its/pkg/pg"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "dev.yaml", "string Path to dev.yaml file")
}

func main() {

	flag.Parse()
	// logger
	log := logger.NewLogger()
	log.Println("app started")

	// config
	conf := internal.GetConfig(cfgPath)

	//db
	postgres := pg.NewDB(&conf.PgDb, &log)

	//repo
	repo := repos.NewDB(postgres)

	// handler
	handler := handlers.NewHandler(&conf.ServerCfg, &log, repo)
	handler.Start()
}
