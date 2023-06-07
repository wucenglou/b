package main

import (
	"fmt"
	"net/http"

	"b/models"
	"b/pkg/gredis"
	"b/pkg/logging"
	"b/pkg/setting"
	"b/pkg/util"
	"b/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	// go func() {
	// 	if err := s.ListenAndServe(); err != nil {
	// 		log.Printf("Listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := s.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }

	// log.Println("Server exiting")
}
