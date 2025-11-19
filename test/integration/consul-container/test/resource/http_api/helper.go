// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/test/resource/http_api/helper.go
Helper module for internal layer

## Linked Modules
- [test/integration/consul-container/libs/cluster](../test/integration/consul-container/libs/cluster)
- [test/integration/consul-container/libs/topology](../test/integration/consul-container/libs/topology)
- [test/integration/consul-container/test/resource/http_api/client](../test/integration/consul-container/test/resource/http_api/client)

## Tags
api, client, internal

## Exports
GVK, ListResponse, Resource, SetupClusterAndClient, Terminate, WriteRequest

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/test/resource/http_api/helper.go> a code:Module ;
    code:name "test/integration/consul-container/test/resource/http_api/helper.go" ;
    code:description "Helper module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "test/integration/consul-container/libs/cluster" ;
        code:path "../test/integration/consul-container/libs/cluster"
    ], [
        code:name "test/integration/consul-container/libs/topology" ;
        code:path "../test/integration/consul-container/libs/topology"
    ], [
        code:name "test/integration/consul-container/test/resource/http_api/client" ;
        code:path "../test/integration/consul-container/test/resource/http_api/client"
    ] ;
    code:exports :GVK, :ListResponse, :Resource, :SetupClusterAndClient, :Terminate, :WriteRequest ;
    code:tags "api", "client", "internal" .
<!-- End LinkedDoc RDF -->
*/
package resource

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	libcluster "github.com/hashicorp/consul/test/integration/consul-container/libs/cluster"
	libtopology "github.com/hashicorp/consul/test/integration/consul-container/libs/topology"
	client "github.com/hashicorp/consul/test/integration/consul-container/test/resource/http_api/client"
)

func makeClusterConfig(numOfServers int, numOfClients int, aclEnabled bool) *libtopology.ClusterConfig {
	return &libtopology.ClusterConfig{
		NumServers:  numOfServers,
		NumClients:  numOfClients,
		LogConsumer: &libtopology.TestLogConsumer{},
		BuildOpts: &libcluster.BuildOptions{
			Datacenter: "dc1",
			ACLEnabled: aclEnabled,
		},
		ApplyDefaultProxySettings: false,
	}
}

type Resource struct {
	HttpClient *client.HttpClient
}

type GVK struct {
	Group   string
	Version string
	Kind    string
}

var demoGVK = GVK{
	Group:   "demo",
	Version: "v2",
	Kind:    "Artist",
}

var defaultTenancyQueryOptions = client.QueryOptions{
	Namespace: "default",
	Partition: "default",
	Peer:      "local",
}

type WriteRequest struct {
	Metadata map[string]string
	Data     map[string]any
}

var demoPayload = WriteRequest{
	Metadata: map[string]string{
		"foo": "bar",
	},
	Data: map[string]any{
		"name": "cool",
	},
}

type config struct {
	gvk          GVK
	resourceName string
	queryOptions client.QueryOptions
	payload      WriteRequest
}

type operation struct {
	action           func(client *Resource, config config) error
	expectedErrorMsg string
	includeToken     bool
}

type testCase struct {
	description string
	operations  []operation
	config      []config
}

var applyResource = func(resource *Resource, config config) error {
	_, err := resource.Apply(&config.gvk, config.resourceName, &config.queryOptions, &config.payload)
	return err
}
var readResource = func(resource *Resource, config config) error {
	_, err := resource.Read(&config.gvk, config.resourceName, &config.queryOptions)
	return err
}
var deleteResource = func(resource *Resource, config config) error {
	err := resource.Delete(&config.gvk, config.resourceName, &config.queryOptions)
	return err
}
var listResource = func(resource *Resource, config config) error {
	_, err := resource.List(&config.gvk, &config.queryOptions)
	return err
}

func (resource *Resource) Read(gvk *GVK, resourceName string, q *client.QueryOptions) (map[string]interface{}, error) {
	r := resource.HttpClient.NewRequest("GET", strings.ToLower(fmt.Sprintf("/api/%s/%s/%s/%s", gvk.Group, gvk.Version, gvk.Kind, resourceName)))
	r.SetQueryOptions(q)
	_, resp, err := resource.HttpClient.DoRequest(r)
	if err != nil {
		return nil, err
	}
	defer client.CloseResponseBody(resp)
	if err := client.RequireOK(resp); err != nil {
		return nil, err
	}

	var out map[string]interface{}
	if err := client.DecodeBody(resp, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (resource *Resource) Delete(gvk *GVK, resourceName string, q *client.QueryOptions) error {
	r := resource.HttpClient.NewRequest("DELETE", strings.ToLower(fmt.Sprintf("/api/%s/%s/%s/%s", gvk.Group, gvk.Version, gvk.Kind, resourceName)))
	r.SetQueryOptions(q)
	_, resp, err := resource.HttpClient.DoRequest(r)
	if err != nil {
		return err
	}
	defer client.CloseResponseBody(resp)
	if err := client.RequireHttpCodes(resp, http.StatusNoContent); err != nil {
		return err
	}
	return nil
}

func (resource *Resource) Apply(gvk *GVK, resourceName string, q *client.QueryOptions, payload *WriteRequest) (*map[string]interface{}, error) {
	url := strings.ToLower(fmt.Sprintf("/api/%s/%s/%s/%s", gvk.Group, gvk.Version, gvk.Kind, resourceName))

	r := resource.HttpClient.NewRequest("PUT", url)
	r.SetQueryOptions(q)
	r.Obj = payload
	_, resp, err := resource.HttpClient.DoRequest(r)
	if err != nil {
		return nil, err
	}
	defer client.CloseResponseBody(resp)
	if err := client.RequireOK(resp); err != nil {
		return nil, err
	}

	var out map[string]interface{}

	if err := client.DecodeBody(resp, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

type ListResponse struct {
	Resources []map[string]interface{} `json:"resources"`
}

func (resource *Resource) List(gvk *GVK, q *client.QueryOptions) (*ListResponse, error) {
	r := resource.HttpClient.NewRequest("GET", strings.ToLower(fmt.Sprintf("/api/%s/%s/%s", gvk.Group, gvk.Version, gvk.Kind)))
	r.SetQueryOptions(q)
	_, resp, err := resource.HttpClient.DoRequest(r)
	if err != nil {
		return nil, err
	}
	defer client.CloseResponseBody(resp)
	if err := client.RequireOK(resp); err != nil {
		return nil, err
	}

	var out *ListResponse
	if err := client.DecodeBody(resp, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func SetupClusterAndClient(t *testing.T, clusterConfig *libtopology.ClusterConfig, getServerHttpClient bool) (*libcluster.Cluster, *client.HttpClient) {
	cluster, _, _ := libtopology.NewCluster(t, clusterConfig)

	// create a http api client for resource service
	var resourceHttpClient *client.HttpClient
	if getServerHttpClient {
		apiClientConfig := cluster.Servers()[0].GetAPIClientConfig()
		apiClientConfig.Token = ""
		resourceClient, err := client.NewClient(&apiClientConfig)
		require.NoError(t, err)

		resourceHttpClient = resourceClient
	} else {
		apiClientConfig := cluster.Clients()[0].GetAPIClientConfig()
		apiClientConfig.Token = ""
		resourceClient, err := client.NewClient(&apiClientConfig)
		require.NoError(t, err)

		resourceHttpClient = resourceClient
	}

	return cluster, resourceHttpClient
}

func Terminate(t *testing.T, cluster *libcluster.Cluster) {
	err := cluster.Terminate()
	require.NoError(t, err)
}
