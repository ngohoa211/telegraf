package services

import (
	"encoding/json"
	v2 "github.com/influxdata/telegraf/plugins/inputs/openstach/api/compute/v2"
)

type Service struct{
	ID             int    `json:"id"`
	Binary         string `json:"binary"`
	DisabledReason string `json:"disabled_reason"`
	Host           string `json:"host"`
	State          string `json:"state"`
	Status         string `json:"status"`
	UpdatedAt      string `json:"updated_at"`
	ForcedDown     bool   `json:"forced_down"`
	Zone           string `json:"zone"`
}
func List(client *v2.ComputeClient) ([]Service, error) {
	api, err := declareListService(client.Endpoint, client.Token)
	err = api.DoReuest()
	result := ListServiceResponse{}
	err = json.Unmarshal([]byte(api.Response),&result)
	services := []Service{}
	for _, v := range result.Services {
		services = append(services, Service{
			ID: v.ID,
			Binary: v.Binary,
			DisabledReason: v.DisabledReason,
			Host: v.Host,
			State: v.State,
			UpdatedAt: v.UpdatedAt,
			ForcedDown: v.ForcedDown,
			Zone: v.Zone,
		})
	}
	return services, err
}