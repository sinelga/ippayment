package CreateHtmlGz

import (
	"testing"
	"domains"
	"log"
	"log/syslog"
	
)
var smsoutTest = []domains.SmsOutHtml{
	{Created:"2014/06/23 10:00:00",From: "07003838338",Text:"Soita nyt!"},
	{Created:"2014/06/24 11:00:00",From: "07003838338",Text:"Soita nyt 222!"},
}

var paymentTest = []domains.Payments{
	{Created:"2014/06/23 10:00:00",Amount: "87",Result: "OK"},
	{Created:"2014/06/24 12:00:00",Amount: "87",Result: "OK"},
}

var hitsTests = []domains.HitHtml{
//	{Createdstr: "2014/06/23 10:00:00", Id: "cd71e4c0-db30-11e3-c47b-71a9d04148af", Msisdn: "358400177858", Site: "tissit.tv", Themes: "adult", Resource: "mobilephone", Provider: "MobileSonera"},
//	{Createdstr: "2014/06/24 13:00:00", Id: "cd71e4c0-db30-11e3-c47b-71a9d04148af", Msisdn: "358400177858", Site: "tissit.tv", Themes: "adult", Resource: "mobilephone", Provider: "MobileSonera"},
	{Createdstr: "2014/06/23 10:00:00", Site: "tissit.tv", Themes: "adult", Resource: "mobilephone", Provider: "MobileSonera"},
	{Createdstr: "2014/06/24 13:00:00", Site: "pillu.tv", Themes: "adult", Resource: "mobilephone", Provider: "MobileSonera"},


}

var  MobClientHtmlTests = domains.MobClientHtml{
	ClPhonenum:  "358400177858", ClBlock: "No", ClHits: hitsTests, ClPayments: paymentTest,ClSmsOut:smsoutTest,

}


func TestCreateFile(t *testing.T) {
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	CreateFile(*golog,MobClientHtmlTests)


}
