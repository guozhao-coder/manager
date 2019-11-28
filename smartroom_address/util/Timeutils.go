package util
import (
	"time"
)

const (
	FIVE_MINUTE_STEP = 5 * 60       //5分钟
	ONE_HOUR_STEP    = 60 * 60      // 一小时
	ONE_DAY_STEP     = 24 * 60 * 60 //一天
)

func GetUnix13NowTime() int64 {
	return time.Now().UTC().UnixNano() / 1000000
}