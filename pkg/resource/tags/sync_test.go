package tags

import (
	"context"
	"errors"
	"testing"

	svcapitypes "github.com/aws-controllers-k8s/bedrock-controller/apis/v1alpha1"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/bedrock"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/bedrock/types"
)

type mockTagsClient struct {
	listTagsResponse *svcsdk.ListTagsForResourceOutput
	listTagsError    error
	tagResourceError error
	untagError       error
}

func (m *mockTagsClient) TagResource(ctx context.Context, input *svcsdk.TagResourceInput, opts ...func(*svcsdk.Options)) (*svcsdk.TagResourceOutput, error) {
	return &svcsdk.TagResourceOutput{}, m.tagResourceError
}

func (m *mockTagsClient) ListTagsForResource(ctx context.Context, input *svcsdk.ListTagsForResourceInput, opts ...func(*svcsdk.Options)) (*svcsdk.ListTagsForResourceOutput, error) {
	return m.listTagsResponse, m.listTagsError
}

func (m *mockTagsClient) UntagResource(ctx context.Context, input *svcsdk.UntagResourceInput, opts ...func(*svcsdk.Options)) (*svcsdk.UntagResourceOutput, error) {
	return &svcsdk.UntagResourceOutput{}, m.untagError
}

type mockMetrics struct{}

func (m *mockMetrics) RecordAPICall(opType string, opID string, err error) {}

func TestGetResourceTags(t *testing.T) {
	tests := []struct {
		name     string
		response *svcsdk.ListTagsForResourceOutput
		err      error
		expected []*svcapitypes.Tag
		wantErr  bool
	}{
		{
			name: "success",
			response: &svcsdk.ListTagsForResourceOutput{
				Tags: []svcsdktypes.Tag{
					{Key: aws.String("key1"), Value: aws.String("value1")},
					{Key: aws.String("key2"), Value: aws.String("value2")},
				},
			},
			expected: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
				{Key: aws.String("key2"), Value: aws.String("value2")},
			},
		},
		{
			name:    "error",
			err:     errors.New("api error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockTagsClient{
				listTagsResponse: tt.response,
				listTagsError:    tt.err,
			}
			metrics := &mockMetrics{}

			result, err := GetResourceTags(context.Background(), client, metrics, "test-arn")

			if tt.wantErr && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d tags, got %d", len(tt.expected), len(result))
			}
		})
	}
}

func TestComputeTagsDelta(t *testing.T) {
	tests := []struct {
		name            string
		desired         []*svcapitypes.Tag
		latest          []*svcapitypes.Tag
		expectedAdded   int
		expectedRemoved int
	}{
		{
			name: "add new tag",
			desired: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			latest:          []*svcapitypes.Tag{},
			expectedAdded:   1,
			expectedRemoved: 0,
		},
		{
			name:    "remove tag",
			desired: []*svcapitypes.Tag{},
			latest: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			expectedAdded:   0,
			expectedRemoved: 1,
		},
		{
			name: "update tag value",
			desired: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("newvalue")},
			},
			latest: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("oldvalue")},
			},
			expectedAdded:   1,
			expectedRemoved: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			added, removed := computeTagsDelta(tt.desired, tt.latest)
			if len(added) != tt.expectedAdded {
				t.Errorf("expected %d added tags, got %d", tt.expectedAdded, len(added))
			}
			if len(removed) != tt.expectedRemoved {
				t.Errorf("expected %d removed tags, got %d", tt.expectedRemoved, len(removed))
			}
		})
	}
}

func TestEqualTags(t *testing.T) {
	tests := []struct {
		name     string
		a        []*svcapitypes.Tag
		b        []*svcapitypes.Tag
		expected bool
	}{
		{
			name: "equal tags",
			a: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			b: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			expected: true,
		},
		{
			name: "different values",
			a: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			b: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value2")},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EqualTags(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
