package middleware

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// These Logger is for access log.
// Error log handling is in `main/httpErrorHandler()`

func logFormat() string {
	// Refer to https://github.com/tkuchiki/alp
	var format string
	format += "time:${time_rfc3339}\t"
	format += "host:${remote_ip}\t"
	format += "forwardedfor:${header:x-forwarded-for}\t"
	format += "req:-\t"
	format += "status:${status}\t"
	format += "method:${method}\t"
	format += "uri:${uri}\t"
	format += "size:${bytes_out}\t"
	format += "referrer:${referrer}\t"
	format += "ua:${user_agent}\t"
	format += "reqtime_ns:${latency}\t"
	format += "cache:-\t"
	format += "runtime:-\t"
	format += "apptime:-\t"
	format += "vhost:${host}\t"
	format += "reqtime_human:${latency_human}\t"
	format += "x-request-id:${id}\t"
	format += "host:${host}\n"
	return format
}

// Logger : Logging middleware
func Logger() echo.MiddlewareFunc {

	ts := time.Now().Format("2006-01-02--15-04-05")
	fn := "./.logs/" + ts + ".access.log"

	if err := os.MkdirAll(".logs", os.ModePerm); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// [TODO] Switch output
	logger := middleware.LoggerWithConfig((middleware.LoggerConfig{
		Format: logFormat(),
		Output: logFile,
	}))

	return logger
}
