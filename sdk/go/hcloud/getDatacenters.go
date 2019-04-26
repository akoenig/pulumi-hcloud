// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hcloud

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a list of available Hetzner Cloud Datacenters.
// This resource may be useful to create highly available infrastructure, distributed across several datacenters.
func LookupDatacenters(ctx *pulumi.Context, args *GetDatacentersArgs) (*GetDatacentersResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["datacenterIds"] = args.DatacenterIds
	}
	outputs, err := ctx.Invoke("hcloud:index/getDatacenters:getDatacenters", inputs)
	if err != nil {
		return nil, err
	}
	return &GetDatacentersResult{
		DatacenterIds: outputs["datacenterIds"],
		Descriptions: outputs["descriptions"],
		Names: outputs["names"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getDatacenters.
type GetDatacentersArgs struct {
	DatacenterIds interface{}
}

// A collection of values returned by getDatacenters.
type GetDatacentersResult struct {
	DatacenterIds interface{}
	Descriptions interface{}
	Names interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}