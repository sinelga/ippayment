package hitshandler

import (
	"checkgzfileexist"
	"domains"
	"github.com/garyburd/redigo/redis"
	"htmlhandler"
	"log/syslog"
	"makepayment"
	"pushsmsout"
	"strconv"
	"time"
	"timeconverter"
)

func ElaborateHit(golog syslog.Writer, c redis.Conn, hit domains.Hit) domains.MobClientHtml {

	nowunix := time.Now().Unix()
	//	c, err := redis.Dial("tcp", ":6379")
	//	if err != nil {
	//		golog.Crit(err.Error())
	//	}

	htmlfile := string("/home/juno/git/ippayment/ippayment/www/tel/" + hit.Msisdn + ".html.gz")
	var mobclienthtml domains.MobClientHtml
	var smsouthtmlarr []domains.SmsOutHtml
	var paymentsarr []domains.Payments
	var hithtmlarr []domains.HitHtml
	var cliblock string
	var blockbool bool
	tvisit := time.Unix(hit.Created, 0)
	tvisitstr := tvisit.Format("2006/01/02 15:04:05")
	//	t := time.Now().Local()
	t := time.Now()
	tstr := t.Format("2006/01/02 15:04:05")

	hithmtl := domains.HitHtml{Createdstr: tvisitstr, Site: hit.Site, Themes: hit.Themes, Resource: hit.Resource, Provider: hit.Provider}

	if checkgzfileexist.CheckFile(golog, htmlfile) {

		golog.Info(htmlfile + " Exist")

		mobclienthtml = htmlhandler.ParseHtmlGzFile(golog, htmlfile)

		if mobclienthtml.ClBlock == "Yes" {

			blockbool = true

		} else {

			blockbool = false
		}

		mobclienthtml.ClHits = append(mobclienthtml.ClHits, hithmtl)

		if hit.Resource == "mobilephone" {

			lastsmsdatestr := mobclienthtml.ClSmsOut[len(mobclienthtml.ClSmsOut)-1].Created

			lastsmsdate := timeconverter.StrToTime(lastsmsdatestr)

			diffsmsouttime := t.Sub(lastsmsdate)

			golog.Info(strconv.FormatFloat(diffsmsouttime.Minutes(), 'f', 6, 64))

			if diffsmsouttime.Minutes() > -170 && !blockbool && (len(mobclienthtml.ClSmsOut) == 1) {

				smsout := domains.SmsOut{
					SmsCreated: nowunix,
					Msisdn:     hit.Msisdn,
					From:       "070095943",
					Text:       "Nyt on aika tutustua. Miia, soita! http://" + hit.Site,
					Provider:   hit.Provider,
				}

				pushsmsout.PushOut(golog, c, smsout)

				smsouthtml := domains.SmsOutHtml{
					Created: tstr,
					From:    "070095943",
					Text:    "Nyt on aika tutustua. Miia, soita! http://" + hit.Site,
				}
				mobclienthtml.ClSmsOut = append(mobclienthtml.ClSmsOut, smsouthtml)

			}

			if diffsmsouttime.Minutes() > 120 && !blockbool && (len(mobclienthtml.ClSmsOut) == 2) {

				smsout := domains.SmsOut{
					SmsCreated: nowunix,
					Msisdn:     hit.Msisdn,
					From:       "070095943",
					Text:       "En ole nähnyt sinua pitkään. Minulle on tullut ikävä. Soita! Miia http://" + hit.Site,
					Provider:   hit.Provider,
				}

				pushsmsout.PushOut(golog, c, smsout)

				smsouthtml := domains.SmsOutHtml{
					Created: tstr,
					From:    "070095943",
					Text:    "En ole nähnyt sinua pitkään. Minulle on tullut ikävä. Soita! Miia http://" + hit.Site,
				}
				mobclienthtml.ClSmsOut = append(mobclienthtml.ClSmsOut, smsouthtml)

			}

		}

		lastpaymentdatestr := mobclienthtml.ClPayments[len(mobclienthtml.ClPayments)-1].Created
		lastpaymentdate := timeconverter.StrToTime(lastpaymentdatestr)

		diffpaymenttime := t.Sub(lastpaymentdate)
		if diffpaymenttime.Minutes() > 43200 && !blockbool {

			mobclienthtml.ClPayments = append(mobclienthtml.ClPayments, makepayment.Pay(golog, hit.Msisdn, hit.Provider, "70", hit.Themes))

		}

	} else {

		golog.Info(htmlfile + " Don't Exist")

		if hit.Resource == "mobilephone" {

			smsout := domains.SmsOut{
				SmsCreated: nowunix,
				Msisdn:     hit.Msisdn,
				From:       "070095943",
				Text:       "Soita nyt! Miia. http://" + hit.Site,
				Provider:   hit.Provider,
			}

			pushsmsout.PushOut(golog, c, smsout)

			smsouthtml := domains.SmsOutHtml{
				Created: tstr,
				From:    "070095943",
				Text:    "Soita nyt! Miia. http://" + hit.Site,
			}

			smsouthtmlarr = append(smsouthtmlarr, smsouthtml)

		}

		paymentsarr = append(paymentsarr, makepayment.Pay(golog, hit.Msisdn, hit.Provider, "70", hit.Themes))

		for _, payment := range paymentsarr {

			golog.Info(payment.Amount + " result " + payment.Result)

		}

		hithtmlarr = append(hithtmlarr, hithmtl)

		cliblock = "No"

		mobclienthtml.ClBlock = cliblock
		mobclienthtml.ClPhonenum = hit.Msisdn
		mobclienthtml.ClHits = hithtmlarr
		mobclienthtml.ClPayments = paymentsarr
		mobclienthtml.ClSmsOut = smsouthtmlarr

	}

	return mobclienthtml

}
