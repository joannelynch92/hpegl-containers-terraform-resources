// (C) Copyright 2022 Hewlett Packard Enterprise Development LP

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Appliance() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"space_id": {
			Type:     schema.TypeString,
			Required: true,
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
			Required: true,
		},
		"id": {
			Type:     schema.TypeString,
			ForceNew: true,
			Computed: true,
		},
	}
}
