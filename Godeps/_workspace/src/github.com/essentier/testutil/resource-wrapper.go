package testutil

import (
	"io"
	"net/http"

	"github.com/essentier/gopencils"
)

type resourceWrapper struct {
	resource   *gopencils.Resource
	errHandler RestErrorHandler
}

func (rw *resourceWrapper) NewChildResource(resourceName string) *resourceWrapper {
	newRes := rw.resource.NewChildResource(resourceName, nil)
	newRW := &resourceWrapper{errHandler: rw.errHandler, resource: newRes}
	return newRW
}

func (rw *resourceWrapper) NewChildIdResource(id string) *resourceWrapper {
	newRes := rw.resource.NewChildIdResource(id)
	newRW := &resourceWrapper{errHandler: rw.errHandler, resource: newRes}
	return newRW
}

func (rw *resourceWrapper) SetQuery(querystring map[string]string) *resourceWrapper {
	rw.resource.SetQuery(querystring)
	return rw
}

func (rw *resourceWrapper) Get(responseBody interface{}) *resourceWrapper {
	rw.resource.Response = responseBody
	_, err := rw.resource.Get()
	rw.handleErrorIfAny(err, "REST GET failed.")
	return rw
}

func (rw *resourceWrapper) Head() *resourceWrapper {
	_, err := rw.resource.Head()
	rw.handleErrorIfAny(err, "REST HEAD failed.")
	return rw
}

func (rw *resourceWrapper) Put(payload interface{}, responseBody interface{}) *resourceWrapper {
	rw.resource.Response = responseBody
	_, err := rw.resource.Put(payload)
	rw.handleErrorIfAny(err, "REST PUT failed.")
	return rw
}

func (rw *resourceWrapper) handleErrorIfAny(err error, message string) {
	if err != nil && rw.errHandler != nil {
		rw.errHandler.HandleError(err, message)
	}
}

func (rw *resourceWrapper) Post(payload interface{}, responseBody interface{}) *resourceWrapper {
	rw.resource.Response = responseBody
	_, err := rw.resource.Post(payload)
	rw.handleErrorIfAny(err, "REST POST failed.")
	return rw
}

func (rw *resourceWrapper) Delete(responseBody interface{}) *resourceWrapper {
	rw.resource.Response = responseBody
	_, err := rw.resource.Delete()
	rw.handleErrorIfAny(err, "REST DELETE failed.")
	return rw
}

func (rw *resourceWrapper) Options(responseBody interface{}) *resourceWrapper {
	rw.resource.Response = responseBody
	_, err := rw.resource.Options()
	rw.handleErrorIfAny(err, "REST OPTIONS failed.")
	return rw
}

func (rw *resourceWrapper) Patch(payload interface{}, responseBody interface{}) *resourceWrapper {
	rw.resource.Response = responseBody
	_, err := rw.resource.Patch(payload)
	rw.handleErrorIfAny(err, "REST PATCH failed.")
	return rw
}

func (rw *resourceWrapper) SetPayload(args interface{}) io.Reader {
	return rw.resource.SetPayload(args)
}

func (rw *resourceWrapper) SetHeader(key string, value string) {
	rw.resource.SetHeader(key, value)
}

func (rw *resourceWrapper) SetClient(c *http.Client) {
	rw.resource.SetClient(c)
}
