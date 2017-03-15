package transport

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/vessel/pkg/logger"
)

var Error, _, Info = log.NewLogger()

type Transport struct {
	Url   string
	Body  string
	Token string
}

func NewTransport(url string) *Transport {
	b, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		Error.Print("Missing token file")
		os.Exit(1)
	}
	token := string(b)

	return &Transport{Token: token, Url: url}
}

func (t Transport) Get() ([]byte, string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, _ := http.NewRequest("GET", t.Url, nil)
	client := &http.Client{Transport: tr}

	bearer := "Bearer " + t.Token
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
