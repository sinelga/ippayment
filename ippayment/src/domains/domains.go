package domains

import ()

type Collection struct {
	ColPhonenum string
	ColCreated  int64
	ColUpdated  int64
	ColThemes   string
	ColResource string
	ColHits     int
}

type Hits struct {
	Created  int64
	Id       string
	Msisdn   string
	Site     string
	Themes   string
	Resource string
}


type SmsOut struct {
	SmsCreated int64
	Msisdn string
	From string
	Text string
}