// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hcloud

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Hetzner Cloud Volume attachment to attach a Volume to a Hetzner Cloud Server. Deleting a Volume Attachment will detach the Volume from the Server.
type VolumeAttachment struct {
	s *pulumi.ResourceState
}

// NewVolumeAttachment registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachment(ctx *pulumi.Context,
	name string, args *VolumeAttachmentArgs, opts ...pulumi.ResourceOpt) (*VolumeAttachment, error) {
	if args == nil || args.ServerId == nil {
		return nil, errors.New("missing required argument 'ServerId'")
	}
	if args == nil || args.VolumeId == nil {
		return nil, errors.New("missing required argument 'VolumeId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["automount"] = nil
		inputs["serverId"] = nil
		inputs["volumeId"] = nil
	} else {
		inputs["automount"] = args.Automount
		inputs["serverId"] = args.ServerId
		inputs["volumeId"] = args.VolumeId
	}
	s, err := ctx.RegisterResource("hcloud:index/volumeAttachment:VolumeAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VolumeAttachment{s: s}, nil
}

// GetVolumeAttachment gets an existing VolumeAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VolumeAttachmentState, opts ...pulumi.ResourceOpt) (*VolumeAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["automount"] = state.Automount
		inputs["serverId"] = state.ServerId
		inputs["volumeId"] = state.VolumeId
	}
	s, err := ctx.ReadResource("hcloud:index/volumeAttachment:VolumeAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VolumeAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VolumeAttachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VolumeAttachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *VolumeAttachment) Automount() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["automount"])
}

func (r *VolumeAttachment) ServerId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["serverId"])
}

func (r *VolumeAttachment) VolumeId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["volumeId"])
}

// Input properties used for looking up and filtering VolumeAttachment resources.
type VolumeAttachmentState struct {
	Automount interface{}
	ServerId interface{}
	VolumeId interface{}
}

// The set of arguments for constructing a VolumeAttachment resource.
type VolumeAttachmentArgs struct {
	Automount interface{}
	ServerId interface{}
	VolumeId interface{}
}
