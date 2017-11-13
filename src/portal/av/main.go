package av

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"strings"
)

//AVStatusLists
type AVStatusLists struct {
	AntiVirusList []struct {
		SiteID          int    `json:"siteId"`
		PartnerID       int    `json:"partnerId"`
		RegID           int    `json:"regId"`
		AntiVirusStatus string `json:"antiVirusStatus"`
	} `json:"antiVirusList"`
	TotalCount int `json:"totalCount"`
}

//ReadAVStatusJSON
func ReadAVStatusJSON() {
	AVURL := "https://rmmapi.dtitsupport247.net/itswebapi/v1/partner/50000001/antivirusstatus?resourceType=desktop&regIds=50046431,50047691"
	res := getReq(AVURL)
	parseResponse(res)
}

func parseResponse(res Response) {
	//log.Println("inside parse response")
	if res.Err != nil {
		log.Println("Error occured while reading stock data ", res.Err)
		return
	}
	var avStatusList AVStatusLists
	var newString = removeCharacterFromString(string(res.Body), "/", 2)
	err := json.Unmarshal([]byte(newString), &avStatusList)
	if err != nil {
		log.Println("Error occured while unmarshaling stock data ", err)
	}
	//log.Println(bankList.AntiVirusList)

	for _, res := range avStatusList.AntiVirusList {
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

func getReq(url string) Response {

	resp, err := http.Get(url)
	if err != nil {
		log.Println("error occured while reading from url", url, err)

		return consolidateResponse(nil, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//log.Println("response:", body)

	return consolidateResponse(body, err)
}

func consolidateResponse(body []byte, err error) Response {
	res := Response{
		Body: body,
		Err:  err,
	}
	return res
}

//Response of the API
type Response struct {
	Body []byte
	Err  error
}
