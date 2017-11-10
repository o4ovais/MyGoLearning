package av

import (
	"encoding/json"
	"lib/networking"
	"log"
	models "portal/av/model"
	"strings"
)

//ReadFinanceXML ssas ss
func ReadFinanceXML(net networking.Network) {
	res := net.GetReq(cAVURL)
	parseResponse(res)
}

func parseResponse(res networking.Response) {
	//log.Println("inside parse response")
	if res.Err != nil {
		log.Println("Error occured while reading stock data ", res.Err)
		return
	}
	var bankList models.Lists
	var newString = removeCharacterFromString(string(res.Body), "/", 2)
	err := json.Unmarshal([]byte(newString), &bankList)
	if err != nil {
		log.Println("Error occured while unmarshaling stock data ", err)
	}
	//log.Println(bankList.AntiVirusList)

	for _, res := range bankList.AntiVirusList {
		var status = "0"
		if res.AntiVirusStatus == "Synced" {
			status = "1"
		}
		log.Println("Site ID : ", res.SiteID, " - RegID : ", res.RegID, " - Status : ", res.AntiVirusStatus, " - update with : ", status)
	}
}
func removeCharacterFromString(actualString string, charRemove string, occurance int) string {
	return strings.Replace(actualString, charRemove, "", occurance)
}
