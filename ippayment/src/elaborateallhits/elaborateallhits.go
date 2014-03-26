package elaborateallhits

import (
	"checkcolexist"
	"domains"
	"elabsmsout"
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/garyburd/redigo/redis"
	"github.com/mitchellh/mapstructure"
	"log/syslog"
	"pushsmsout"
	"time"
)

func ElabAllHits(golog syslog.Writer, c redis.Conn, tdDB db.DB) {

	if quan_hits, err := redis.Int(c.Do("LLEN", "hits")); err != nil {

		golog.Crit(err.Error())

	} else {

		nowunix := time.Now().Unix()

		col := tdDB.Use("mobclients")

		if quan_hits > 0 {

			for i := 0; i < quan_hits; i++ {

				var hit domains.Hit

				bhit, _ := redis.Bytes(c.Do("RPOP", "hits"))

				err := json.Unmarshal(bhit, &hit)
				if err != nil {

					golog.Crit(err.Error())

				}

				docID := checkcolexist.ChecExist(col, hit.Msisdn)
				if docID == 0 {

					var hitsarr []domains.Hit
					var smsoutarr []domains.SmsOut

					hitsarr = append(hitsarr, hit)

					if hit.Resource == "mobilephone" {

						smsout := domains.SmsOut{
							SmsCreated: nowunix,
							Msisdn:     hit.Msisdn,
							From:       "070095943",
							Text:       "Soita nyt! Miia. " + hit.Site,
							Provider:   hit.Provider,
						}
						pushsmsout.PushOut(golog, c, smsout)
						smsoutarr = append(smsoutarr, smsout)

						col.Insert(map[string]interface{}{
							"ClPhonenum": hit.Msisdn,
							"ClCreated":  hit.Created,
							"ClUpdated":  hit.Created,
							"ClThemes":   hit.Themes,
							"ClResource": hit.Resource,
							"ClProvider": hit.Provider,
							"ClBlock":    0,
							"ClHits":     hitsarr,
							"ClSmsOut":   smsoutarr,
						})

					} else {

						col.Insert(map[string]interface{}{
							"ClPhonenum": hit.Msisdn,
							"ClCreated":  hit.Created,
							"ClUpdated":  hit.Created,
							"ClThemes":   hit.Themes,
							"ClResource": hit.Resource,
							"ClProvider": hit.Provider,
							"ClBlock":    0,
							"ClHits":     hitsarr,
						})

					}

				} else {

					var readBack interface{}

					col.Read(docID, &readBack)

					mobclientval := readBack.(map[string]interface{})
					var mobclientobj domains.MobClient

					err := mapstructure.Decode(mobclientval, &mobclientobj)
					if err != nil {
						panic(err)
					}

					mobclientobj.ClHits = append(mobclientobj.ClHits, hit)

					if hit.Resource == "mobilephone" {

						smsoutarrupdated := elabsmsout.Elab(golog, c, hit.Msisdn, hit.Provider, hit.Site, mobclientobj.ClSmsOut)

						col.Update(docID, map[string]interface{}{
							"ClPhonenum": mobclientobj.ClPhonenum,
							"ClCreated":  mobclientobj.ClCreated,
							"ClUpdated":  nowunix,
							"ClThemes":   mobclientobj.ClThemes,
							"ClResource": mobclientobj.ClResource,
							"ClProvider": mobclientobj.ClProvider,
							"ClBlock":    mobclientobj.ClBlock,
							"ClHits":     mobclientobj.ClHits,
							"ClSmsOut":   smsoutarrupdated,
						})

					} else {

						col.Update(docID, map[string]interface{}{
							"ClPhonenum": mobclientobj.ClPhonenum,
							"ClCreated":  mobclientobj.ClCreated,
							"ClUpdated":  nowunix,
							"ClThemes":   mobclientobj.ClThemes,
							"ClResource": mobclientobj.ClResource,
							"ClProvider": mobclientobj.ClProvider,
							"ClBlock":    mobclientobj.ClBlock,
							"ClHits":     mobclientobj.ClHits,
						})

					}

				}

			}

		}

	}

}
