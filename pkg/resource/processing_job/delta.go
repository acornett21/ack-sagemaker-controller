// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package processing_job

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customSetDefaults(a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AppSpecification, b.ko.Spec.AppSpecification) {
		delta.Add("Spec.AppSpecification", a.ko.Spec.AppSpecification, b.ko.Spec.AppSpecification)
	} else if a.ko.Spec.AppSpecification != nil && b.ko.Spec.AppSpecification != nil {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.AppSpecification.ContainerArguments, b.ko.Spec.AppSpecification.ContainerArguments) {
			delta.Add("Spec.AppSpecification.ContainerArguments", a.ko.Spec.AppSpecification.ContainerArguments, b.ko.Spec.AppSpecification.ContainerArguments)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.AppSpecification.ContainerEntrypoint, b.ko.Spec.AppSpecification.ContainerEntrypoint) {
			delta.Add("Spec.AppSpecification.ContainerEntrypoint", a.ko.Spec.AppSpecification.ContainerEntrypoint, b.ko.Spec.AppSpecification.ContainerEntrypoint)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.AppSpecification.ImageURI, b.ko.Spec.AppSpecification.ImageURI) {
			delta.Add("Spec.AppSpecification.ImageURI", a.ko.Spec.AppSpecification.ImageURI, b.ko.Spec.AppSpecification.ImageURI)
		} else if a.ko.Spec.AppSpecification.ImageURI != nil && b.ko.Spec.AppSpecification.ImageURI != nil {
			if *a.ko.Spec.AppSpecification.ImageURI != *b.ko.Spec.AppSpecification.ImageURI {
				delta.Add("Spec.AppSpecification.ImageURI", a.ko.Spec.AppSpecification.ImageURI, b.ko.Spec.AppSpecification.ImageURI)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Environment, b.ko.Spec.Environment) {
		delta.Add("Spec.Environment", a.ko.Spec.Environment, b.ko.Spec.Environment)
	} else if a.ko.Spec.Environment != nil && b.ko.Spec.Environment != nil {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.Environment, b.ko.Spec.Environment) {
			delta.Add("Spec.Environment", a.ko.Spec.Environment, b.ko.Spec.Environment)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig, b.ko.Spec.ExperimentConfig) {
		delta.Add("Spec.ExperimentConfig", a.ko.Spec.ExperimentConfig, b.ko.Spec.ExperimentConfig)
	} else if a.ko.Spec.ExperimentConfig != nil && b.ko.Spec.ExperimentConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig.ExperimentName, b.ko.Spec.ExperimentConfig.ExperimentName) {
			delta.Add("Spec.ExperimentConfig.ExperimentName", a.ko.Spec.ExperimentConfig.ExperimentName, b.ko.Spec.ExperimentConfig.ExperimentName)
		} else if a.ko.Spec.ExperimentConfig.ExperimentName != nil && b.ko.Spec.ExperimentConfig.ExperimentName != nil {
			if *a.ko.Spec.ExperimentConfig.ExperimentName != *b.ko.Spec.ExperimentConfig.ExperimentName {
				delta.Add("Spec.ExperimentConfig.ExperimentName", a.ko.Spec.ExperimentConfig.ExperimentName, b.ko.Spec.ExperimentConfig.ExperimentName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig.TrialComponentDisplayName, b.ko.Spec.ExperimentConfig.TrialComponentDisplayName) {
			delta.Add("Spec.ExperimentConfig.TrialComponentDisplayName", a.ko.Spec.ExperimentConfig.TrialComponentDisplayName, b.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
		} else if a.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil && b.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil {
			if *a.ko.Spec.ExperimentConfig.TrialComponentDisplayName != *b.ko.Spec.ExperimentConfig.TrialComponentDisplayName {
				delta.Add("Spec.ExperimentConfig.TrialComponentDisplayName", a.ko.Spec.ExperimentConfig.TrialComponentDisplayName, b.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig.TrialName, b.ko.Spec.ExperimentConfig.TrialName) {
			delta.Add("Spec.ExperimentConfig.TrialName", a.ko.Spec.ExperimentConfig.TrialName, b.ko.Spec.ExperimentConfig.TrialName)
		} else if a.ko.Spec.ExperimentConfig.TrialName != nil && b.ko.Spec.ExperimentConfig.TrialName != nil {
			if *a.ko.Spec.ExperimentConfig.TrialName != *b.ko.Spec.ExperimentConfig.TrialName {
				delta.Add("Spec.ExperimentConfig.TrialName", a.ko.Spec.ExperimentConfig.TrialName, b.ko.Spec.ExperimentConfig.TrialName)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig, b.ko.Spec.NetworkConfig) {
		delta.Add("Spec.NetworkConfig", a.ko.Spec.NetworkConfig, b.ko.Spec.NetworkConfig)
	} else if a.ko.Spec.NetworkConfig != nil && b.ko.Spec.NetworkConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption, b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption) {
			delta.Add("Spec.NetworkConfig.EnableInterContainerTrafficEncryption", a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption, b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption)
		} else if a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != nil && b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != nil {
			if *a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != *b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption {
				delta.Add("Spec.NetworkConfig.EnableInterContainerTrafficEncryption", a.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption, b.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig.EnableNetworkIsolation, b.ko.Spec.NetworkConfig.EnableNetworkIsolation) {
			delta.Add("Spec.NetworkConfig.EnableNetworkIsolation", a.ko.Spec.NetworkConfig.EnableNetworkIsolation, b.ko.Spec.NetworkConfig.EnableNetworkIsolation)
		} else if a.ko.Spec.NetworkConfig.EnableNetworkIsolation != nil && b.ko.Spec.NetworkConfig.EnableNetworkIsolation != nil {
			if *a.ko.Spec.NetworkConfig.EnableNetworkIsolation != *b.ko.Spec.NetworkConfig.EnableNetworkIsolation {
				delta.Add("Spec.NetworkConfig.EnableNetworkIsolation", a.ko.Spec.NetworkConfig.EnableNetworkIsolation, b.ko.Spec.NetworkConfig.EnableNetworkIsolation)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.NetworkConfig.VPCConfig, b.ko.Spec.NetworkConfig.VPCConfig) {
			delta.Add("Spec.NetworkConfig.VPCConfig", a.ko.Spec.NetworkConfig.VPCConfig, b.ko.Spec.NetworkConfig.VPCConfig)
		} else if a.ko.Spec.NetworkConfig.VPCConfig != nil && b.ko.Spec.NetworkConfig.VPCConfig != nil {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs, b.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs) {
				delta.Add("Spec.NetworkConfig.VPCConfig.SecurityGroupIDs", a.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs, b.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs)
			}
			if !ackcompare.SliceStringPEqual(a.ko.Spec.NetworkConfig.VPCConfig.Subnets, b.ko.Spec.NetworkConfig.VPCConfig.Subnets) {
				delta.Add("Spec.NetworkConfig.VPCConfig.Subnets", a.ko.Spec.NetworkConfig.VPCConfig.Subnets, b.ko.Spec.NetworkConfig.VPCConfig.Subnets)
			}
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ProcessingInputs, b.ko.Spec.ProcessingInputs) {
		delta.Add("Spec.ProcessingInputs", a.ko.Spec.ProcessingInputs, b.ko.Spec.ProcessingInputs)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ProcessingJobName, b.ko.Spec.ProcessingJobName) {
		delta.Add("Spec.ProcessingJobName", a.ko.Spec.ProcessingJobName, b.ko.Spec.ProcessingJobName)
	} else if a.ko.Spec.ProcessingJobName != nil && b.ko.Spec.ProcessingJobName != nil {
		if *a.ko.Spec.ProcessingJobName != *b.ko.Spec.ProcessingJobName {
			delta.Add("Spec.ProcessingJobName", a.ko.Spec.ProcessingJobName, b.ko.Spec.ProcessingJobName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ProcessingOutputConfig, b.ko.Spec.ProcessingOutputConfig) {
		delta.Add("Spec.ProcessingOutputConfig", a.ko.Spec.ProcessingOutputConfig, b.ko.Spec.ProcessingOutputConfig)
	} else if a.ko.Spec.ProcessingOutputConfig != nil && b.ko.Spec.ProcessingOutputConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ProcessingOutputConfig.KMSKeyID, b.ko.Spec.ProcessingOutputConfig.KMSKeyID) {
			delta.Add("Spec.ProcessingOutputConfig.KMSKeyID", a.ko.Spec.ProcessingOutputConfig.KMSKeyID, b.ko.Spec.ProcessingOutputConfig.KMSKeyID)
		} else if a.ko.Spec.ProcessingOutputConfig.KMSKeyID != nil && b.ko.Spec.ProcessingOutputConfig.KMSKeyID != nil {
			if *a.ko.Spec.ProcessingOutputConfig.KMSKeyID != *b.ko.Spec.ProcessingOutputConfig.KMSKeyID {
				delta.Add("Spec.ProcessingOutputConfig.KMSKeyID", a.ko.Spec.ProcessingOutputConfig.KMSKeyID, b.ko.Spec.ProcessingOutputConfig.KMSKeyID)
			}
		}
		if !reflect.DeepEqual(a.ko.Spec.ProcessingOutputConfig.Outputs, b.ko.Spec.ProcessingOutputConfig.Outputs) {
			delta.Add("Spec.ProcessingOutputConfig.Outputs", a.ko.Spec.ProcessingOutputConfig.Outputs, b.ko.Spec.ProcessingOutputConfig.Outputs)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ProcessingResources, b.ko.Spec.ProcessingResources) {
		delta.Add("Spec.ProcessingResources", a.ko.Spec.ProcessingResources, b.ko.Spec.ProcessingResources)
	} else if a.ko.Spec.ProcessingResources != nil && b.ko.Spec.ProcessingResources != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ProcessingResources.ClusterConfig, b.ko.Spec.ProcessingResources.ClusterConfig) {
			delta.Add("Spec.ProcessingResources.ClusterConfig", a.ko.Spec.ProcessingResources.ClusterConfig, b.ko.Spec.ProcessingResources.ClusterConfig)
		} else if a.ko.Spec.ProcessingResources.ClusterConfig != nil && b.ko.Spec.ProcessingResources.ClusterConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount, b.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount) {
				delta.Add("Spec.ProcessingResources.ClusterConfig.InstanceCount", a.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount, b.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount)
			} else if a.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount != nil && b.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount != nil {
				if *a.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount != *b.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount {
					delta.Add("Spec.ProcessingResources.ClusterConfig.InstanceCount", a.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount, b.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ProcessingResources.ClusterConfig.InstanceType, b.ko.Spec.ProcessingResources.ClusterConfig.InstanceType) {
				delta.Add("Spec.ProcessingResources.ClusterConfig.InstanceType", a.ko.Spec.ProcessingResources.ClusterConfig.InstanceType, b.ko.Spec.ProcessingResources.ClusterConfig.InstanceType)
			} else if a.ko.Spec.ProcessingResources.ClusterConfig.InstanceType != nil && b.ko.Spec.ProcessingResources.ClusterConfig.InstanceType != nil {
				if *a.ko.Spec.ProcessingResources.ClusterConfig.InstanceType != *b.ko.Spec.ProcessingResources.ClusterConfig.InstanceType {
					delta.Add("Spec.ProcessingResources.ClusterConfig.InstanceType", a.ko.Spec.ProcessingResources.ClusterConfig.InstanceType, b.ko.Spec.ProcessingResources.ClusterConfig.InstanceType)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID, b.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID) {
				delta.Add("Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID", a.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID, b.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID)
			} else if a.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID != nil && b.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID != nil {
				if *a.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID != *b.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID {
					delta.Add("Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID", a.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID, b.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB, b.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB) {
				delta.Add("Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB", a.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB, b.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB)
			} else if a.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB != nil && b.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB != nil {
				if *a.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB != *b.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB {
					delta.Add("Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB", a.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB, b.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RoleARN, b.ko.Spec.RoleARN) {
		delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
	} else if a.ko.Spec.RoleARN != nil && b.ko.Spec.RoleARN != nil {
		if *a.ko.Spec.RoleARN != *b.ko.Spec.RoleARN {
			delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StoppingCondition, b.ko.Spec.StoppingCondition) {
		delta.Add("Spec.StoppingCondition", a.ko.Spec.StoppingCondition, b.ko.Spec.StoppingCondition)
	} else if a.ko.Spec.StoppingCondition != nil && b.ko.Spec.StoppingCondition != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds) {
			delta.Add("Spec.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
		} else if a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil && b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil {
			if *a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != *b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds {
				delta.Add("Spec.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
			}
		}
	}

	return delta
}
