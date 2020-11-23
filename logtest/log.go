package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
	"os"
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
		level)
	logger := zap.New(core)
	logger.Info("DefaultLogger init success")
	return logger
}

func main() {
	//testlog()

	//dynaic()
	mutiOutput()
}

// 动态修改日志级别
func dynaic() {
	logger := initLogger("./all.log", "")

	alevel := zap.NewAtomicLevel()
	http.HandleFunc("/handle/level", alevel.ServeHTTP)
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			panic(err)
		}
	}()
	// 默认是Info级别
	logcfg := zap.NewProductionConfig()
	logcfg.Level = alevel
	logger, err := logcfg.Build()
	if err != nil {
		fmt.Println("err", err)
	}
	defer logger.Sync()
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Second)
		logger.Debug("debug log", zap.Int("line", 66), zap.String("level", alevel.String()), zap.String("debug", "debug-debug"))
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

// 打印到控制台,文件,kafka
func mutiOutput() {
	// First, define our level-handling logic.
	// 仅打印Error级别以上的日志
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	// 打印所有级别的日志
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	hook := lumberjack.Logger{
		Filename:   "/tmp/abc.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 3,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}

	topicErrors := zapcore.AddSync(ioutil.Discard)
	fileWriter := zapcore.AddSync(&hook)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		// 打印在kafka topic中（伪造的case）
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		// 打印在控制台
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		// 打印在文件中
		zapcore.NewCore(consoleEncoder, fileWriter, highPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a info logger", zap.Int("test", 1))
	logger.Error("constructed a error logger", zap.Int("test", 2))
}
