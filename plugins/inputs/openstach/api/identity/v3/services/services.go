package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/influxdata/telegraf/plugins/inputs/openstach/api/identity/v3"
	"io/ioutil"
	"net/http"
)
// Service represents an OpenStack Service.
type Service struct {
	// ID is the unique ID of the service.
	ID string `json:"id"`

	// Type is the type of the service.
	Type string `json:"type"`

	// Enabled is whether or not the service is enabled.
	Enabled bool `json:"enabled"`

	// Links contains referencing links to the service.
	Links struct{ Self string `json:"self"` }
}

func List(client *v3.IdentityClient) ([]Service, error){
	api := declareListService(client.Token)

	jsonBody, err := json.Marshal(api.Request)

	if err != nil {
		panic(err.Error())
	}

	httpClient := &http.Client{}
	request, err := http.NewRequest(api.Method, client.Endpoint+api.Path, bytes.NewBuffer(jsonBody))
	for k, v := range api.Header {
		request.Header.Add(k,v)
	}
	resp, err := httpClient.Do(request)
	defer resp.Body.Close()

	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("List service successful ")
	} else {
		err := errors.New("List service respond status code "+ string(resp.StatusCode))
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(body), &api.Response)

	services := []Service{}
	for _,v := range api.Response.Services{
		services = append(services, Service{
			ID: v.ID,
			Type: v.Type,
			Enabled: v.Enabled,
			Links: v.Links,
		})
	}

	return services, err
}