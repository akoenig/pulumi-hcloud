# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetSshKeyResult:
    """
    A collection of values returned by getSshKey.
    """
    def __init__(__self__, fingerprint=None, id=None, labels=None, name=None, public_key=None, selector=None, with_selector=None):
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        __self__.fingerprint = fingerprint
        if id and not isinstance(id, float):
            raise TypeError("Expected argument 'id' to be a float")
        __self__.id = id
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if public_key and not isinstance(public_key, str):
            raise TypeError("Expected argument 'public_key' to be a str")
        __self__.public_key = public_key
        if selector and not isinstance(selector, str):
            raise TypeError("Expected argument 'selector' to be a str")
        __self__.selector = selector
        if with_selector and not isinstance(with_selector, str):
            raise TypeError("Expected argument 'with_selector' to be a str")
        __self__.with_selector = with_selector

async def get_ssh_key(fingerprint=None,id=None,name=None,selector=None,with_selector=None,opts=None):
    """
    Provides details about a Hetzner Cloud SSH Key.
    This resource is useful if you want to use a non-terraform managed SSH Key.
    """
    __args__ = dict()

    __args__['fingerprint'] = fingerprint
    __args__['id'] = id
    __args__['name'] = name
    __args__['selector'] = selector
    __args__['withSelector'] = with_selector
    __ret__ = await pulumi.runtime.invoke('hcloud:index/getSshKey:getSshKey', __args__, opts=opts)

    return GetSshKeyResult(
        fingerprint=__ret__.get('fingerprint'),
        id=__ret__.get('id'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        public_key=__ret__.get('publicKey'),
        selector=__ret__.get('selector'),
        with_selector=__ret__.get('withSelector'))
