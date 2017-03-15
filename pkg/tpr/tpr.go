package tpr

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/vessel/pkg/logger"
	"github.com/vessel/pkg/transport"
)

type TPR struct {
	Url        string
	Body       string
	Token      string
	Spec       map[string]interface{}
	VesselSpec map[string]interface{}
}

var Error, _, Info = log.NewLogger()

func NewTPR(url string) *TPR {
	b, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		Error.Print("Missing token file")
		os.Exit(1)
	}
	token := string(b)

	return &TPR{Token: token, Url: url}
}

func (tpr *TPR) Read() string {

	body, status := transport.Get(tpr.Url, tpr.Token)
	tpr.Body = string(body)

	json.Unmarshal(body, &tpr.Spec)
	tpr.VesselSpec = tpr.Spec["vesselSpec"].(map[string]interface{})

	return status
}
