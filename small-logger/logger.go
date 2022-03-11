package small_logger

import (
	"context"
	"fmt"
	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

var SmallLogger = smallLogger{logLevel: logger2.Info}

var infoStr = "%s\n[info] "
var warnStr = "%s\n[warn] "
var errStr = "%s\n[error] "
var traceStr = "%s\n[%.3fms] [rows:%v] %s"

type smallLogger struct {
	logLevel logger2.LogLevel
}

func (l smallLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger2.Info {
		fmt.Printf(infoStr+msg+"\n", append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l smallLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger2.Warn {
		fmt.Printf(warnStr+msg+"\n", append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l smallLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger2.Error {
		fmt.Printf(errStr+msg+"\n", append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l smallLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if l.logLevel <= logger2.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	if rows == -1 {
		fmt.Printf(traceStr+"\n", utils.FileWithLineNum(), float64(elapsed.Milliseconds())/1e3, "-", sql)
	} else {
		fmt.Printf(traceStr+"\n", utils.FileWithLineNum(), float64(elapsed.Milliseconds())/1e3, rows, sql)
	}
}

// LogMode log mode
func (l *smallLogger) LogMode(level logger2.LogLevel) logger2.Interface {
	newlogger := *l
	newlogger.logLevel = level
	return &newlogger
}
