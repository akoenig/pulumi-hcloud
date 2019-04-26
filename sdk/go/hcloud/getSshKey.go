// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hcloud

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides details about a Hetzner Cloud SSH Key.
// This resource is useful if you want to use a non-terraform managed SSH Key.
func LookupSshKey(ctx *pulumi.Context, args *GetSshKeyArgs) (*GetSshKeyResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["fingerprint"] = args.Fingerprint
		inputs["id"] = args.Id
		inputs["name"] = args.Name
		inputs["selector"] = args.Selector
		inputs["withSelector"] = args.WithSelector
	}
	outputs, err := ctx.Invoke("hcloud:index/getSshKey:getSshKey", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSshKeyResult{
		Fingerprint: outputs["fingerprint"],
		Id: outputs["id"],
		Labels: outputs["labels"],
		Name: outputs["name"],
		PublicKey: outputs["publicKey"],
		Selector: outputs["selector"],
		WithSelector: outputs["withSelector"],
	}, nil
}

// A collection of arguments for invoking getSshKey.
type GetSshKeyArgs struct {
	Fingerprint interface{}
	Id interface{}
	Name interface{}
	Selector interface{}
	WithSelector interface{}
}

// A collection of values returned by getSshKey.
type GetSshKeyResult struct {
	Fingerprint interface{}
	Id interface{}
	Labels interface{}
	Name interface{}
	PublicKey interface{}
	Selector interface{}
	WithSelector interface{}
}
