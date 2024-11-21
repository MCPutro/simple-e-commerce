package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/MCPutro/E-commerce/pkg/constant"
	"github.com/sirupsen/logrus"
)

var (
	logger     *logrus.Logger
	loggerInit sync.Once
)

func NewLogger(level logrus.Level) *logrus.Logger {
	loggerInit.Do(func() {
		//check directory is existing
		logFilePath := "logs"
		if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
			// Directory does not exist, create it
			err := os.MkdirAll(logFilePath, 0755) // 0755 sets permissions (read/write for owner, read-only for others)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
			fmt.Println("Directory created:", logFilePath)
		} else {
			fmt.Println("Directory already exists:", logFilePath)
		}

		// set file log
		fileName := "logData.log"
		logFile, err := os.OpenFile(path.Join(logFilePath, fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}

		logger = logrus.New()
		//logger.SetFormatter(&logrus.TextFormatter{
		//	FullTimestamp:   true,
		//	TimestampFormat: customTimeFormat(),
		//})
		logger.SetOutput(logFile)
		logger.SetLevel(level)
	})

	return logger
}

func GetLogger() *logrus.Logger {
	return logger
}

func ContextLogger(ctx context.Context) *logrus.Entry {
	mlogger := GetLogger()

	var fields logrus.Fields

	if ctxRqId, ok := ctx.Value(constant.HeaderXRequestID).(string); ok {
		fields = logrus.Fields{
			"requestId": ctxRqId,
		}
	}

	return mlogger.WithFields(fields)
}
