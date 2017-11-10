package models

//Lists will be the structure for the
type Lists struct {
	AntiVirusList []struct {
		SiteID          int    `json:"siteId"`
		PartnerID       int    `json:"partnerId"`
		RegID           int    `json:"regId"`
		AntiVirusStatus string `json:"antiVirusStatus"`
	} `json:"antiVirusList"`
	TotalCount int `json:"totalCount"`
}
