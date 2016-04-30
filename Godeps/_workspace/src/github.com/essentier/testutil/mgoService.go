package testutil

import (
	"strconv"
	"testing"

	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/model"
)

func CreateMgoService(serviceName string, t *testing.T) *mgoService {
	service, err := spickspan.GetMongoDBService(provider, serviceName)
	if err != nil {
		t.Fatalf("Failed to create service %v. Error is: %v", serviceName, err)
	}
	mgoService := &mgoService{provider: provider, service: service}
	return mgoService
}

type mgoService struct {
	provider model.Provider
	service  model.Service
}

func (s *mgoService) Release() {
	s.provider.Release(s.service)
}

func (s *mgoService) GetUrl() string {
	return s.service.IP + ":" + strconv.Itoa(s.service.Port)
}
