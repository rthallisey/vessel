package transport

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/vessel/pkg/logger"
)

var Error, _, Info = log.NewLogger()

func Get(url string, token string) ([]byte, string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{Transport: tr}

	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)

	res, err := client.Do(req)
	if err != nil {
		Error.Print("Http request failed")
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(res.Body)
	Info.Print(string(body))

	return body, res.Status
}
