package tpr

import (
	"encoding/json"

	log "github.com/vessel/pkg/logger"
	"github.com/vessel/pkg/transport"
)

var Error, _, Info = log.NewLogger()
var Spec map[string]interface{}
var VesselSpec map[string]interface{}

func Read(url string) (string, map[string]interface{}, map[string]interface{}) {

	t := transport.NewTransport(url)
	body, status := t.Get()

	json.Unmarshal(body, &Spec)
	VesselSpec = Spec["vesselSpec"].(map[string]interface{})

	return status, Spec, VesselSpec
}
