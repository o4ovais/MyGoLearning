package networking

import (
	"net/http"
	"io/ioutil"
	"log"
)

type NetworkImpl struct {
}

func GetInstance() Network {
	return new(NetworkImpl)
}

func (net *NetworkImpl) GetReq(url string) Response {
	ch := getReq(url)
	return <-ch
}

func getReq(url string) chan Response {
	chanOut := make(chan Response)
	go func() {
		defer close(chanOut)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("error occured while reading from url", url, err)
			chanOut <- consolidateResponse(nil, err)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		//log.Println("response:", body)

		chanOut <- consolidateResponse(body, err)
	}()
	return chanOut
}

func consolidateResponse(body []byte, err error) Response {
	res := Response{
		Body: body,
		Err:  err,
	}
	return res
}
