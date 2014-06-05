package checkgzfileexist

import (
//	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"testing"
)

var htmlfilearr = []string{"../../www/tel/358400177858.html.gz", "../../www/tel/358400177858.html", "../../www/tel/smallfile"}

func TestCheckFile(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	d1 := []byte("hello\nSmall file\n")
	err = ioutil.WriteFile("../../www/tel/smallfile", d1, 0644)
	if err != nil {
		panic(err)

	}

	for _, htmlfile := range htmlfilearr {

		check := CheckFile(*golog, htmlfile)
		t.Log(check)

	}

}
