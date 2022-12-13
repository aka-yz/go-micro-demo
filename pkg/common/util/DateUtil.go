package util

import "time"

// CurrentSecond 获取当前时间戳（秒）
func CurrentSecond() (current int) {
	return int(time.Now().Unix())
}
