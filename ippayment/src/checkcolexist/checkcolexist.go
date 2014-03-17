package checkcolexist

import ()

func ChecExist(collections []string, msisdn string) bool {

	var exist bool = false

	for _, msisdnloop := range collections {

		if msisdnloop == msisdn {
			exist = true

		}

	}

	return exist

}
