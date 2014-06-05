package htmlhandler

import (
	"code.google.com/p/go.net/html"
	"compress/gzip"
	"domains"
	//	"fmt"
	"log/syslog"
	"os"
)

func ParseHtmlGzFile(golog syslog.Writer, htmlfile string) domains.MobClientHtml {

	var clphonenum string
	var block string
	var mobClientHtml domains.MobClientHtml

	var doc *html.Node

	sf, err := os.Open(htmlfile)
	if err != nil {
		golog.Err("ParseHtmlGzFile: " + htmlfile + err.Error())
	}
	defer sf.Close()
	s, err := gzip.NewReader(sf)
	if err != nil {
		golog.Err("ParseHtmlGzFile: " + htmlfile + err.Error())
	}
	defer s.Close()

	doc, err = html.Parse(s)
	if err != nil {
		golog.Err("ParseHtmlGzFile: " + htmlfile + err.Error())
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "div" && n.Attr[0].Val == "tel" {
				if n.FirstChild != nil {
					clphonenum = n.FirstChild.Data
					mobClientHtml.ClPhonenum = clphonenum
				}
			}
			if n.Data == "div" && n.Attr[0].Val == "block" {
				if n.FirstChild != nil {
					block = n.FirstChild.Data
					mobClientHtml.ClBlock = block
				}
			}
			if n.Data == "table" && n.Attr[0].Val == "payments" {

				paymentsarr := ParseTable(*n)

				for _, paymentstrarr := range paymentsarr {

					if len(paymentstrarr) == 3 {
						payment := domains.Payments{
							Created: paymentstrarr[0],
							Amount:  paymentstrarr[1],
							Result:  paymentstrarr[2],
						}
						mobClientHtml.ClPayments = append(mobClientHtml.ClPayments, payment)

					} else {
						golog.Err("ParseHtmlGzFile: paymentstrarr len!!! " + htmlfile)

					}

				}

			}
			if n.Data == "table" && n.Attr[0].Val == "hits" {
				hitsarr := ParseTable(*n)

				for _, hitstrarr := range hitsarr {

					if len(hitstrarr) == 5 {
						hithtml := domains.HitHtml{

							Createdstr: hitstrarr[0],
							Site:       hitstrarr[1],
							Themes:     hitstrarr[2],
							Resource:   hitstrarr[3],
							Provider:   hitstrarr[4],
						}
						mobClientHtml.ClHits = append(mobClientHtml.ClHits, hithtml)
					} else {
						golog.Err("ParseHtmlGzFile: hitstrarr len!!! " + htmlfile)
					}

				}

			}

			if n.Data == "table" && n.Attr[0].Val == "smsout" {
				smsoutarr := ParseTable(*n)
				for _, smsoutstrarr := range smsoutarr {

					if len(smsoutstrarr) == 3 {
						smsouthtml := domains.SmsOutHtml{

							Created: smsoutstrarr[0],
							From:    smsoutstrarr[1],
							Text:    smsoutstrarr[2],
						}
						mobClientHtml.ClSmsOut = append(mobClientHtml.ClSmsOut, smsouthtml)
					} else {

						golog.Err("ParseHtmlGzFile:  len!!! smsoutstrarr" + htmlfile)
					}

				}

			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return mobClientHtml

}
func ParseTable(n html.Node) [][]string {

	var retarr [][]string

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		if c.Data == "tbody" {

			for tb := c.FirstChild; tb != nil; tb = tb.NextSibling {

				var trvar []string
				if tb.Data == "tr" {

					for td := tb.FirstChild; td != nil; td = td.NextSibling {

						if td.Data == "td" {

							trvar = append(trvar, td.FirstChild.Data)

						}
					}

				}

				if len(trvar) > 0 {

					retarr = append(retarr, trvar)
				}

			}

		}

	}

	return retarr

}
