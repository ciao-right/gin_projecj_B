package main

import (
	"fmt"
	"gin_project_B/pkg/setting"
	"gin_project_B/routers"
	"net/http"
)

func main() {
	//r := gin.Default()
	//r.Run()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
