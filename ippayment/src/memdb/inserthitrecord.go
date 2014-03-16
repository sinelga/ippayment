package memdb

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	//"strconv"
	"time"
)

func InsertHit(golog syslog.Writer, c redis.Conn, record []string) {

	nowUnix :=time.Now().Unix()
	nowUnixInt := int(nowUnix)

	hit := domains.Hits{
		Created:  nowUnixInt,
		Id:       record[0],
		Msisdn:   record[1],
		Site:     record[2],
		Themes:   record[3],
		Resource: record[4],
	}

	bithit, _ := json.Marshal(hit)

	if _, err := c.Do("LPUSH", "hits", bithit); err != nil {

		golog.Crit(err.Error())

	}

}
