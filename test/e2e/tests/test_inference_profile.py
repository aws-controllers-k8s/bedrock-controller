# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for Inference Profile API.
"""

import time
import pytest

from acktest.k8s import condition
from acktest.k8s import resource as k8s
from acktest import tags
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_bedrock_resource
from e2e.replacement_values import REPLACEMENT_VALUES

INFERENCE_PROFILE_RESOURCE_PLURAL="inferenceprofiles"
TEST_DESCRIPTION = "Test profile"
TEST_MODEL_ID = "arn:aws:bedrock:us-west-2::foundation-model/amazon.titan-image-generator-v2:0"
TAG_KEY_1 = "environment"
TAG_VALUE_1 = "dev"
TAG_KEY_2 = "project"
TAG_VALUE_2 = "ack"

CHECK_WAIT_SECONDS = 5

@pytest.fixture(scope="module")
def simple_inference_profile():
    profile_name = random_suffix_name("bedrock-test-profile", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["INFERENCE_PROFILE_NAME"] = profile_name
    replacements["INFERENCE_PROFILE_DESCRIPTION"] = TEST_DESCRIPTION
    replacements["SOURCE_MODEL_ID"] = TEST_MODEL_ID
    replacements["TAG_KEY_1"] = TAG_KEY_1
    replacements["TAG_VALUE_1"] = TAG_VALUE_1
    replacements["TAG_KEY_2"] = TAG_KEY_2
    replacements["TAG_VALUE_2"] = TAG_VALUE_2

    resource_data = load_bedrock_resource(
        "inference_profile",
        additional_replacements=replacements,
    )

    ref = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        INFERENCE_PROFILE_RESOURCE_PLURAL,
        profile_name,
        namespace="default",
    )

    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    yield (ref, cr)

    # Delete Inference Profile
    try:
        _, deleted = k8s.delete_custom_resource(ref, wait_periods=3, period_length=10)
        assert deleted
    except:
        pass

@service_marker
@pytest.mark.canary
class TestInferenceProfile:
    def test_create_delete_inference_profile(self, simple_inference_profile, bedrock_client):
        (ref, cr) = simple_inference_profile

        k8s.wait_on_condition(ref, "ACK.ResourceSynced", "True", wait_periods=5)
        condition.assert_synced(ref)

        cr = k8s.get_resource(ref)
        inference_profile_arn = cr["status"]["ackResourceMetadata"]["arn"]
        assert inference_profile_arn is not None
        assert cr["spec"]["description"] == TEST_DESCRIPTION
        assert cr["spec"]["modelSource"]["copyFrom"] == TEST_MODEL_ID
        assert cr["status"]["status"] == "ACTIVE"

        expected_tags = [
            {"key": TAG_KEY_1, "value": TAG_VALUE_1},
            {"key": TAG_KEY_2, "value": TAG_VALUE_2},
        ]
        tags.assert_equal_without_ack_tags(expected_tags, cr["spec"]["tags"], key_member_name="key", value_member_name="value")

        aws_profile = bedrock_client.get_inference_profile(inferenceProfileIdentifier=inference_profile_arn)
        assert aws_profile["description"] == TEST_DESCRIPTION
        assert aws_profile["models"][0]["modelArn"] == TEST_MODEL_ID

        aws_tags_response = bedrock_client.list_tags_for_resource(resourceARN=inference_profile_arn)
        aws_tags = [{"key": tag["key"], "value": tag["value"]} for tag in aws_tags_response["tags"]]
        tags.assert_equal_without_ack_tags(expected_tags, aws_tags, key_member_name="key", value_member_name="value")
        
        updated_tags = [
            {"key": TAG_KEY_1, "value": "updated-dev"},
            {"key": TAG_KEY_2, "value": "updated-ack"},
        ]
        cr["spec"]["tags"] = updated_tags
        k8s.patch_custom_resource(ref, cr)
        time.sleep(CHECK_WAIT_SECONDS)
        
        k8s.wait_on_condition(ref, "ACK.ResourceSynced", "True", wait_periods=5)
        
        cr_updated = k8s.get_resource(ref)
        tags.assert_equal_without_ack_tags(updated_tags, cr_updated["spec"]["tags"], key_member_name="key", value_member_name="value")
        
        aws_tags_updated_response = bedrock_client.list_tags_for_resource(resourceARN=inference_profile_arn)
        aws_tags_updated = [{"key": tag["key"], "value": tag["value"]} for tag in aws_tags_updated_response["tags"]]
        tags.assert_equal_without_ack_tags(updated_tags, aws_tags_updated, key_member_name="key", value_member_name="value")