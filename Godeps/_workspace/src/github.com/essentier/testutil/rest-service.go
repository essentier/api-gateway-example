package testutil

import (
	"testing"

	"github.com/essentier/gopencils"
	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/config"
	"github.com/essentier/spickspan/model"
	"github.com/essentier/spickspan/probe"
)

var provider model.Provider

func init() {
	config, err := config.GetConfig()
	if err != nil {
		panic("Failed to find and parse spickspan.json. The error is " + err.Error())
	}

	provider, err = spickspan.GetNomockProvider(config)
	if err != nil {
		panic("Failed to get nomock provider. The error is " + err.Error())
	}

	err = spickspan.BuildAllInConfig(config)
	if err != nil {
		panic("Failed to build projects. The error is " + err.Error())
	}
}

func CreateRestService(serviceName string, readinessPath string, t *testing.T) *restService {
	service, err := provider.GetService(serviceName)
	if err != nil {
		t.Fatalf("Failed to create service %v. Error is: %v.", serviceName, err)
	}

	errHandler := &failTestRestErrHanlder{t: t}
	api := gopencils.Api(service.GetUrl())
	rw := &resourceWrapper{resource: api, errHandler: errHandler}
	restService := &restService{provider: provider, service: service, api: rw}

	serviceReady := probe.ProbeHttpService(restService.service, readinessPath)
	if serviceReady {
		return restService
	} else {
		defer restService.Release()
		t.Fatalf("Service is not ready. The service is %v", serviceName)
		return nil
	}
}

type restService struct {
	api      *resourceWrapper
	provider model.Provider
	service  model.Service
}

func (s *restService) Release() {
	s.provider.Release(s.service)
}

func (s *restService) Resource(resourceName string) *resourceWrapper {
	return s.api.NewChildResource(resourceName)
}
