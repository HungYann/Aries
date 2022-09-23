package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"net/http"
	"server/internal/config"
	"server/internal/service"
	"time"
)

const VERSION = "achilles-20220922-version1"

func initLog() {
	log.SetLevel(log.InfoLevel)

	achillesLog := config.GetLog()
	// 配置日志每隔7天轮转一个新文件，保留最近14天的日志文件，多余的自动清理掉
	writer, _ := rotatelogs.New(
		achillesLog.OutputPath+".%Y%m%d",
		rotatelogs.WithLinkName(achillesLog.OutputPath),
		rotatelogs.WithMaxAge(time.Duration(achillesLog.MaxTime)*time.Hour*24),
		rotatelogs.WithRotationTime(time.Duration(achillesLog.RotateTime)*time.Hour*24),
	)

	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(writer)
	log.Info("initialized log config success")
}

func initHttpServer() {
	achillesApplication := config.GetApplication()
	server := service.NewServer(achillesApplication)
	http.ListenAndServe(server.Addr, server.Handler)
	log.Info("initialized http server success")
}

func main() {
	log.Infof("version %v running...", VERSION)
	initLog()
	initHttpServer()
}
