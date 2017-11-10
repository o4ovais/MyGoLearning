package currency

import (
	"encoding/xml"
	"portal/currency/model"
	"lib/networking"
	"log"
)

func ReadFinanceXML(net networking.Network) {
	res := net.GetReq(cCurrencyUrl)
	parseResponse(res)
}

func parseResponse(res networking.Response) {
	if res.Err != nil {
		log.Println(res.Err)
		return
	}
	var finApiXML models.Lists
	err := xml.Unmarshal(res.Body, &finApiXML)
	if err != nil {
		log.Println("Error occured while unmarshalling xml file", err)
	}
	iterateResponse(finApiXML, cCurrencyCode)
}

func iterateResponse(finApiXML models.Lists, desiredCurrency string) {

	for _, res := range finApiXML.Resources.Resource {
		currRateMap := make(map[string]string)
		for _, field := range res.Field {
			currRateMap[field.Key] = field.Value
		}

		if currRateMap["name"] == desiredCurrency {
			log.Println("Currency Code:", desiredCurrency, " Rate :", currRateMap["price"])
			break
		}

	}
}
