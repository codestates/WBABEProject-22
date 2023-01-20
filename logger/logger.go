package logger

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"oos/config"
)

var lg *zap.Logger

func InitLogger(cfg *config.Config) (err error) {
	cf := cfg.Log

	now := time.Now()
	lPath := fmt.Sprintf("%s_%s.log", cf.Fpath, now.Format("2006-01-02"))

	writeSyncer := getLogWriter(lPath, cf.Msize, cf.Mbackup, cf.Mage)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cf.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return
}

func Debug(ctx ...interface{}) {
	var b bytes.Buffer
	for _, str := range ctx {
		b.WriteString(str.(string))
		b.WriteString(" ")
	}

	lg.Debug("debug", zap.String("-", b.String()))
}

func Info(ctx ...interface{}) {
	var b bytes.Buffer
	for _, str := range ctx {
		b.WriteString(str.(string))
		b.WriteString(" ")
	}

	lg.Info("info", zap.String("-", b.String()))
}

func Warn(ctx ...interface{}) {
	var b bytes.Buffer
	for _, str := range ctx {
		b.WriteString(str.(string))
		b.WriteString(" ")
	}

	lg.Warn("warn", zap.String("-", b.String()))
}

func Error(ctx ...interface{}) {
	var b bytes.Buffer
	for _, str := range ctx {
		b.WriteString(str.(string))
		b.WriteString(" ")
	}

	lg.Error("error", zap.String("-", b.String()))
}
/* [코드리뷰]
 * 로그 레벨을 잘 나누어 주셨습니다. 
 * 이것은 개발자의 스타일이여서 꼭 필요한 comment는 아니지만 
 * 기본적으로 해당 log package에서 제공되는 log level을 모두 구현해주어도 좋을 것 같습니다.
 * zap은 Debug, Info, Warning, Error, DPanic, Panic, Fatal 총 7가지 디버깅 레벨을 제공해줍니다.
 * 사용되지 않더라도, 해당 레벨들을 미리 구현해놓으면 이후에 필요한 상황에 새롭게 추가하는 번거로움이 적어질 것 같습니다.
 */

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		lg.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection,
				// as it is not really a condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// References
// Class material: lecture 12