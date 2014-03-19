package checkindex

import (
	"domains"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
)

func CheckIndx(golog syslog.Writer, tdDB db.DB, hit domains.Hits) {

	msisdn := tdDB.Use(hit.Msisdn)

	var indexarr []string

	for index := range msisdn.SecIndexes {
		//		fmt.Printf("I have a collection called %s\n", name)
		indexarr = append(indexarr, index)
	}

	if len(indexarr) == 0 {

		if err := msisdn.Index([]string{"Created"}); err != nil {
			panic(err)
		}
		if err := msisdn.Index([]string{"ColCreated"}); err != nil {
			panic(err)
		}

		if hit.Resource == "mobilephone" {

			if err := msisdn.Index([]string{"SmsCreated"}); err != nil {
				panic(err)
			}

		}

	}
}
