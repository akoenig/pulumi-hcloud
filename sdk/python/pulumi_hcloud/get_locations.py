# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetLocationsResult:
    """
    A collection of values returned by getLocations.
    """
    def __init__(__self__, descriptions=None, location_ids=None, names=None, id=None):
        if descriptions and not isinstance(descriptions, list):
            raise TypeError("Expected argument 'descriptions' to be a list")
        __self__.descriptions = descriptions
        if location_ids and not isinstance(location_ids, list):
            raise TypeError("Expected argument 'location_ids' to be a list")
        __self__.location_ids = location_ids
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_locations(location_ids=None,opts=None):
    """
    Provides a list of available Hetzner Cloud Locations.
    This resource may be useful to create highly available infrastructure, distributed across several locations.
    """
    __args__ = dict()

    __args__['locationIds'] = location_ids
    __ret__ = await pulumi.runtime.invoke('hcloud:index/getLocations:getLocations', __args__, opts=opts)

    return GetLocationsResult(
        descriptions=__ret__.get('descriptions'),
        location_ids=__ret__.get('locationIds'),
        names=__ret__.get('names'),
        id=__ret__.get('id'))
