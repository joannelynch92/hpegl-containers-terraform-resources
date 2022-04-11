# Copyright 2020 Hewlett Packard Enterprise Development LP

provider hpegl {
  caas {
    api_url = "https://client.greenlake.hpe.com/api/caas/mcaas/v1"
  }
}

resource hpegl_caas_cluster test {
  name         = var.cluster_name
  blueprint_id = "935bc1ed-ae52-41b0-b577-9527eaca2885"
  appliance_id = "233eead2-20de-47ab-b266-2413cdaa3685"
  space_id     = "f866c9bd-2d2c-4e60-aab0-64737df96273"
}
