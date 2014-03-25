package elabsmsout

import (
	"domains"
	"log/syslog"
	"strconv"
	"time"
)

func Elab(golog syslog.Writer, clphonenum string, smsoutarr []domains.SmsOut) {

	nowunix := time.Now().Unix()

	var lastdate int64

	if len(smsoutarr) == 0 {

		golog.Info("!!! not smsout for??? " + clphonenum)

	} else {

		for _, smsout := range smsoutarr {

			if lastdate < smsout.SmsCreated {
				lastdate = smsout.SmsCreated

			}
		}
		difftime := (nowunix - lastdate)

		golog.Info("Time sends SMS min? last " + clphonenum + " " + strconv.FormatInt(lastdate, 10) + " diff " + strconv.FormatInt(difftime, 10))

	}

}
