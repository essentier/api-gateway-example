package main

import (
	"testing"

	"github.com/essentier/testutil"
)

func TestTodoRestAPI(t *testing.T) {
	t.Parallel()
	gatewayService := testutil.CreateRestService("api-gateway-example", "/hello-rest/hello", t)
	defer gatewayService.Release()

	var helloResult map[string]string
	gatewayService.Resource("hello-rest").Resource("hello").Get(&helloResult)
	t.Logf("helloResult is %v", helloResult)

	// api := gopencils.Api(hostUrl)
	// var result todo.Todos
	// todosRes := api.Res("todo-rest/todos", &result)
	// _, err := todosRes.Get()
	// if err != nil {
	// 	t.Errorf("Failed to call the todo rest api. Error is: %#v. Error string is %v", err, err.Error())
	// }
	// log.Printf("todos are: %#v", result)
}

// func TestLogin(t *testing.T) {
// 	token := loginToEssentier("http://10.20.132.206:8083", "user1@user1.com", "user1password")
// 	log.Printf("toke: %v", token)
// }

// func loginToEssentier(url, username, password string) string {
// 	essentierRest := gopencils.Api(url + "/essentier-rest")
// 	token := &jwtauth.TokenAuthentication{}
// 	loginData := &jwtauth.LoginData{Email: username, Password: password}
// 	_, err := essentierRest.Res("login", token).Post(loginData)
// 	if err != nil {
// 		log.Printf("Failed to call the login rest api. Error is: %#v", err)
// 	}
// 	log.Printf("Received token is: %#v", token)
// 	return token.Token
// }

// func TestNomockServerRestAPI(t *testing.T) {
// 	hostUrl := "http://10.20.132.206:8083"
// 	api := gopencils.Api(hostUrl)
// 	var result spickspan.Service
// 	servicesResource := api.Res("nomockserver/services", &result)

// 	jwtService := jwtauth.CreateJWTService()
// 	token, _ := jwtService.GenerateToken("user1")
// 	servicesResource.SetHeader("Authorization", "Bearer " + token)

// 	serviceConfig := config.Service{
// 		ServiceName: "testnomockservermongo1",
// 		ContainerImage: "mongo",
// 		ProjectSrcRoot: "/",
// 		Port: 27017,
// 	}

// 	_, err := servicesResource.Post(serviceConfig)
// 	if err != nil {
// 		t.Errorf("Failed to call the service rest api. Error is: %#v. Error string is %v", err, err.Error())
// 	}
// 	log.Printf("service is: %#v", result)

// 	servicesResource = api.Res("nomockserver/services")
// 	servicesResource = servicesResource.Id(result.Id)
// 	servicesResource.SetHeader("Authorization", "Bearer " + token)
// 	servicesResource.Delete()
// }
