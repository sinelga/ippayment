package hitshandler

import (
	"CreateHtmlGz"
	"domains"
	"log"
	"log/syslog"
	"testing"
)

var hitsTests = []domains.Hit{
	{Created: 1400800260, Id: "cd71e4c0-db30-11e3-c47b-71a9d04148af", Msisdn: "358400177858", Site: "tissit.tv", Themes: "adult", Resource: "mobilephone", Provider: "MobileSonera"},
	{Created: 1400800188, Id: "8f544070-db30-11e3-e046-03985e155ee2", Msisdn: "358401950828", Site: "ilmaista.rakkauden.fi", Themes: "adult", Resource: "mobilephone", Provider: "MobileSonera"},
	{Created: 1400800331, Id: "f93c3880-db30-11e3-8c5b-29d98b33d5f6", Msisdn: "358401395127", Site: "pillua.tissit.me", Themes: "adult", Resource: "mobiledesk", Provider: "MobileSonera"},
	{Created: 1400901131, Id: "f93c3880-db30-11e3-8c5b-29d98b33d5f7", Msisdn: "358451202801", Site: "porno.miia.me", Themes: "adult", Resource: "mobilephone", Provider: "MobileSonera"},
}

func TestElaborateHit(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	for _, hit := range hitsTests {

		MobClientHtml := ElaborateHit(*golog, hit)

		CreateHtmlGz.CreateFile(*golog, MobClientHtml)

	}

}
