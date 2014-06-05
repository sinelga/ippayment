package checkgzfileexist

import (
	"fmt"
	"log/syslog"
	"os"
	"strconv"
)

func CheckFile(golog syslog.Writer, htmlfile string) bool {

	retbool := false
	var i64sizesmall int64 = 100

	fmt.Println("start CheckFile")

	finfo, err := os.Stat(htmlfile)
	if err != nil {

		return retbool
	}

	if !finfo.IsDir() {

		filessize := strconv.FormatInt(finfo.Size(), 10)
		
		if finfo.Size() < i64sizesmall {

			golog.Err("Small < 100 filessize delete !!!" + filessize)
			
			os.Remove(htmlfile) 

		} else {

			golog.Err("OK filessize " + filessize)
			retbool = true

		}
	}

	return retbool
}
