// Package util provides ...
package utime

import (
	"time"
)

func FormatLongTimeStrUTC(seconds int) string {
	times := time.Unix(int64(seconds), 0)
	localUTC, _ := time.LoadLocation("")
	timeUTC := times.In(localUTC)
	return timeUTC.Format("2006-01-02 15:04:05")
}
