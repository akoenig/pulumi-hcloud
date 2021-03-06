# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetDatacentersResult:
    """
    A collection of values returned by getDatacenters.
    """
    def __init__(__self__, datacenter_ids=None, descriptions=None, names=None, id=None):
        if datacenter_ids and not isinstance(datacenter_ids, list):
            raise TypeError("Expected argument 'datacenter_ids' to be a list")
        __self__.datacenter_ids = datacenter_ids
        if descriptions and not isinstance(descriptions, list):
            raise TypeError("Expected argument 'descriptions' to be a list")
        __self__.descriptions = descriptions
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_datacenters(datacenter_ids=None,opts=None):
    """
    Provides a list of available Hetzner Cloud Datacenters.
    This resource may be useful to create highly available infrastructure, distributed across several datacenters.
    """
    __args__ = dict()

    __args__['datacenterIds'] = datacenter_ids
    __ret__ = await pulumi.runtime.invoke('hcloud:index/getDatacenters:getDatacenters', __args__, opts=opts)

    return GetDatacentersResult(
        datacenter_ids=__ret__.get('datacenterIds'),
        descriptions=__ret__.get('descriptions'),
        names=__ret__.get('names'),
        id=__ret__.get('id'))
