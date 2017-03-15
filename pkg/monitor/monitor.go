package monitor

import (
	"encoding/json"

	log "github.com/vessel/pkg/logger"
	"github.com/vessel/pkg/transport"
)

var _, _, Info = log.NewLogger()

type Monitor struct {
	ContainerCount int
	Spec           map[string]interface{}
}

func (m Monitor) CheckContainerCount() string {

	Info.Print("Checking container count")
	url := "https://192.168.121.29:6443/api/v1/namespaces/vessels/replicationcontrollers/busybox-sleep-rc"
	t := transport.NewTransport(url)
	body, status := t.Get()

	json.Unmarshal(body, &m.Spec)
	podStatus := m.Spec["status"].(map[string]interface{})
	Info.Print(podStatus)
	return status
}
