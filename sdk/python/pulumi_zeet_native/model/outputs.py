# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CreateAppBuildInput',
    'CreateAppDeployInput',
    'CreateAppEnvironmentVariableInput',
    'CreateAppGithubInput',
    'CreateAppResourcesInput',
]

@pulumi.output_type
class CreateAppBuildInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dockerfilePath":
            suggest = "dockerfile_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CreateAppBuildInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CreateAppBuildInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CreateAppBuildInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 dockerfile_path: Optional[str] = None):
        pulumi.set(__self__, "type", type)
        if dockerfile_path is not None:
            pulumi.set(__self__, "dockerfile_path", dockerfile_path)

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="dockerfilePath")
    def dockerfile_path(self) -> Optional[str]:
        return pulumi.get(self, "dockerfile_path")


@pulumi.output_type
class CreateAppDeployInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deployTarget":
            suggest = "deploy_target"
        elif key == "clusterId":
            suggest = "cluster_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CreateAppDeployInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CreateAppDeployInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CreateAppDeployInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 deploy_target: str,
                 cluster_id: Optional[str] = None):
        pulumi.set(__self__, "deploy_target", deploy_target)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)

    @property
    @pulumi.getter(name="deployTarget")
    def deploy_target(self) -> str:
        return pulumi.get(self, "deploy_target")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        return pulumi.get(self, "cluster_id")


@pulumi.output_type
class CreateAppEnvironmentVariableInput(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str,
                 sealed: Optional[bool] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if sealed is not None:
            pulumi.set(__self__, "sealed", sealed)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def sealed(self) -> Optional[bool]:
        return pulumi.get(self, "sealed")


@pulumi.output_type
class CreateAppGithubInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "productionBranch":
            suggest = "production_branch"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CreateAppGithubInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CreateAppGithubInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CreateAppGithubInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 production_branch: str,
                 url: str):
        pulumi.set(__self__, "production_branch", production_branch)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="productionBranch")
    def production_branch(self) -> str:
        return pulumi.get(self, "production_branch")

    @property
    @pulumi.getter
    def url(self) -> str:
        return pulumi.get(self, "url")


@pulumi.output_type
class CreateAppResourcesInput(dict):
    def __init__(__self__, *,
                 cpu: str,
                 memory: str):
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "memory", memory)

    @property
    @pulumi.getter
    def cpu(self) -> str:
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter
    def memory(self) -> str:
        return pulumi.get(self, "memory")


