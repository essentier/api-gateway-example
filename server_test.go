package main

import (
	"testing"

	"github.com/essentier/testutil"
)

func TestTodoRestAPI(t *testing.T) {
	t.Parallel()
	gatewayService := testutil.CreateRestService("api-gateway-example", "/hello-example/hello", t)
	defer gatewayService.Release()

	var helloResult map[string]string
	gatewayService.Resource("hello-example").NewChildResource("hello").Get(&helloResult)
	t.Logf("helloResult is %v", helloResult)
}
