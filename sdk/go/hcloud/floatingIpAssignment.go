// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hcloud

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Hetzner Cloud Floating IP Assignment to assign a Floating IP to a Hetzner Cloud Server. Deleting a Floating IP Assignment will unassign the Floating IP from the Server.
type FloatingIpAssignment struct {
	s *pulumi.ResourceState
}

// NewFloatingIpAssignment registers a new resource with the given unique name, arguments, and options.
func NewFloatingIpAssignment(ctx *pulumi.Context,
	name string, args *FloatingIpAssignmentArgs, opts ...pulumi.ResourceOpt) (*FloatingIpAssignment, error) {
	if args == nil || args.FloatingIpId == nil {
		return nil, errors.New("missing required argument 'FloatingIpId'")
	}
	if args == nil || args.ServerId == nil {
		return nil, errors.New("missing required argument 'ServerId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["floatingIpId"] = nil
		inputs["serverId"] = nil
	} else {
		inputs["floatingIpId"] = args.FloatingIpId
		inputs["serverId"] = args.ServerId
	}
	s, err := ctx.RegisterResource("hcloud:index/floatingIpAssignment:FloatingIpAssignment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FloatingIpAssignment{s: s}, nil
}

// GetFloatingIpAssignment gets an existing FloatingIpAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIpAssignment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FloatingIpAssignmentState, opts ...pulumi.ResourceOpt) (*FloatingIpAssignment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["floatingIpId"] = state.FloatingIpId
		inputs["serverId"] = state.ServerId
	}
	s, err := ctx.ReadResource("hcloud:index/floatingIpAssignment:FloatingIpAssignment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FloatingIpAssignment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FloatingIpAssignment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FloatingIpAssignment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *FloatingIpAssignment) FloatingIpId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["floatingIpId"])
}

func (r *FloatingIpAssignment) ServerId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["serverId"])
}

// Input properties used for looking up and filtering FloatingIpAssignment resources.
type FloatingIpAssignmentState struct {
	FloatingIpId interface{}
	ServerId interface{}
}

// The set of arguments for constructing a FloatingIpAssignment resource.
type FloatingIpAssignmentArgs struct {
	FloatingIpId interface{}
	ServerId interface{}
}
