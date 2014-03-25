package elabsmsout

import (
	"domains"
	"time"
	"log/syslog"
	"strconv"
)

func Elab(golog syslog.Writer,clphonenum string,smsoutarr []domains.SmsOut) {
	
	nowunix := time.Now().Unix()
		
	var lastdate int64
	for _, smsout := range smsoutarr {

		if lastdate < smsout.SmsCreated {
			lastdate = smsout.SmsCreated

		}
	}
	difftime := (nowunix - lastdate)
	
	golog.Info("Time sends SMS min? last "+clphonenum +" "+ strconv.FormatInt(lastdate,10)+" diff "+ strconv.FormatInt(difftime,10))
		

}
