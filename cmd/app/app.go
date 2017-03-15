package app

import (
	"time"

	log "github.com/vessel/pkg/logger"
	"github.com/vessel/pkg/tpr"
)

func Run() *int {
	Tpr := tpr.NewTPR("https://192.168.121.11:6443/apis/nave.vessel/v1/namespaces/vessels/servicevessels/mariadb-vessel")
	status := Tpr.Read()

	_, _, Info := log.NewLogger()

	Info.Print(status)

	for {
		Info.Print("Starting main loop")
		time.Sleep(time.Duration(1) * time.Second)
		a := new(int)
		return a
	}
}
