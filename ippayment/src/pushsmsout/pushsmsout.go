package pushsmsout

import (
	"domains"
	"log/syslog"
	"github.com/garyburd/redigo/redis"
	"encoding/json"

)

func PushOut(golog syslog.Writer,c redis.Conn,smsout domains.SmsOut) {

	bitsmsout, _ := json.Marshal(smsout)

	if _, err := c.Do("LPUSH", "smsout", bitsmsout); err != nil {

		golog.Crit(err.Error())

	}


}

