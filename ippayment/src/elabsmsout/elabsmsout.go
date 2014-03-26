package elabsmsout

import (
	"domains"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"pushsmsout"
	"strconv"
	"time"
)

func Elab(golog syslog.Writer, c redis.Conn, clphonenum string, provider string, site string, smsoutarr []domains.SmsOut) []domains.SmsOut {

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

		golog.Info("Time sends Second SMS must be more 300 " + clphonenum + " " + strconv.FormatInt(lastdate, 10) + " diff " + strconv.FormatInt(difftime, 10))

		if difftime > 300 {

			smsout := domains.SmsOut{
				SmsCreated: nowunix,
				Msisdn:     clphonenum,
				From:       "070095943",
				Text:       "On aika tutustua! Miia soita. "+site,
				Provider:   provider,
				
			}
			pushsmsout.PushOut(golog, c, smsout)
			smsoutarr = append(smsoutarr, smsout)

		}

	}

	return smsoutarr

}
