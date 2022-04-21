# Copyright 2020 Hewlett Packard Enterprise Development LP

provider hpegl {
  caas {
    api_url = "https://client.greenlake.hpe.com/api/caas/mcaas/v1"
  }
}

resource hpegl_caas_cluster test {
  name         = var.cluster_name
  blueprint_id = ""
  appliance_id = ""
  space_id     = ""
}
