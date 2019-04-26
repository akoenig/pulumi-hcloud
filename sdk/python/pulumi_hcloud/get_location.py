# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetLocationResult:
    """
    A collection of values returned by getLocation.
    """
    def __init__(__self__, city=None, country=None, description=None, id=None, latitude=None, longitude=None, name=None):
        if city and not isinstance(city, str):
            raise TypeError("Expected argument 'city' to be a str")
        __self__.city = city
        if country and not isinstance(country, str):
            raise TypeError("Expected argument 'country' to be a str")
        __self__.country = country
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if id and not isinstance(id, float):
            raise TypeError("Expected argument 'id' to be a float")
        __self__.id = id
        if latitude and not isinstance(latitude, float):
            raise TypeError("Expected argument 'latitude' to be a float")
        __self__.latitude = latitude
        if longitude and not isinstance(longitude, float):
            raise TypeError("Expected argument 'longitude' to be a float")
        __self__.longitude = longitude
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name

async def get_location(id=None,name=None,opts=None):
    """
    Provides details about a specific Hetzner Cloud Location.
    Use this resource to get detailed information about specific location.
    """
    __args__ = dict()

    __args__['id'] = id
    __args__['name'] = name
    __ret__ = await pulumi.runtime.invoke('hcloud:index/getLocation:getLocation', __args__, opts=opts)

    return GetLocationResult(
        city=__ret__.get('city'),
        country=__ret__.get('country'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        latitude=__ret__.get('latitude'),
        longitude=__ret__.get('longitude'),
        name=__ret__.get('name'))
