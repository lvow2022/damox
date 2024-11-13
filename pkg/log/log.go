package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"refactory/pkg/jasmine"
	"time"
)

// Logger 包装 zap.Logger 和 logmessage
type Logger struct {
	l *zap.Logger
}

// NewLogger 创建一个新的日志实例，注入 zap.Logger 实例
func NewLogger(zapLogger *zap.Logger) *Logger {
	return &Logger{
		l: zapLogger,
	}
}

// Info 打印 Info 级别的日志
func (l *Logger) Info(msg string) {
	l.l.Info(msg)
}

// Error 打印 Error 级别的日志
func (l *Logger) Error(msg string) {
	l.l.Error(msg)
}

func NewMsg() jasmine.MessageBuilder {
	return jasmine.NewLogMessageBuilder()
}

func main() {
	// 自定义 EncoderConfig，添加文件名和行号
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "timestamp",  // 时间戳字段名
		LevelKey:      "level",      // 日志级别字段名
		MessageKey:    "message",    // 消息字段名
		CallerKey:     "caller",     // 调用者字段名
		StacktraceKey: "stacktrace", // 堆栈字段名
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			// 自定义时间格式：2024-11-07 15:35:43
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 日志级别格式：大写
		EncodeCaller: zapcore.ShortCallerEncoder,  // 使用简短格式的调用者（文件名:行号）
	}

	// 控制台输出编码器
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 设置日志级别
	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	// 创建日志核心
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)

	// 创建 logger
	logger := zap.New(core, zap.AddCaller()) // 启用 Caller 信息

	// 使用 context 传递一些键值对
	ctx := context.Background()
	ctx = context.WithValue(ctx, "retry", 2)
	ctx = context.WithValue(ctx, "sessionID", "sess-789")

	l := NewLogger(logger)

	// Example large JSON data
	largeData := map[string]interface{}{
		"user_id":   123,
		"status":    "failed",
		"attempts":  3,
		"error":     "database connection error",
		"timestamp": time.Now().String(),
		"details": map[string]interface{}{
			"retry_count": 5,
			"source":      "es",
			"duration":    "200ms",
		},
	}
	type UserInfo struct {
		UserID   int    `json:"user_id"`
		Status   string `json:"status"`
		Attempts int    `json:"attempts"`
	}
	// 示例结构体数据
	userData := UserInfo{
		UserID:   123,
		Status:   "failed",
		Attempts: 3,
	}

	// Convert the JSON data to a formatted JSON string
	l.Info(NewMsg().WithContext(ctx, "retry", "sessionID").
		WithDescf("[uid %s] user login fail", "123").
		WithField("retry", "5").WithField("source", "es").WithField("spend", "200ms").
		WithError(errors.New("sql connect err")).
		WithJSON("DetailedLog", largeData).
		WithJSON("UserDetails", userData).
		WithTimeStamp().Build())

	// 打印日志信息
	// 记录日志

	defer logger.Sync() // 确保日志被写入

}
