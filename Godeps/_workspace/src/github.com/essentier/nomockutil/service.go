package nomockutil

import (
	"regexp"

	"github.com/go-errors/errors"
)

var serviceNameRegexp, _ = regexp.Compile("^([a-z0-9-_.]+)$")

func ValidateServiceName(serviceName string) error {
	//service name can contain only these characters: [a-z0-9-_.]
	//This is because when we build a container image of a service,
	//the service name will be part of the container image name. And
	//a container image name can contain only [a-z0-9-_.]
	if serviceNameRegexp.MatchString(serviceName) {
		return nil
	} else {
		return errors.Errorf("Service name %v is invalid. It should contain only [a-z0-9-_.]", serviceName)
	}
}
