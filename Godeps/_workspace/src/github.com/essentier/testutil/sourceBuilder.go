package testutil

import (
	"github.com/essentier/servicebuilder"
	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/config"
	"github.com/essentier/spickspan/model"
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

	err = servicebuilder.BuildAllInConfig(config)
	if err != nil {
		panic("Failed to build projects. The error is " + err.Error())
	}
}
