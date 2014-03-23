package checkipnet

import (
	"fmt"
	"net"
)

func Check(ipstr string, fieldsarr [][]string) {

	for _, field := range fieldsarr {

		_, ipnet, _ := net.ParseCIDR(field[0])
		ipo := net.ParseIP(ipstr)

		if ipnet.Contains(ipo) {

			fmt.Println("OK provider ",field[1])

		} else {
		
			fmt.Println("NotMobile")
			
		}

	}

}
