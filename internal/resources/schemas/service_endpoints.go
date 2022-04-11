package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
)

func ServiceEndpoints() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"endpoint": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"namespace": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
	}
}

func FlattenServiceEndpoints(serviceEndpoints *[]mcaasapi.ServiceEndpoints) []interface{} {
	if serviceEndpoints == nil {
		return nil
	}

	serviceSets := make([]interface{}, len(*serviceEndpoints))
	for i, service := range *serviceEndpoints {
		serv := make(map[string]interface{})

		serv["endpoint"] = service.Endpoint
		serv["name"] = service.Name
		serv["namespace"] = service.Namespace
		serv["type"] = service.Type_
		serviceSets[i] = serv
	}

	return serviceSets
}
