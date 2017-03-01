package tpr

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/vessel/pkg/logger"
)

type TPR struct {
	Url        string
	Body       string
	Token      string
	Spec       map[string]interface{}
	VesselSpec map[string]interface{}
}

func NewTPR(url string) *TPR {
	b, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		log.Error("Missing token file")
		os.Exit(1)
	}
	token := string(b)

	return &TPR{Token: token, Url: url}
}

func (tpr *TPR) Read() string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, _ := http.NewRequest("GET", tpr.Url, nil)
	client := &http.Client{Transport: tr}

	bearer := "Bearer " + tpr.Token
	req.Header.Set("Authorization", bearer)

	res, err := client.Do(req)
	if err != nil {
		log.Error("Http request failed")
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(res.Body)
	tpr.Body = string(body)

	json.Unmarshal(body, &tpr.Spec)
	tpr.VesselSpec = tpr.Spec["vesselSpec"].(map[string]interface{})

	return res.Status
}
