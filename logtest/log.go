package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

func initLogger(logpath string, loglevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logpath, //日志文件路径
		MaxSize:    1024,    //megabytes
		MaxBackups: 3,       //最对保留3个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩,默认false
	}
	w := zapcore.AddSync(&hook)
	var level zapcore.Level
	//switch loglevel {
	//case "debug":
	//	level = zap.DebugLevel
	//case "info":
	//	level = zap.InfoLevel
	//case "error":
	//	level = zap.ErrorLevel
	//default:
	//	level = zap.InfoLevel
	//}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		w,
		level, )
	logger := zap.New(core)
	logger.Info("DefaultLogger init success")
	return logger
}

func main() {
	//testlog()

	logger := initLogger("./all.log", "")

	alevel := zap.NewAtomicLevel()
	http.HandleFunc("/handle/level", alevel.ServeHTTP)
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			panic(err)
		}
	}()
	// 默认是Info级别
	//logcfg := zap.NewProductionConfig()
	//logcfg.Level = alevel
	//logger, err := logcfg.Build()
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	defer logger.Sync()
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Second)
		logger.Debug("debug log", zap.Int("line", 66), zap.String("level", alevel.String()))
		logger.Info("Info log", zap.Int("line", 67), zap.String("level", alevel.String()))
	}
}

func panic1() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func testlog() {
	url := "Hello"
	logger, _ := zap.NewProduction()
	defer panic1()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Warn("debug log", zap.String("level", url))
	logger.Error("Error Message", zap.String("error", url))
	logger.Panic("Panic log", zap.String("panic", url))
}
