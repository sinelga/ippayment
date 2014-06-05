package htmlhandler

import (
	"log"
	"log/syslog"
	"testing"
	"fmt"
)

var htmlfile = "../../www/tel/358400177858.html.gz"

func TestParseHtmlGzFile(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	mobClientHtml :=ParseHtmlGzFile(*golog, htmlfile)
	
	for _,record := range mobClientHtml.ClPayments {
		
		fmt.Println(record)
	
	}
	
	for _,recordhit := range mobClientHtml.ClHits {
		
		fmt.Println(recordhit)
	
	}
	for _,recordsms := range mobClientHtml.ClSmsOut {
		
		fmt.Println(recordsms)
	
	}	
		
	
}
