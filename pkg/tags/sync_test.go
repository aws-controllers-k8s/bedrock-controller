package tags

import (
	"context"
	"errors"
	"slices"
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

	tagCalls   []*svcsdk.TagResourceInput
	untagCalls []*svcsdk.UntagResourceInput
}

func (m *mockTagsClient) TagResource(ctx context.Context, input *svcsdk.TagResourceInput, opts ...func(*svcsdk.Options)) (*svcsdk.TagResourceOutput, error) {
	m.tagCalls = append(m.tagCalls, input)
	return &svcsdk.TagResourceOutput{}, m.tagResourceError
}

func (m *mockTagsClient) ListTagsForResource(ctx context.Context, input *svcsdk.ListTagsForResourceInput, opts ...func(*svcsdk.Options)) (*svcsdk.ListTagsForResourceOutput, error) {
	return m.listTagsResponse, m.listTagsError
}

func (m *mockTagsClient) UntagResource(ctx context.Context, input *svcsdk.UntagResourceInput, opts ...func(*svcsdk.Options)) (*svcsdk.UntagResourceOutput, error) {
	m.untagCalls = append(m.untagCalls, input)
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

func TestSyncResourceTags(t *testing.T) {
	errUntag := errors.New("untag failed")
	errTag := errors.New("tag failed")

	tests := []struct {
		name            string
		desired         []*svcapitypes.Tag
		latest          []*svcapitypes.Tag
		expectedTagged  map[string]string
		expectedUntag   []string
		expectTagCall   bool
		expectUntagCall bool
		untagErr        error
		tagErr          error
		wantErr         error
		unchanged       []string
	}{
		{
			name: "no-op when tags are equal",
			desired: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			latest: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
		},
		{
			name:    "no-op when both are empty",
			desired: []*svcapitypes.Tag{},
			latest:  []*svcapitypes.Tag{},
		},
		{
			name: "add new tags",
			desired: []*svcapitypes.Tag{
				{Key: aws.String("key2"), Value: aws.String("value2")},
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			latest:         []*svcapitypes.Tag{},
			expectedTagged: map[string]string{"key1": "value1", "key2": "value2"},
			expectTagCall:  true,
		},
		{
			name:    "remove tag",
			desired: []*svcapitypes.Tag{},
			latest: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("value1")},
			},
			expectedUntag:   []string{"key1"},
			expectUntagCall: true,
		},
		{
			name:    "remove multiple tags sends sorted keys",
			desired: []*svcapitypes.Tag{},
			latest: []*svcapitypes.Tag{
				{Key: aws.String("zeta"), Value: aws.String("z")},
				{Key: aws.String("alpha"), Value: aws.String("a")},
				{Key: aws.String("mid"), Value: aws.String("m")},
			},
			expectedUntag:   []string{"zeta", "alpha", "mid"}, // deliberately unsorted
			expectUntagCall: true,
		},
		{
			name: "value change only calls TagResource",
			desired: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("newvalue")},
			},
			latest: []*svcapitypes.Tag{
				{Key: aws.String("key1"), Value: aws.String("oldvalue")},
			},
			expectedTagged: map[string]string{"key1": "newvalue"},
			expectTagCall:  true,
		},
		{
			name: "add, remove, update and keep unchanged together",
			desired: []*svcapitypes.Tag{
				{Key: aws.String("keep"), Value: aws.String("v")},
				{Key: aws.String("new"), Value: aws.String("n")},
				{Key: aws.String("same"), Value: aws.String("s")},
			},
			latest: []*svcapitypes.Tag{
				{Key: aws.String("keep"), Value: aws.String("old")},
				{Key: aws.String("gone"), Value: aws.String("g")},
				{Key: aws.String("same"), Value: aws.String("s")},
			},
			expectedTagged:  map[string]string{"keep": "v", "new": "n"},
			expectedUntag:   []string{"gone"},
			expectTagCall:   true,
			expectUntagCall: true,
			unchanged:       []string{"same"},
		},
		{
			name: "nil key does not panic and is skipped",
			desired: []*svcapitypes.Tag{
				{Key: nil, Value: aws.String("orphan")},
				{Key: aws.String("key1"), Value: nil},
			},
			latest:         []*svcapitypes.Tag{nil},
			expectedTagged: map[string]string{"key1": ""},
			expectTagCall:  true,
		},
		{
			name:            "UntagResource error stops before TagResource",
			desired:         []*svcapitypes.Tag{{Key: aws.String("new"), Value: aws.String("n")}},
			latest:          []*svcapitypes.Tag{{Key: aws.String("gone"), Value: aws.String("g")}},
			untagErr:        errUntag,
			wantErr:         errUntag,
			expectedUntag:   []string{"gone"},
			expectUntagCall: true,
			expectTagCall:   false,
		},
		{
			name:            "TagResource error is propagated",
			desired:         []*svcapitypes.Tag{{Key: aws.String("new"), Value: aws.String("n")}},
			latest:          []*svcapitypes.Tag{{Key: aws.String("gone"), Value: aws.String("g")}},
			tagErr:          errTag,
			wantErr:         errTag,
			expectedUntag:   []string{"gone"},
			expectUntagCall: true,
			expectedTagged:  map[string]string{"new": "n"},
			expectTagCall:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockTagsClient{untagError: tt.untagErr, tagResourceError: tt.tagErr}
			metrics := &mockMetrics{}

			err := SyncResourceTags(context.Background(), client, metrics, "test-arn", tt.desired, tt.latest)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tt.expectTagCall && len(client.tagCalls) != 0 {
				t.Errorf("expected no TagResource call, got %d", len(client.tagCalls))
			}
			if tt.expectTagCall {
				if len(client.tagCalls) != 1 {
					t.Fatalf("expected 1 TagResource call, got %d", len(client.tagCalls))
				}
				got := map[string]string{}
				for _, tag := range client.tagCalls[0].Tags {
					got[*tag.Key] = *tag.Value
				}
				if len(got) != len(tt.expectedTagged) {
					t.Errorf("expected tagged %v, got %v", tt.expectedTagged, got)
				}
				for k, v := range tt.expectedTagged {
					if got[k] != v {
						t.Errorf("expected tag %s=%s, got %q", k, v, got[k])
					}
				}
			}

			if !tt.expectUntagCall && len(client.untagCalls) != 0 {
				t.Errorf("expected no UntagResource call, got %d", len(client.untagCalls))
			}
			if tt.expectUntagCall {
				if len(client.untagCalls) != 1 {
					t.Fatalf("expected 1 UntagResource call, got %d", len(client.untagCalls))
				}
				got := client.untagCalls[0].TagKeys
				want := slices.Sorted(slices.Values(tt.expectedUntag))
				if !slices.Equal(got, want) {
					t.Errorf("untag keys: got %v, want %v", got, want)
				}
			}

			for _, k := range tt.unchanged {
				for _, call := range client.tagCalls {
					for _, t2 := range call.Tags {
						if t2.Key != nil && *t2.Key == k {
							t.Errorf("unchanged tag %q was sent to TagResource", k)
						}
					}
				}
				for _, call := range client.untagCalls {
					if slices.Contains(call.TagKeys, k) {
						t.Errorf("unchanged tag %q was sent to UntagResource", k)
					}
				}
			}
		})
	}
}
