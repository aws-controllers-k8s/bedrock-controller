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

package v1alpha1

type ApplicationType string

const (
	ApplicationType_ModelEvaluation ApplicationType = "ModelEvaluation"
	ApplicationType_RagEvaluation   ApplicationType = "RagEvaluation"
)

type CommitmentDuration string

const (
	CommitmentDuration_OneMonth  CommitmentDuration = "OneMonth"
	CommitmentDuration_SixMonths CommitmentDuration = "SixMonths"
)

type CustomizationType string

const (
	CustomizationType_CONTINUED_PRE_TRAINING CustomizationType = "CONTINUED_PRE_TRAINING"
	CustomizationType_DISTILLATION           CustomizationType = "DISTILLATION"
	CustomizationType_FINE_TUNING            CustomizationType = "FINE_TUNING"
)

type EvaluationJobStatus string

const (
	EvaluationJobStatus_Completed  EvaluationJobStatus = "Completed"
	EvaluationJobStatus_Deleting   EvaluationJobStatus = "Deleting"
	EvaluationJobStatus_Failed     EvaluationJobStatus = "Failed"
	EvaluationJobStatus_InProgress EvaluationJobStatus = "InProgress"
	EvaluationJobStatus_Stopped    EvaluationJobStatus = "Stopped"
	EvaluationJobStatus_Stopping   EvaluationJobStatus = "Stopping"
)

type EvaluationJobType string

const (
	EvaluationJobType_Automated EvaluationJobType = "Automated"
	EvaluationJobType_Human     EvaluationJobType = "Human"
)

type EvaluationTaskType string

const (
	EvaluationTaskType_Classification    EvaluationTaskType = "Classification"
	EvaluationTaskType_Custom            EvaluationTaskType = "Custom"
	EvaluationTaskType_Generation        EvaluationTaskType = "Generation"
	EvaluationTaskType_QuestionAndAnswer EvaluationTaskType = "QuestionAndAnswer"
	EvaluationTaskType_Summarization     EvaluationTaskType = "Summarization"
)

type ExternalSourceType string

const (
	ExternalSourceType_BYTE_CONTENT ExternalSourceType = "BYTE_CONTENT"
	ExternalSourceType_S3           ExternalSourceType = "S3"
)

type FineTuningJobStatus string

const (
	FineTuningJobStatus_Completed  FineTuningJobStatus = "Completed"
	FineTuningJobStatus_Failed     FineTuningJobStatus = "Failed"
	FineTuningJobStatus_InProgress FineTuningJobStatus = "InProgress"
	FineTuningJobStatus_Stopped    FineTuningJobStatus = "Stopped"
	FineTuningJobStatus_Stopping   FineTuningJobStatus = "Stopping"
)

type FoundationModelLifecycleStatus string

const (
	FoundationModelLifecycleStatus_ACTIVE FoundationModelLifecycleStatus = "ACTIVE"
	FoundationModelLifecycleStatus_LEGACY FoundationModelLifecycleStatus = "LEGACY"
)

type GuardrailContentFilterType string

const (
	GuardrailContentFilterType_HATE          GuardrailContentFilterType = "HATE"
	GuardrailContentFilterType_INSULTS       GuardrailContentFilterType = "INSULTS"
	GuardrailContentFilterType_MISCONDUCT    GuardrailContentFilterType = "MISCONDUCT"
	GuardrailContentFilterType_PROMPT_ATTACK GuardrailContentFilterType = "PROMPT_ATTACK"
	GuardrailContentFilterType_SEXUAL        GuardrailContentFilterType = "SEXUAL"
	GuardrailContentFilterType_VIOLENCE      GuardrailContentFilterType = "VIOLENCE"
)

type GuardrailContextualGroundingFilterType string

const (
	GuardrailContextualGroundingFilterType_GROUNDING GuardrailContextualGroundingFilterType = "GROUNDING"
	GuardrailContextualGroundingFilterType_RELEVANCE GuardrailContextualGroundingFilterType = "RELEVANCE"
)

type GuardrailFilterStrength string

const (
	GuardrailFilterStrength_HIGH   GuardrailFilterStrength = "HIGH"
	GuardrailFilterStrength_LOW    GuardrailFilterStrength = "LOW"
	GuardrailFilterStrength_MEDIUM GuardrailFilterStrength = "MEDIUM"
	GuardrailFilterStrength_NONE   GuardrailFilterStrength = "NONE"
)

type GuardrailManagedWordsType string

const (
	GuardrailManagedWordsType_PROFANITY GuardrailManagedWordsType = "PROFANITY"
)

type GuardrailModality string

const (
	GuardrailModality_IMAGE GuardrailModality = "IMAGE"
	GuardrailModality_TEXT  GuardrailModality = "TEXT"
)

type GuardrailPiiEntityType string

const (
	GuardrailPiiEntityType_ADDRESS                                 GuardrailPiiEntityType = "ADDRESS"
	GuardrailPiiEntityType_AGE                                     GuardrailPiiEntityType = "AGE"
	GuardrailPiiEntityType_AWS_ACCESS_KEY                          GuardrailPiiEntityType = "AWS_ACCESS_KEY"
	GuardrailPiiEntityType_AWS_SECRET_KEY                          GuardrailPiiEntityType = "AWS_SECRET_KEY"
	GuardrailPiiEntityType_CA_HEALTH_NUMBER                        GuardrailPiiEntityType = "CA_HEALTH_NUMBER"
	GuardrailPiiEntityType_CA_SOCIAL_INSURANCE_NUMBER              GuardrailPiiEntityType = "CA_SOCIAL_INSURANCE_NUMBER"
	GuardrailPiiEntityType_CREDIT_DEBIT_CARD_CVV                   GuardrailPiiEntityType = "CREDIT_DEBIT_CARD_CVV"
	GuardrailPiiEntityType_CREDIT_DEBIT_CARD_EXPIRY                GuardrailPiiEntityType = "CREDIT_DEBIT_CARD_EXPIRY"
	GuardrailPiiEntityType_CREDIT_DEBIT_CARD_NUMBER                GuardrailPiiEntityType = "CREDIT_DEBIT_CARD_NUMBER"
	GuardrailPiiEntityType_DRIVER_ID                               GuardrailPiiEntityType = "DRIVER_ID"
	GuardrailPiiEntityType_EMAIL                                   GuardrailPiiEntityType = "EMAIL"
	GuardrailPiiEntityType_INTERNATIONAL_BANK_ACCOUNT_NUMBER       GuardrailPiiEntityType = "INTERNATIONAL_BANK_ACCOUNT_NUMBER"
	GuardrailPiiEntityType_IP_ADDRESS                              GuardrailPiiEntityType = "IP_ADDRESS"
	GuardrailPiiEntityType_LICENSE_PLATE                           GuardrailPiiEntityType = "LICENSE_PLATE"
	GuardrailPiiEntityType_MAC_ADDRESS                             GuardrailPiiEntityType = "MAC_ADDRESS"
	GuardrailPiiEntityType_NAME                                    GuardrailPiiEntityType = "NAME"
	GuardrailPiiEntityType_PASSWORD                                GuardrailPiiEntityType = "PASSWORD"
	GuardrailPiiEntityType_PHONE                                   GuardrailPiiEntityType = "PHONE"
	GuardrailPiiEntityType_PIN                                     GuardrailPiiEntityType = "PIN"
	GuardrailPiiEntityType_SWIFT_CODE                              GuardrailPiiEntityType = "SWIFT_CODE"
	GuardrailPiiEntityType_UK_NATIONAL_HEALTH_SERVICE_NUMBER       GuardrailPiiEntityType = "UK_NATIONAL_HEALTH_SERVICE_NUMBER"
	GuardrailPiiEntityType_UK_NATIONAL_INSURANCE_NUMBER            GuardrailPiiEntityType = "UK_NATIONAL_INSURANCE_NUMBER"
	GuardrailPiiEntityType_UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER     GuardrailPiiEntityType = "UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER"
	GuardrailPiiEntityType_URL                                     GuardrailPiiEntityType = "URL"
	GuardrailPiiEntityType_USERNAME                                GuardrailPiiEntityType = "USERNAME"
	GuardrailPiiEntityType_US_BANK_ACCOUNT_NUMBER                  GuardrailPiiEntityType = "US_BANK_ACCOUNT_NUMBER"
	GuardrailPiiEntityType_US_BANK_ROUTING_NUMBER                  GuardrailPiiEntityType = "US_BANK_ROUTING_NUMBER"
	GuardrailPiiEntityType_US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER GuardrailPiiEntityType = "US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER"
	GuardrailPiiEntityType_US_PASSPORT_NUMBER                      GuardrailPiiEntityType = "US_PASSPORT_NUMBER"
	GuardrailPiiEntityType_US_SOCIAL_SECURITY_NUMBER               GuardrailPiiEntityType = "US_SOCIAL_SECURITY_NUMBER"
	GuardrailPiiEntityType_VEHICLE_IDENTIFICATION_NUMBER           GuardrailPiiEntityType = "VEHICLE_IDENTIFICATION_NUMBER"
)

type GuardrailSensitiveInformationAction string

const (
	GuardrailSensitiveInformationAction_ANONYMIZE GuardrailSensitiveInformationAction = "ANONYMIZE"
	GuardrailSensitiveInformationAction_BLOCK     GuardrailSensitiveInformationAction = "BLOCK"
)

type GuardrailStatus string

const (
	GuardrailStatus_CREATING   GuardrailStatus = "CREATING"
	GuardrailStatus_DELETING   GuardrailStatus = "DELETING"
	GuardrailStatus_FAILED     GuardrailStatus = "FAILED"
	GuardrailStatus_READY      GuardrailStatus = "READY"
	GuardrailStatus_UPDATING   GuardrailStatus = "UPDATING"
	GuardrailStatus_VERSIONING GuardrailStatus = "VERSIONING"
)

type GuardrailTopicType string

const (
	GuardrailTopicType_DENY GuardrailTopicType = "DENY"
)

type InferenceProfileStatus string

const (
	InferenceProfileStatus_ACTIVE InferenceProfileStatus = "ACTIVE"
)

type InferenceProfileType string

const (
	InferenceProfileType_APPLICATION    InferenceProfileType = "APPLICATION"
	InferenceProfileType_SYSTEM_DEFINED InferenceProfileType = "SYSTEM_DEFINED"
)

type InferenceType string

const (
	InferenceType_ON_DEMAND   InferenceType = "ON_DEMAND"
	InferenceType_PROVISIONED InferenceType = "PROVISIONED"
)

type ModelCopyJobStatus string

const (
	ModelCopyJobStatus_Completed  ModelCopyJobStatus = "Completed"
	ModelCopyJobStatus_Failed     ModelCopyJobStatus = "Failed"
	ModelCopyJobStatus_InProgress ModelCopyJobStatus = "InProgress"
)

type ModelCustomization string

const (
	ModelCustomization_CONTINUED_PRE_TRAINING ModelCustomization = "CONTINUED_PRE_TRAINING"
	ModelCustomization_DISTILLATION           ModelCustomization = "DISTILLATION"
	ModelCustomization_FINE_TUNING            ModelCustomization = "FINE_TUNING"
)

type ModelCustomizationJobStatus string

const (
	ModelCustomizationJobStatus_Completed  ModelCustomizationJobStatus = "Completed"
	ModelCustomizationJobStatus_Failed     ModelCustomizationJobStatus = "Failed"
	ModelCustomizationJobStatus_InProgress ModelCustomizationJobStatus = "InProgress"
	ModelCustomizationJobStatus_Stopped    ModelCustomizationJobStatus = "Stopped"
	ModelCustomizationJobStatus_Stopping   ModelCustomizationJobStatus = "Stopping"
)

type ModelImportJobStatus string

const (
	ModelImportJobStatus_Completed  ModelImportJobStatus = "Completed"
	ModelImportJobStatus_Failed     ModelImportJobStatus = "Failed"
	ModelImportJobStatus_InProgress ModelImportJobStatus = "InProgress"
)

type ModelInvocationJobStatus string

const (
	ModelInvocationJobStatus_Completed          ModelInvocationJobStatus = "Completed"
	ModelInvocationJobStatus_Expired            ModelInvocationJobStatus = "Expired"
	ModelInvocationJobStatus_Failed             ModelInvocationJobStatus = "Failed"
	ModelInvocationJobStatus_InProgress         ModelInvocationJobStatus = "InProgress"
	ModelInvocationJobStatus_PartiallyCompleted ModelInvocationJobStatus = "PartiallyCompleted"
	ModelInvocationJobStatus_Scheduled          ModelInvocationJobStatus = "Scheduled"
	ModelInvocationJobStatus_Stopped            ModelInvocationJobStatus = "Stopped"
	ModelInvocationJobStatus_Stopping           ModelInvocationJobStatus = "Stopping"
	ModelInvocationJobStatus_Submitted          ModelInvocationJobStatus = "Submitted"
	ModelInvocationJobStatus_Validating         ModelInvocationJobStatus = "Validating"
)

type ModelModality string

const (
	ModelModality_EMBEDDING ModelModality = "EMBEDDING"
	ModelModality_IMAGE     ModelModality = "IMAGE"
	ModelModality_TEXT      ModelModality = "TEXT"
)

type PerformanceConfigLatency string

const (
	PerformanceConfigLatency_optimized PerformanceConfigLatency = "optimized"
	PerformanceConfigLatency_standard  PerformanceConfigLatency = "standard"
)

type PromptRouterStatus string

const (
	PromptRouterStatus_AVAILABLE PromptRouterStatus = "AVAILABLE"
)

type PromptRouterType string

const (
	PromptRouterType_custom  PromptRouterType = "custom"
	PromptRouterType_default PromptRouterType = "default"
)

type ProvisionedModelStatus string

const (
	ProvisionedModelStatus_Creating  ProvisionedModelStatus = "Creating"
	ProvisionedModelStatus_Failed    ProvisionedModelStatus = "Failed"
	ProvisionedModelStatus_InService ProvisionedModelStatus = "InService"
	ProvisionedModelStatus_Updating  ProvisionedModelStatus = "Updating"
)

type QueryTransformationType string

const (
	QueryTransformationType_QUERY_DECOMPOSITION QueryTransformationType = "QUERY_DECOMPOSITION"
)

type RetrieveAndGenerateType string

const (
	RetrieveAndGenerateType_EXTERNAL_SOURCES RetrieveAndGenerateType = "EXTERNAL_SOURCES"
	RetrieveAndGenerateType_KNOWLEDGE_BASE   RetrieveAndGenerateType = "KNOWLEDGE_BASE"
)

type S3InputFormat string

const (
	S3InputFormat_JSONL S3InputFormat = "JSONL"
)

type SearchType string

const (
	SearchType_HYBRID   SearchType = "HYBRID"
	SearchType_SEMANTIC SearchType = "SEMANTIC"
)

type SortByProvisionedModels string

const (
	SortByProvisionedModels_CreationTime SortByProvisionedModels = "CreationTime"
)

type SortJobsBy string

const (
	SortJobsBy_CreationTime SortJobsBy = "CreationTime"
)

type SortModelsBy string

const (
	SortModelsBy_CreationTime SortModelsBy = "CreationTime"
)

type SortOrder string

const (
	SortOrder_Ascending  SortOrder = "Ascending"
	SortOrder_Descending SortOrder = "Descending"
)

type Status string

const (
	Status_INCOMPATIBLE_ENDPOINT Status = "INCOMPATIBLE_ENDPOINT"
	Status_REGISTERED            Status = "REGISTERED"
)
