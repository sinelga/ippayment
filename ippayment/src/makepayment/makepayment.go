package makepayment

import (
	"domains"
	"log/syslog"
	"net/http"
	"net/url"
	"time"
)

func Pay(golog syslog.Writer, clphonenum string, provider string, amount string, themes string) domains.Payments {
	var retpayment domains.Payments
	var urlstr string

	if Url, err := url.Parse("http://ippayment.info:8080"); err != nil {

		golog.Err(err.Error())
	} else {

		Url.Path += "/MobileCharger"

		parameters := url.Values{}
		parameters.Add("Action", "Check")
		parameters.Add("MSISDN", "+"+clphonenum)
		parameters.Add("ServiceId", "2")
		parameters.Add("Price", amount)

		Url.RawQuery = parameters.Encode()
		urlstr = Url.String()

	}

	resp, err := http.Get(urlstr)
	defer resp.Body.Close()
	if err != nil {

		golog.Crit("makepayment: " + err.Error())
	} else {

		status := resp.Header.Get("Status")
		golog.Info("Status1 " + status)
		transactionid := resp.Header.Get("TransactionId")
		golog.Info(transactionid)

		if status == "0" {

			if Url, err := url.Parse("http://ippayment.info:8080"); err != nil {

				golog.Err(err.Error())
			} else {

				Url.Path += "/MobileCharger"

				parameters := url.Values{}
				parameters.Add("Action", "Charge")
				parameters.Add("MSISDN", "+"+clphonenum)
				parameters.Add("ServiceId", "2")
				parameters.Add("Price", amount)
				parameters.Add("TransactionId", transactionid)

				Url.RawQuery = parameters.Encode()
				urlstr = Url.String()
				golog.Info(urlstr)

				resp, err := http.Get(urlstr)
				defer resp.Body.Close()
				if err != nil {

					golog.Crit("makepayment-2: " + err.Error())
				} else {
					status = resp.Header.Get("Status")
					golog.Info("StatusCharge " + status)

				}

			}
		}

		t := time.Now().Local()
		tstr := t.Format("2006/01/02 15:04:05")
		retpayment.Created = tstr
		retpayment.Amount = amount
		retpayment.Themes = themes
		retpayment.Result = status

	}

	return retpayment
}
