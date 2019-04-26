// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcloud

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-hcloud/hcloud"
)

// all of the token components used below.
const (
	// packages:
	hcloudPkg = "hcloud"
	// modules:
	hcloudMod = "index"
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(hcloudPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := hcloud.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "hcloud",
		Description:          "A Pulumi package for creating and managing hcloud cloud resources.",
		Keywords:             []string{"pulumi", "hcloud"},
		License:              "MIT",
		Homepage:             "https://github.com/akoenig/pulumi-hcloud",
		Repository:           "https://github.com/akoenig/pulumi-hcloud",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"hcloud_server":                 {Tok: makeResource(hcloudMod, "Server")},
			"hcloud_floating_ip":            {Tok: makeResource(hcloudMod, "FloatingIp")},
			"hcloud_floating_ip_assignment": {Tok: makeResource(hcloudMod, "FloatingIpAssignment")},
			"hcloud_rdns":                   {Tok: makeResource(hcloudMod, "Rdns")},
			"hcloud_ssh_key":                {Tok: makeResource(hcloudMod, "SshKey")},
			"hcloud_volume":                 {Tok: makeResource(hcloudMod, "Volume")},
			"hcloud_volume_attachment":      {Tok: makeResource(hcloudMod, "VolumeAttachment")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"hcloud_datacenter":  {Tok: makeDataSource(hcloudMod, "getDatacenter")},
			"hcloud_datacenters": {Tok: makeDataSource(hcloudMod, "getDatacenters")},
			"hcloud_floating_ip": {Tok: makeDataSource(hcloudMod, "getFloatingIp")},
			"hcloud_image":       {Tok: makeDataSource(hcloudMod, "getImage")},
			"hcloud_location":    {Tok: makeDataSource(hcloudMod, "getLocation")},
			"hcloud_locations":   {Tok: makeDataSource(hcloudMod, "getLocations")},
			"hcloud_server":      {Tok: makeDataSource(hcloudMod, "getServer")},
			"hcloud_ssh_key":     {Tok: makeDataSource(hcloudMod, "getSshKey")},
			"hcloud_volume":      {Tok: makeDataSource(hcloudMod, "getVolume")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "latest",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=0.16.4,<0.17.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
