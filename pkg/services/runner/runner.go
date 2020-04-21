package servicerunner

import (
	"github.com/ovh/utask/engine/step"
	"github.com/ovh/utask/pkg/services"
)

func Init() error {
	for _, serviceName := range services.List() {
		service, _ := services.Get(serviceName)
		if err := step.RegisterRunner(serviceName, service); err != nil {
			return err
		}
	}
	return nil
}
