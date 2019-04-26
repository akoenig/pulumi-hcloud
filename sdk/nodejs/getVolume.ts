// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getVolume(args?: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    args = args || {};
    return pulumi.runtime.invoke("hcloud:index/getVolume:getVolume", {
        "id": args.id,
        "location": args.location,
        "name": args.name,
        "selector": args.selector,
        "server": args.server,
        "withSelector": args.withSelector,
        "withStatuses": args.withStatuses,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeArgs {
    readonly id?: number;
    readonly location?: string;
    readonly name?: string;
    readonly selector?: string;
    readonly server?: string;
    readonly withSelector?: string;
    readonly withStatuses?: string[];
}

/**
 * A collection of values returned by getVolume.
 */
export interface GetVolumeResult {
    readonly id?: number;
    readonly labels: {[key: string]: any};
    readonly linuxDevice: string;
    readonly location?: string;
    readonly name: string;
    readonly selector?: string;
    readonly server?: string;
    readonly size: number;
    readonly withSelector?: string;
    readonly withStatuses?: string[];
}
