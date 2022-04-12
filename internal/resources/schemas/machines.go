package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
)

func Machines() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"state": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"health": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"created_date": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"last_update_date": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"hostname": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
		"id": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
	}
}

func FlattenMachines(machines *[]mcaasapi.Machine) []interface{} {
	if machines == nil {
		return nil
	}

	machinesOut := make([]interface{}, len(*machines))
	for i, machine := range *machines {
		mc := make(map[string]interface{})

		createdDate, _ := machine.CreatedDate.MarshalText()

		lastUpdateDate, _ := machine.LastUpdateDate.MarshalText()

		mc["state"] = machine.State
		mc["health"] = machine.Health
		mc["created_date"] = string(createdDate)
		mc["last_update_date"] = string(lastUpdateDate)
		mc["name"] = machine.Name
		mc["hostname"] = machine.Hostname
		mc["id"] = machine.Id
		machinesOut[i] = mc
	}

	return machinesOut
}
