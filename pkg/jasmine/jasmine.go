package jasmine

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type MessageBuilder interface {
	WithDesc(desc string) MessageBuilder
	WithDescf(format string, args ...interface{}) MessageBuilder
	WithError(err error) MessageBuilder
	WithContext(ctx context.Context, keys ...string) MessageBuilder
	WithField(key, value string) MessageBuilder
	WithTimeStamp() MessageBuilder
	WithLargeInfo(key, info string) MessageBuilder
	WithJSON(key string, data interface{}) MessageBuilder
	WithJSONIndent(key string, data interface{}) MessageBuilder
	Build() string
}

// LogMessageBuilder 是日志消息的构建器
type LogMessageBuilder struct {
	fields    []string          // 动态字段列表，用于存储要输出的日志字段
	ctx       []string          // 保存传入的 context 字段
	err       error             // 错误信息
	timestamp int64             // 时间戳
	largeInfo map[string]string // 大信息
}

type MessageOption func(*LogMessageBuilder)

// NewLogMessageBuilder 初始化一个新的日志构建器，直接持有传入的 context.Context
func NewLogMessageBuilder() *LogMessageBuilder {
	b := &LogMessageBuilder{}
	b.largeInfo = make(map[string]string)
	return b
}

func (b *LogMessageBuilder) WithDesc(desc string) MessageBuilder {
	b.fields = append(b.fields, desc)
	return b
}

// WithDescf 使用格式化的描述信息
func (b *LogMessageBuilder) WithDescf(format string, args ...interface{}) MessageBuilder {
	b.fields = append(b.fields, fmt.Sprintf(format, args...))
	return b
}

func (b *LogMessageBuilder) WithError(err error) MessageBuilder {
	b.err = err
	return b
}

func (b *LogMessageBuilder) WithContext(ctx context.Context, keys ...string) MessageBuilder {
	for _, key := range keys {
		value := ctx.Value(key)
		if value != nil {
			b.ctx = append(b.ctx, fmt.Sprintf("%s=%v", key, value))
		}
	}
	return b
}

func (b *LogMessageBuilder) WithField(key, value string) MessageBuilder {
	b.fields = append(b.fields, fmt.Sprintf("%s=%v", key, value))
	return b
}

func (b *LogMessageBuilder) WithTimeStamp() MessageBuilder {
	b.timestamp = time.Now().Unix()
	return b
}

// WithLargeInfo 添加大信息，另起一行显示
func (b *LogMessageBuilder) WithLargeInfo(key, info string) MessageBuilder {

	b.largeInfo[key] = info
	return b
}

func (b *LogMessageBuilder) WithJSON(key string, data interface{}) MessageBuilder {
	jsonData, err := json.Marshal(data)
	if err != nil {
		jsonData = []byte(fmt.Sprintf("error marshaling JSON: %v", err))
	}
	b.largeInfo[key] = string(jsonData)
	return b
}
func (b *LogMessageBuilder) WithJSONIndent(key string, data interface{}) MessageBuilder {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		jsonData = []byte(fmt.Sprintf("error marshaling JSON: %v", err))
	}
	b.largeInfo[key] = string(jsonData)
	return b
}

// Build 构建最终的日志字符串
func (b *LogMessageBuilder) Build() string {
	var msg strings.Builder

	// 拼接所有动态添加的字段
	if len(b.fields) > 0 {
		msg.WriteString(strings.Join(b.fields, " | "))
	}

	// 单独添加 Context 信息
	if len(b.ctx) > 0 {
		msg.WriteString("\nContext: ")
		msg.WriteString(strings.Join(b.ctx, " | "))
	}

	// 添加时间戳

	if b.timestamp > 0 {
		msg.WriteString("\nTimestamp: ")
		msg.WriteString(fmt.Sprintf("%d", b.timestamp))
	}

	// 输出大信息
	if len(b.largeInfo) > 0 {
		for key, info := range b.largeInfo {
			msg.WriteString(fmt.Sprintf("\n%s:", key))
			msg.WriteString(fmt.Sprintf(" %s", info))
		}
	}

	// 单独添加错误信息
	if b.err != nil {
		msg.WriteString("\nError: ")
		msg.WriteString(b.err.Error())
	}

	return msg.String()
}
