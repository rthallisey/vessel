package app

import (
	"time"

	log "github.com/vessel/pkg/logger"
	"github.com/vessel/pkg/tpr"
)

func Run() *int {
	Tpr := tpr.NewTPR("https://192.168.121.11:6443/apis/nave.vessel/v1/namespaces/vessels/servicevessels/mariadb-vessel")
	status := Tpr.Read()

	for {
		log.Info("Starting main loop")
		a := new(int)
		time.Sleep(time.Duration(1) * time.Second)
		return a
	}
}
