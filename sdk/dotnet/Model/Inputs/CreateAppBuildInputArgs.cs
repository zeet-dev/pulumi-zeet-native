// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ZeetNative.Model.Inputs
{

    public sealed class CreateAppBuildInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("dockerfilePath")]
        public Input<string>? DockerfilePath { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CreateAppBuildInputArgs()
        {
        }
        public static new CreateAppBuildInputArgs Empty => new CreateAppBuildInputArgs();
    }
}
