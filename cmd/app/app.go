package app

import (
	"time"

	log "github.com/vessel/pkg/logger"
	"github.com/vessel/pkg/monitor"
	"github.com/vessel/pkg/tpr"
)

func Run() *int {
	_, _, Info := log.NewLogger()
	status, _, _ := tpr.Read("https://192.168.121.29:6443/apis/nave.vessel/v1/namespaces/vessels/servicevessels/mariadb-vessel")
	Info.Print(status)

	elk := monitor.Monitor{ContainerCount: 0}

	for {
		Info.Print("Starting main loop")
		elk.CheckContainerCount()
		time.Sleep(time.Duration(1) * time.Second)
	}
	a := new(int)
	return a
}
