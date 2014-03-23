package loadsubnets

import (
	"encoding/csv"
	"log/syslog"
	"os"
)

func Load(golog syslog.Writer, file string) [][]string {

	csvFile, err := os.Open(file)
	defer csvFile.Close()
	if err != nil {

		golog.Crit("loadsubnets: " + err.Error())
	}
	var fieldsarr [][]string
	csvReader := csv.NewReader(csvFile)

	lines, err := csvReader.ReadAll()
	if err != nil {

		golog.Crit("loadsubnets: " + err.Error())
	}
	for _, line := range lines {

		fieldsarr = append(fieldsarr, line)

	}

	return fieldsarr
}
