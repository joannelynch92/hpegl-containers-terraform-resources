// (C) Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"

	"github.com/hewlettpackard/hpegl-provider-lib/pkg/client"

	"github.com/HewlettPackard/hpegl-containers-terraform-resources/pkg/constants"
)

// keyForGLClientMap is the key in the map[string]interface{} that is passed down by hpegl used to store *Client
// This must be unique, hpegl will error-out if it isn't
const keyForGLClientMap = "caasClient"

// Assert that InitialiseClient satisfies the client.Initialisation interface
var _ client.Initialisation = (*InitialiseClient)(nil)

// Client is the client struct that is used by the provider code
type Client struct {
	CaasClient *mcaasapi.APIClient
}

// InitialiseClient is imported by hpegl from each service repo
type InitialiseClient struct{}

// NewClient takes an argument of all of the provider.ConfigData, and returns an interface{} and error
// If there is no error interface{} will contain *Client.
// The hpegl provider will put *Client at the value of keyForGLClientMap (returned by ServiceName) in
// the map of clients that it creates and passes down to provider code.  hpegl executes NewClient for each service.
func (i InitialiseClient) NewClient(r *schema.ResourceData) (interface{}, error) {
	// Get CaaS settings from the CaaS block
	caasProviderSettings, err := client.GetServiceSettingsMap(constants.ServiceName, r)
	if err != nil {
		return nil, nil
	}
	apiURL := caasProviderSettings[constants.APIURL].(string)

	caasCfg := mcaasapi.Configuration{
		BasePath:      apiURL,
		DefaultHeader: make(map[string]string),
		UserAgent:     "hpegl-terraform",
	}

	cli := new(Client)
	cli.CaasClient = mcaasapi.NewAPIClient(&caasCfg)

	return cli, nil
}

// ServiceName is used to return the value of keyForGLClientMap, for use by hpegl
func (i InitialiseClient) ServiceName() string {
	return keyForGLClientMap
}

// GetClientFromMetaMap is a convenience function used by provider code to extract *Client from the
// meta argument passed-in by terraform
func GetClientFromMetaMap(meta interface{}) (*Client, error) {
	cli := meta.(map[string]interface{})[keyForGLClientMap]
	if cli == nil {
		return nil, fmt.Errorf("client is not initialised, make sure that caas block is defined in hpegl stanza")
	}

	return cli.(*Client), nil
}
