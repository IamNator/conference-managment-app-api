package main

import (
	"conference/handler"
	"conference/routes"
	"conference/service"
	"conference/storage"
	"errors"
	"github.com/joho/godotenv"
	log "github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := log.New(os.Stdout).With().Str("app", "conf_mgmt_sys").Logger()

	//export env from .env file
	if er := godotenv.Load(); er != nil {
		logger.Warn().Err(er).Msg("no env file found")
	}

	port := os.Getenv("PORT")
	store, er := storage.New()
	if er != nil {
		logger.Fatal().Err(er).Msg("unable to connect to database")
		return
	}
	//if er := store.RunMigration(); er != nil { //
	//	logger.Fatal().Err(er).Msg("database migration failed")
	//}
	userSrv := service.NewUserService(store)
	confSrv := service.NewConferenceService(store)
	handlers := handler.NewHandler(userSrv, confSrv)

	erChan := make(chan error)
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if er := routes.Run(handlers, port); er != nil {
			erChan <- er
		}
	}()

	select {
	case er := <-erChan:
		logger.Fatal().Err(er).Msg("server shut down")
	case <-quit:
		logger.Info().Err(errors.New("shutdown signal received")).Msg("server shut down")
	}

}
