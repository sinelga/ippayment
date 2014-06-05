package makepayment

import (
	"fmt"
	"log"
	"log/syslog"
	"testing"
)

var clphonenumtest = "358451202801"

//var clphonenumtest = "358404859709"

var providertest = "MobileSonera"
var amounttest = "70"
var themestest = "adult"

func TestPay(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	paymentest := Pay(*golog, clphonenumtest, providertest, amounttest,themestest )

	fmt.Println("test payment-> ", paymentest.Created, paymentest.Amount, paymentest.Result)
	if paymentest.Result != "0" {
	
		t.Fail()
	
	} 
	

}
