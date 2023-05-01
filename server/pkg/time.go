package pkg

import (
	"time"
)

const (
	DateFormat    = "2006-01-02 15:04:05"
	DayFormatFile = "2006_01_02"
	DayFormat     = "2006-01-02"
	DayIntFormat  = "20060102"
	DayFormatYmdH = "2006-01-02-15"
)

// TimeGetChineseTime 获取中国时间戳
func TimeGetChineseTime() time.Time {
	// 中国比美国快8个小时
	return time.Now().Add(time.Hour * 8)
}

func TimeGetChineseTimeStr() string {
	return TimeGetChineseTime().Format(DayFormat)
}
