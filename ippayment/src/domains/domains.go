package domains

import ()

type MobClient struct {
	ClPhonenum string
	ClCreated  int64
	ClUpdated  int64
	ClThemes   string
	ClResource string
	ClProvider string
	ClBlock    int
	ClHits       []Hit
	ClSmsOut  []SmsOut
	//	ColHits     int
}

type Hit struct {
	Created  int64
	Id       string
	Msisdn   string
	Site     string
	Themes   string
	Resource string
	Provider string
}

type SmsOut struct {
	SmsCreated int64
	Msisdn     string
	From       string
	Text       string
	Provider	string
}
