# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Server(pulumi.CustomResource):
    backup_window: pulumi.Output[str]
    backups: pulumi.Output[bool]
    datacenter: pulumi.Output[str]
    image: pulumi.Output[str]
    ipv4_address: pulumi.Output[str]
    ipv6_address: pulumi.Output[str]
    ipv6_network: pulumi.Output[str]
    iso: pulumi.Output[str]
    keep_disk: pulumi.Output[bool]
    labels: pulumi.Output[dict]
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    rescue: pulumi.Output[str]
    server_type: pulumi.Output[str]
    ssh_keys: pulumi.Output[list]
    status: pulumi.Output[str]
    user_data: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, backups=None, datacenter=None, image=None, iso=None, keep_disk=None, labels=None, location=None, name=None, rescue=None, server_type=None, ssh_keys=None, user_data=None, __name__=None, __opts__=None):
        """
        Provides an Hetzner Cloud server resource. This can be used to create, modify, and delete servers. Servers also support [provisioning](https://www.terraform.io/docs/provisioners/index.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['backups'] = backups

        __props__['datacenter'] = datacenter

        if image is None:
            raise TypeError("Missing required property 'image'")
        __props__['image'] = image

        __props__['iso'] = iso

        __props__['keep_disk'] = keep_disk

        __props__['labels'] = labels

        __props__['location'] = location

        __props__['name'] = name

        __props__['rescue'] = rescue

        if server_type is None:
            raise TypeError("Missing required property 'server_type'")
        __props__['server_type'] = server_type

        __props__['ssh_keys'] = ssh_keys

        __props__['user_data'] = user_data

        __props__['backup_window'] = None
        __props__['ipv4_address'] = None
        __props__['ipv6_address'] = None
        __props__['ipv6_network'] = None
        __props__['status'] = None

        super(Server, __self__).__init__(
            'hcloud:index/server:Server',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
