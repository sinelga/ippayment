package timeconverter

import (
	"time"
)

func StrToTime(timestr string) time.Time {

	layout := "2006/01/02 15:04:05"
	rettime, _ := time.Parse(layout, timestr)

	return rettime

}
