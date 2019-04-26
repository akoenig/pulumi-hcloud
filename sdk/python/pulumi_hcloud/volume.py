# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Volume(pulumi.CustomResource):
    automount: pulumi.Output[bool]
    format: pulumi.Output[str]
    labels: pulumi.Output[dict]
    linux_device: pulumi.Output[str]
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    server_id: pulumi.Output[float]
    size: pulumi.Output[float]
    def __init__(__self__, resource_name, opts=None, automount=None, format=None, labels=None, location=None, name=None, server_id=None, size=None, __name__=None, __opts__=None):
        """
        Provides a Hetzner Cloud volume resource to manage volumes.
        
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

        __props__['automount'] = automount

        __props__['format'] = format

        __props__['labels'] = labels

        __props__['location'] = location

        __props__['name'] = name

        __props__['server_id'] = server_id

        if size is None:
            raise TypeError("Missing required property 'size'")
        __props__['size'] = size

        __props__['linux_device'] = None

        super(Volume, __self__).__init__(
            'hcloud:index/volume:Volume',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
