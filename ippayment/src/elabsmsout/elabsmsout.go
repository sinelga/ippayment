package elabsmsout

import (
	"domains"
	"time"
	"log/syslog"
	"strconv"
)

func Elab(golog syslog.Writer,smsoutarr []domains.SmsOut) {
	
	nowunix := time.Now().Unix()
		
	var lastdate int64
	for _, smsout := range smsoutarr {

		if lastdate < smsout.SmsCreated {
			lastdate = smsout.SmsCreated

		}
	}
	difftime := (nowunix - lastdate)/60
	
	golog.Info("Time sends SMS min? "+strconv.FormatInt(difftime,10))
		

}
