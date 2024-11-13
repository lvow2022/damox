package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// QueryParam 在 gin.Context 中扩展一个泛型的 QueryParam 方法
func QueryParam[T any](c *gin.Context, key string, defaultValue T) (T, error) {
	var result T

	// 获取查询参数值
	val := c.DefaultQuery(key, fmt.Sprintf("%v", defaultValue))

	// 使用类型断言直接对 result 进行类型转换
	switch v := any(&result).(type) { // 转换成 interface{}
	case *int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return defaultValue, err
		}
		*v = i
	case *string:
		*v = val
	case *bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return defaultValue, err
		}
		*v = b
	case *float64:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return defaultValue, err
		}
		*v = f
	case *int64:
		i64, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return defaultValue, err
		}
		*v = i64
	case *time.Duration:
		d, err := time.ParseDuration(val)
		if err != nil {
			return defaultValue, err
		}
		*v = d
	case *time.Time:
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return defaultValue, err
		}
		*v = t
	}

	return result, nil
}
