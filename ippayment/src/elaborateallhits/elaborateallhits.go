package elaborateallhits

import (
//"log"
"log/syslog"
"github.com/garyburd/redigo/redis"

)

func ElabAllHits(golog syslog.Writer, c redis.Conn,collections []string) {

	if quan_hits, err := redis.Int(c.Do("LLEN", "hits")); err != nil {

		golog.Crit(err.Error())

	} else {
	
	if quan_hits > 1 {
	
		for i :=0; i< quan_hits; {
		
		println(i)
		
		
		
		}
	
	
	
	}
	
	}
	
	

}