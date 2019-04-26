// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hcloud

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides details about a specific Hetzner Cloud Location.
// Use this resource to get detailed information about specific location.
func LookupLocation(ctx *pulumi.Context, args *GetLocationArgs) (*GetLocationResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["id"] = args.Id
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("hcloud:index/getLocation:getLocation", inputs)
	if err != nil {
		return nil, err
	}
	return &GetLocationResult{
		City: outputs["city"],
		Country: outputs["country"],
		Description: outputs["description"],
		Id: outputs["id"],
		Latitude: outputs["latitude"],
		Longitude: outputs["longitude"],
		Name: outputs["name"],
	}, nil
}

// A collection of arguments for invoking getLocation.
type GetLocationArgs struct {
	Id interface{}
	Name interface{}
}

// A collection of values returned by getLocation.
type GetLocationResult struct {
	City interface{}
	Country interface{}
	Description interface{}
	Id interface{}
	Latitude interface{}
	Longitude interface{}
	Name interface{}
}
