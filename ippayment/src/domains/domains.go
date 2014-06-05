package domains

import (
	"net"
)
type Payments struct {
	Created string
	Amount string
	Result string
	Themes string

}
type SmsOutHtml struct {
	Created string
	From string
	Text string

}

type MobClientHtml struct {
	ClPhonenum string
	ClBlock    string
	ClHits     []HitHtml
	ClPayments  []Payments
	ClSmsOut   []SmsOutHtml
}

type HitHtml struct {
	
	Createdstr string
	Site     string
	Themes   string
	Resource string
	Provider string

}
type MobClient struct {
	ClPhonenum string
	ClCreated  int64
	ClUpdated  int64
	ClThemes   string
	ClResource string
	ClProvider string
	ClBlock    int
	ClHits     []Hit
	ClSmsOut   []SmsOut
	//	ColHits     int
}


type Hit struct {
	Created  int64
	Createdstr string
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
	Provider   string
	
}

type ProviderSubnet struct {
	IpNet    net.IPNet
	Provider string
}

type Provider struct {
	Provider  string
	UserAgent string
}
