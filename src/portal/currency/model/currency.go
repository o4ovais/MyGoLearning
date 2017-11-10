package models

type Lists struct {
	Resources struct {
		Resource []struct {
			//SField  []string `xml:"field"`
			Field  []struct {
				////value   string `xml:"field"`
				//name   string `xml:"type,attr"`
				Key   string `xml:"name,attr"`
				Value string `xml:",chardata"`
			} `xml:"field"`
		}`xml:"resource"`
	}`xml:"resources"`
}