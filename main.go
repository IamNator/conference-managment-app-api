package main

import (
	"conference/handler"
	"conference/routes"
	"conference/service"
	"conference/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := os.Getenv("PORT")
	store := storage.New()
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
		log.Fatalln(er)
	case <-quit:
		log.Println("server shut down")
	}

}
