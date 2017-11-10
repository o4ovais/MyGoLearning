package main

import (
	"lib/networking"
	"portal/av"
	"portal/currency"
	"portal/stock"
)

const (
	STOCK    = 1
	CURRENCY = 2
	AV       = 3
)

func main() {

	network := networking.GetInstance()
	networking.SendTestMail()
	module := AV
	switch module {
	case STOCK:
		stock.ReadGoogleFinanceXML(network)
	case CURRENCY:
		currency.ReadFinanceXML(network)
	case AV:
		av.ReadFinanceXML(network)
	}
}
