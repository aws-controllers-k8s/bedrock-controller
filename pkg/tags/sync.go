package tags

import (
	"context"

	svcapitypes "github.com/aws-controllers-k8s/bedrock-controller/apis/v1alpha1"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/bedrock"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/bedrock/types"
)

type metricsRecorder interface {
	RecordAPICall(opType string, opID string, err error)
}

type tagsClient interface {
	TagResource(context.Context, *svcsdk.TagResourceInput, ...func(*svcsdk.Options)) (*svcsdk.TagResourceOutput, error)
	ListTagsForResource(context.Context, *svcsdk.ListTagsForResourceInput, ...func(*svcsdk.Options)) (*svcsdk.ListTagsForResourceOutput, error)
	UntagResource(context.Context, *svcsdk.UntagResourceInput, ...func(*svcsdk.Options)) (*svcsdk.UntagResourceOutput, error)
}

// GetResourceTags retrieves a resource list of tags.
func GetResourceTags(
	ctx context.Context,
	client tagsClient,
	mr metricsRecorder,
	resourceARN string,
) ([]*svcapitypes.Tag, error) {
	listTagsForResourceResponse, err := client.ListTagsForResource(
		ctx,
		&svcsdk.ListTagsForResourceInput{
			ResourceARN: &resourceARN,
		},
	)
	mr.RecordAPICall("GET", "ListTagsForResource", err)
	if err != nil {
		return nil, err
	}

	tags := make([]*svcapitypes.Tag, 0)
	for _, tag := range listTagsForResourceResponse.Tags {
		tags = append(tags, &svcapitypes.Tag{
			Key:   tag.Key,
			Value: tag.Value,
		})
	}

	return tags, nil
}

// SyncResourceTags uses TagDeliveryStream and UntagDeliveryStream API Calls to add, remove
// and update resource tags.
func SyncResourceTags(
	ctx context.Context,
	client tagsClient,
	mr metricsRecorder,
	resourceARN string,
	desiredTags []*svcapitypes.Tag,
	latestTags []*svcapitypes.Tag,
) error {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("common.SyncResourceTags")
	defer func() {
		exit(err)
	}()

	addedOrUpdated, removed := computeTagsDelta(desiredTags, latestTags)

	if len(removed) > 0 {
		_, err = client.UntagResource(
			ctx,
			&svcsdk.UntagResourceInput{
				ResourceARN: aws.String(resourceARN),
				TagKeys:     removed,
			},
		)
		mr.RecordAPICall("UPDATE", "UntagResource", err)
		if err != nil {
			return err
		}
	}

	if len(addedOrUpdated) > 0 {
		_, err = client.TagResource(
			ctx,
			&svcsdk.TagResourceInput{
				ResourceARN: aws.String(resourceARN),
				Tags:        addedOrUpdated,
			},
		)
		mr.RecordAPICall("UPDATE", "TagResource", err)
		if err != nil {
			return err
		}
	}
	return nil
}

// computeTagsDelta compares two Tag arrays and return two different list
// containing the addedOrupdated and removed tags. The removed tags array
// only contains the tags Keys.
func computeTagsDelta(
	a []*svcapitypes.Tag,
	b []*svcapitypes.Tag,
) (addedOrUpdated []svcsdktypes.Tag, removed []string) {

	// Find the keys in the Spec have either been added or updated.
	addedOrUpdated = make([]svcsdktypes.Tag, 0)
	for _, aTag := range a {
		found := false
		for _, bTag := range b {
			if *aTag.Key == *bTag.Key {
				if *aTag.Value == *bTag.Value {
					found = true
				}

				break
			}
		}
		if !found {
			addedOrUpdated = append(addedOrUpdated, svcsdktypes.Tag{
				Key:   aTag.Key,
				Value: aTag.Value,
			})
		}
	}

	for _, bTag := range b {
		found := false
		for _, aTag := range a {
			if *aTag.Key == *bTag.Key {
				found = true
				break
			}
		}

		if !found {
			removed = append(removed, *bTag.Key)
		}
	}

	return addedOrUpdated, removed
}

// equalTags returns true if two Tag arrays are equal regardless of the order
// of their elements.
func EqualTags(
	a []*svcapitypes.Tag,
	b []*svcapitypes.Tag,
) bool {
	addedOrUpdated, removed := computeTagsDelta(a, b)
	return len(addedOrUpdated) == 0 && len(removed) == 0
}
