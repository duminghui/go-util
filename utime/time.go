// Package util provides ...
package utime

import (
	"time"
)

func FormatLongTimeStrUTC(seconds int64) string {
	times := time.Unix(seconds, 0)
	localUTC, _ := time.LoadLocation("")
	timeUTC := times.In(localUTC)
	return timeUTC.Format("2006-01-02 15:04:05")
}

func FormatTimeStrUTC(t time.Time) string {
	localUTC, _ := time.LoadLocation("")
	timeUTC := t.In(localUTC)
	return timeUTC.Format("2006-01-02 15:04:05")
}
