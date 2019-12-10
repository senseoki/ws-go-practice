package main

import (
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	app := iris.New()

	logger := zapConfig()

	app.Get("/ping", func(ctx iris.Context) {
		for i := 0; i < 1000000; i++ {
			logger.Error("This is an ERROR message")
		}
		ctx.WriteString("pong")
	})

	app.Get("/test01", func(ctx iris.Context) {
		ctx.WriteString("test01")
	})

	logger.Info("Logger Test Info",
		zap.String("url", "http://test.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger.Error("This is an ERROR message")

	//logger.DPanic("This is a DPANIC message")

	app.Run(
		iris.Addr("0.0.0.0:9100"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithoutInterruptHandler,
	)

}

func zapConfig() *zap.Logger {
	//now := time.Now()
	hook := lumberjack.Logger{
		Filename: "logs/iris-test.log",
		MaxSize:  1, // megabytes
		//MaxBackups: 30,
		MaxAge:   30, // days
		Compress: false,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)

	caller := zap.AddCaller()
	development := zap.Development()
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	logger := zap.New(core, caller, development, filed)

	return logger
}

// func zapConfig() *zap.Logger {

// 	// w := zapcore.AddSync(&lumberjack.Logger{
// 	// 	Filename: "logdata/foo.log",
// 	// 	MaxSize:  1, // megabytes
// 	// 	//MaxBackups: 3,
// 	// 	MaxAge: 30, // days
// 	// })
// 	// core := zapcore.NewCore(
// 	// 	zapcore.NewJSONEncoder(
// 	// 		// zap.NewProductionEncoderConfig(),
// 	// 		zap.NewDevelopmentEncoderConfig(),
// 	// 	),
// 	// 	w,
// 	// 	zap.InfoLevel,
// 	// )

// 	// logger := zap.New(core)

// 	now := time.Now()
// 	hook := &lumberjack.Logger{
// 		Filename: fmt.Sprintf("log/%04d%02d%02d%02d%02d%02d.log", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), //filePath
// 		MaxSize:  1,                                                                                                                           // megabytes
// 		//MaxBackups: 10000,
// 		MaxAge:   30,    //days
// 		Compress: false, // disabled by default
// 	}
// 	defer hook.Close()

// 	enConfig := zap.NewProductionEncoderConfig()
// 	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
// 	level := zap.InfoLevel
// 	w := zapcore.AddSync(hook)
// 	core := zapcore.NewCore(
// 		zapcore.NewConsoleEncoder(enConfig),
// 		w,
// 		level,
// 	)

// 	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
// 	return logger
// }
