# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class FloatingIp(pulumi.CustomResource):
    description: pulumi.Output[str]
    home_location: pulumi.Output[str]
    ip_address: pulumi.Output[str]
    ip_network: pulumi.Output[str]
    labels: pulumi.Output[dict]
    server_id: pulumi.Output[float]
    type: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, description=None, home_location=None, labels=None, server_id=None, type=None, __name__=None, __opts__=None):
        """
        Provides a Hetzner Cloud Floating IP to represent a publicly-accessible static IP address that can be mapped to one of your servers.
        
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

        __props__['description'] = description

        __props__['home_location'] = home_location

        __props__['labels'] = labels

        __props__['server_id'] = server_id

        if type is None:
            raise TypeError("Missing required property 'type'")
        __props__['type'] = type

        __props__['ip_address'] = None
        __props__['ip_network'] = None

        super(FloatingIp, __self__).__init__(
            'hcloud:index/floatingIp:FloatingIp',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

