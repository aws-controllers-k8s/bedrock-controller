package tags

import (
	"context"
	"maps"
	"slices"

	svcapitypes "github.com/aws-controllers-k8s/bedrock-controller/apis/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
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

// SyncResourceTags uses TagResource and UntagResource API Calls to add, remove
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

	from := toACKTags(latestTags)
	to := toACKTags(desiredTags)
	added, _, removed := ackcompare.GetTagsDifference(from, to)

	if len(added) == 0 && len(removed) == 0 {
		return nil
	}

	if len(removed) > 0 {
		_, err = client.UntagResource(
			ctx,
			&svcsdk.UntagResourceInput{
				ResourceARN: aws.String(resourceARN),
				TagKeys:     slices.Sorted(maps.Keys(removed)),
			},
		)
		mr.RecordAPICall("UPDATE", "UntagResource", err)
		if err != nil {
			return err
		}
	}

	if len(added) > 0 {
		addedOrUpdated := make([]svcsdktypes.Tag, 0, len(added))
		for _, k := range slices.Sorted(maps.Keys(added)) {
			addedOrUpdated = append(addedOrUpdated, svcsdktypes.Tag{
				Key:   aws.String(k),
				Value: aws.String(added[k]),
			})
		}
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

// toACKTags converts a Bedrock tag list into the ACK runtime tags map.
// Nil tags and tags without a key are skipped.
func toACKTags(tags []*svcapitypes.Tag) acktags.Tags {
	m := make(acktags.Tags)
	for _, t := range tags {
		if t == nil || t.Key == nil {
			continue
		}
		v := ""
		if t.Value != nil {
			v = *t.Value
		}
		m[*t.Key] = v
	}
	return m
}
