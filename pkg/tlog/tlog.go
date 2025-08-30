package tlog

import (
	"fmt"
	"log"
	"os"
)

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

var currentLevel = LevelDebug // lets prioritize debug messages by default

func SetLevel(level LogLevel) {
	currentLevel = level
}

func logMessage(level LogLevel, color string, levelStr string, format string, v ...any) {
	if level < currentLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	output := fmt.Sprintf("%s [%-5s] %s%s", color, levelStr, ColorReset, msg)

	if level == LevelFatal {
		log.Fatalln(output)
		os.Exit(1)
	} else {
		log.Println(output)
	}
}

func Debug(format string, v ...any) {
	logMessage(LevelDebug, ColorCyan, "DEBUG", format, v...)
}

func Info(format string, v ...any) {
	logMessage(LevelInfo, ColorGreen, "INFO", format, v...)
}

func Warn(format string, v ...any) {
	logMessage(LevelWarning, ColorYellow, "WARN", format, v...)
}

func Error(format string, v ...any) {
	logMessage(LevelError, ColorRed, "ERROR", format, v...)
}

func Fatal(format string, v ...any) {
	logMessage(LevelFatal, ColorRed, "FATAL", format, v...)
}

// ------- Gin integration -------

// func init() {
// 	// Redirect Gin’s default logs to tlog
// 	gin.DefaultWriter = GetWriter(LevelInfo)
// 	gin.DefaultErrorWriter = GetWriter(LevelError)
// }

// func TlogMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		start := time.Now()

// 		// Process request
// 		c.Next()

// 		// After request
// 		latency := time.Since(start)
// 		status := c.Writer.Status()
// 		clientIP := c.ClientIP()
// 		method := c.Request.Method
// 		path := c.Request.URL.Path

// 		if len(c.Errors) > 0 {
// 			// Gin collected errors (c.Error)
// 			for _, e := range c.Errors {
// 				Error("%s %s -> %d (%s) | %s", method, path, status, latency, e.Err)
// 			}
// 		} else {
// 			Info("%s %s -> %d (%s) | %s", method, path, status, latency, clientIP)
// 		}
// 	}
// }

// type logWriter struct {
// 	level LogLevel
// }

// func (lw *logWriter) Write(p []byte) (n int, err error) {
// 	// Trim newlines (Gin includes them)
// 	msg := string(p)
// 	switch lw.level {
// 	case LevelError:
// 		Error("%s", msg)
// 	default:
// 		Info("%s", msg)
// 	}
// 	return len(p), nil
// }

// func GetWriter(level LogLevel) io.Writer {
// 	return &logWriter{level: level}
// }

// ------- End of Gin integration -------
