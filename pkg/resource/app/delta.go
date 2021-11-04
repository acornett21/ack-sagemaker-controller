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

package app

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

	if ackcompare.HasNilDifference(a.ko.Spec.AppName, b.ko.Spec.AppName) {
		delta.Add("Spec.AppName", a.ko.Spec.AppName, b.ko.Spec.AppName)
	} else if a.ko.Spec.AppName != nil && b.ko.Spec.AppName != nil {
		if *a.ko.Spec.AppName != *b.ko.Spec.AppName {
			delta.Add("Spec.AppName", a.ko.Spec.AppName, b.ko.Spec.AppName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AppType, b.ko.Spec.AppType) {
		delta.Add("Spec.AppType", a.ko.Spec.AppType, b.ko.Spec.AppType)
	} else if a.ko.Spec.AppType != nil && b.ko.Spec.AppType != nil {
		if *a.ko.Spec.AppType != *b.ko.Spec.AppType {
			delta.Add("Spec.AppType", a.ko.Spec.AppType, b.ko.Spec.AppType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DomainID, b.ko.Spec.DomainID) {
		delta.Add("Spec.DomainID", a.ko.Spec.DomainID, b.ko.Spec.DomainID)
	} else if a.ko.Spec.DomainID != nil && b.ko.Spec.DomainID != nil {
		if *a.ko.Spec.DomainID != *b.ko.Spec.DomainID {
			delta.Add("Spec.DomainID", a.ko.Spec.DomainID, b.ko.Spec.DomainID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResourceSpec, b.ko.Spec.ResourceSpec) {
		delta.Add("Spec.ResourceSpec", a.ko.Spec.ResourceSpec, b.ko.Spec.ResourceSpec)
	} else if a.ko.Spec.ResourceSpec != nil && b.ko.Spec.ResourceSpec != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceSpec.InstanceType, b.ko.Spec.ResourceSpec.InstanceType) {
			delta.Add("Spec.ResourceSpec.InstanceType", a.ko.Spec.ResourceSpec.InstanceType, b.ko.Spec.ResourceSpec.InstanceType)
		} else if a.ko.Spec.ResourceSpec.InstanceType != nil && b.ko.Spec.ResourceSpec.InstanceType != nil {
			if *a.ko.Spec.ResourceSpec.InstanceType != *b.ko.Spec.ResourceSpec.InstanceType {
				delta.Add("Spec.ResourceSpec.InstanceType", a.ko.Spec.ResourceSpec.InstanceType, b.ko.Spec.ResourceSpec.InstanceType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceSpec.SageMakerImageARN, b.ko.Spec.ResourceSpec.SageMakerImageARN) {
			delta.Add("Spec.ResourceSpec.SageMakerImageARN", a.ko.Spec.ResourceSpec.SageMakerImageARN, b.ko.Spec.ResourceSpec.SageMakerImageARN)
		} else if a.ko.Spec.ResourceSpec.SageMakerImageARN != nil && b.ko.Spec.ResourceSpec.SageMakerImageARN != nil {
			if *a.ko.Spec.ResourceSpec.SageMakerImageARN != *b.ko.Spec.ResourceSpec.SageMakerImageARN {
				delta.Add("Spec.ResourceSpec.SageMakerImageARN", a.ko.Spec.ResourceSpec.SageMakerImageARN, b.ko.Spec.ResourceSpec.SageMakerImageARN)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceSpec.SageMakerImageVersionARN, b.ko.Spec.ResourceSpec.SageMakerImageVersionARN) {
			delta.Add("Spec.ResourceSpec.SageMakerImageVersionARN", a.ko.Spec.ResourceSpec.SageMakerImageVersionARN, b.ko.Spec.ResourceSpec.SageMakerImageVersionARN)
		} else if a.ko.Spec.ResourceSpec.SageMakerImageVersionARN != nil && b.ko.Spec.ResourceSpec.SageMakerImageVersionARN != nil {
			if *a.ko.Spec.ResourceSpec.SageMakerImageVersionARN != *b.ko.Spec.ResourceSpec.SageMakerImageVersionARN {
				delta.Add("Spec.ResourceSpec.SageMakerImageVersionARN", a.ko.Spec.ResourceSpec.SageMakerImageVersionARN, b.ko.Spec.ResourceSpec.SageMakerImageVersionARN)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.UserProfileName, b.ko.Spec.UserProfileName) {
		delta.Add("Spec.UserProfileName", a.ko.Spec.UserProfileName, b.ko.Spec.UserProfileName)
	} else if a.ko.Spec.UserProfileName != nil && b.ko.Spec.UserProfileName != nil {
		if *a.ko.Spec.UserProfileName != *b.ko.Spec.UserProfileName {
			delta.Add("Spec.UserProfileName", a.ko.Spec.UserProfileName, b.ko.Spec.UserProfileName)
		}
	}

	return delta
}
