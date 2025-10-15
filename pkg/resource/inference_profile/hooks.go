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

package inference_profile

import (
	"context"

	svcapitypes "github.com/aws-controllers-k8s/bedrock-controller/apis/v1alpha1"
	"github.com/aws-controllers-k8s/bedrock-controller/pkg/tags"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// getTags retrieves the resource's associated tags.
func (rm *resourceManager) getTags(
	ctx context.Context,
	resourceARN string,
) ([]*svcapitypes.Tag, error) {
	return tags.GetResourceTags(ctx, rm.sdkapi, rm.metrics, resourceARN)
}

// syncTags keeps the resource's tags in sync.
func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	return tags.SyncResourceTags(
		ctx,
		rm.sdkapi,
		rm.metrics,
		string(*latest.ko.Status.ACKResourceMetadata.ARN),
		desired.ko.Spec.Tags,
		latest.ko.Spec.Tags,
	)
}

// customCheckRequired returns true if the resource's ARN is not available,
// which is required for the GetInferenceProfile API call
func (rm *resourceManager) customCheckRequired(r *resource) bool {
	return r.ko.Status.ACKResourceMetadata == nil || r.ko.Status.ACKResourceMetadata.ARN == nil
}

// custom_update handles updates for inference profiles. Since inference profiles
// are immutable resources, only tag updates are supported.
func (rm *resourceManager) custom_update(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	if delta.DifferentAt("Spec.Tags") {
		if err := rm.syncTags(ctx, desired, latest); err != nil {
			return nil, err
		}
	}

	return desired, nil
}
