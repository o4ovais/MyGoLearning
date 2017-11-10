package stock

import (
	"encoding/json"
	"strings"
	"portal/stock/model"
	"lib/networking"
	"log"
)

func ReadGoogleFinanceXML(net networking.Network) {
	res := net.GetReq(cStockUrl)
	parseResponse(res)
}

func parseResponse(res networking.Response) {
	//log.Println("inside parse response")
	if res.Err != nil {
		log.Println("Error occured while reading stock data ", res.Err)
		return
	}
	var bankList []models.Bank
	var newString = removeCharacterFromString(string(res.Body), "/", 2)
	err := json.Unmarshal([]byte(newString), &bankList)
	if err != nil {
		log.Println("Error occured while unmarshaling stock data ", err)
	}
	log.Println(bankList)
}

func removeCharacterFromString(actualString string, charRemove string, occurance int) string {
	return strings.Replace(actualString, charRemove, "", occurance)
}
