// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace model {
    export interface CreateAppBuildInputArgs {
        dockerfilePath?: pulumi.Input<string>;
        type: pulumi.Input<string>;
    }

    export interface CreateAppDeployInputArgs {
        clusterId?: pulumi.Input<string>;
        deployTarget: pulumi.Input<string>;
    }

    export interface CreateAppEnvironmentVariableInputArgs {
        name: pulumi.Input<string>;
        sealed?: pulumi.Input<boolean>;
        value: pulumi.Input<string>;
    }

    export interface CreateAppGithubInputArgs {
        productionBranch: pulumi.Input<string>;
        url: pulumi.Input<string>;
    }

    export interface CreateAppResourcesInputArgs {
        cpu: pulumi.Input<string>;
        memory: pulumi.Input<string>;
    }
}

export namespace time {
}
