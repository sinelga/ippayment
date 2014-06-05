package CreateHtmlGz

import (
	"bytes"
	"compress/gzip"
	"domains"
	"fmt"
	"html/template"
	"io/ioutil"
	"log/syslog"
)

func CreateFile(golog syslog.Writer, hit domains.MobClientHtml) {

	fmt.Println(hit.ClPhonenum)
	var htmlpage = template.Must(template.ParseFiles(
		"/home/juno/git/ippayment/ippayment/templ/_base.html",
		"/home/juno/git/ippayment/ippayment/templ/contents.html",
	))
	webpage := bytes.NewBuffer(nil)
	if err := htmlpage.Execute(webpage, hit); err != nil {
		golog.Err(err.Error())

	}
	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	var b bytes.Buffer

	w := gzip.NewWriter(&b)
	w.Write(webpagebytes)
	w.Close()

	if err := ioutil.WriteFile("/home/juno/git/ippayment/ippayment/www/tel/"+hit.ClPhonenum+".html.gz", b.Bytes(), 0666); err != nil {
		golog.Crit(err.Error())

	}

}
