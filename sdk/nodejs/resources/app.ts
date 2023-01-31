// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zeet-native:resources:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    public /*out*/ readonly appId!: pulumi.Output<string>;
    public readonly build!: pulumi.Output<outputs.model.CreateAppBuildInput | undefined>;
    public readonly deploy!: pulumi.Output<outputs.model.CreateAppDeployInput>;
    public readonly docker!: pulumi.Output<outputs.model.CreateAppDockerInput | undefined>;
    public readonly enabled!: pulumi.Output<boolean>;
    public readonly environmentId!: pulumi.Output<string>;
    public readonly environmentVariables!: pulumi.Output<outputs.model.CreateAppEnvironmentVariableInput[] | undefined>;
    public readonly github!: pulumi.Output<outputs.model.CreateAppGithubInput | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly projectId!: pulumi.Output<string>;
    public readonly resources!: pulumi.Output<outputs.model.CreateAppResourcesInput>;
    public /*out*/ readonly updatedAt!: pulumi.Output<outputs.time.Time>;
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deploy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploy'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.resources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resources'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["build"] = args ? args.build : undefined;
            resourceInputs["deploy"] = args ? args.deploy : undefined;
            resourceInputs["docker"] = args ? args.docker : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["github"] = args ? args.github : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["build"] = undefined /*out*/;
            resourceInputs["deploy"] = undefined /*out*/;
            resourceInputs["docker"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["environmentVariables"] = undefined /*out*/;
            resourceInputs["github"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(App.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    build?: pulumi.Input<inputs.model.CreateAppBuildInputArgs>;
    deploy: pulumi.Input<inputs.model.CreateAppDeployInputArgs>;
    docker?: pulumi.Input<inputs.model.CreateAppDockerInputArgs>;
    enabled: pulumi.Input<boolean>;
    environmentId: pulumi.Input<string>;
    environmentVariables?: pulumi.Input<pulumi.Input<inputs.model.CreateAppEnvironmentVariableInputArgs>[]>;
    github?: pulumi.Input<inputs.model.CreateAppGithubInputArgs>;
    name: pulumi.Input<string>;
    projectId: pulumi.Input<string>;
    resources: pulumi.Input<inputs.model.CreateAppResourcesInputArgs>;
    userId: pulumi.Input<string>;
}
