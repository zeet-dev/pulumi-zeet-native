// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zeet.Model.Outputs
{

    [OutputType]
    public sealed class CreateAppGithubInput
    {
        public readonly string ProductionBranch;
        public readonly string Url;

        [OutputConstructor]
        private CreateAppGithubInput(
            string productionBranch,

            string url)
        {
            ProductionBranch = productionBranch;
            Url = url;
        }
    }
}
