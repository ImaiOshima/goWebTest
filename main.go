package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWebTest/global"
	"goWebTest/internal/model"
	"goWebTest/internal/routers"
	setting "goWebTest/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init(){
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v",err)
	}
}

func init(){
	err := setupDBEngine()
	if err != nil{
		log.Fatalf("init .setupDBEngine err:%s",err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting,err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server",&global.ServerSetting)
	fmt.Println(*global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}