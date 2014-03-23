package main 

import (
    "flag"
    "fmt"
    "log"
    "log/syslog"
    "loadsubnets"
    "checkipnet"
    "strings"
    
    
)

var ipflag *string = flag.String("ip", "", "Ip like 85.76.65.103")

func main() {
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
    flag.Parse() // Scan the arguments list 

    if *ipflag == ""  {
        fmt.Println("try -h")
    } else {
    	
    	ip := *ipflag
    	ipstr :=strings.Split(ip,":")
    	
 		fieldsarr := loadsubnets.Load(*golog,"allsubnet.csv")
 		
 		for _,fields := range fieldsarr {
 		
 			fmt.Println(fields)
 			 			
 		
 		}
 		
 		checkipnet.Check(ipstr[0],fieldsarr)
 		
 		
 		
    
    }
    
    
}

