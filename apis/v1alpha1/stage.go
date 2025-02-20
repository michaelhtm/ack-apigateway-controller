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

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StageSpec defines the desired state of Stage.
//
// Represents a unique identifier for a version of a deployed RestApi that is
// callable by users.
type StageSpec struct {

// Whether cache clustering is enabled for the stage.
CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty"`
// The stage's cache capacity in GB. For more information about choosing a cache
// size, see Enabling API caching to enhance responsiveness (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).
CacheClusterSize *string `json:"cacheClusterSize,omitempty"`
// The canary deployment settings of this stage.
CanarySettings *CanarySettings `json:"canarySettings,omitempty"`
// The identifier of the Deployment resource for the Stage resource.
// +kubebuilder:validation:Required
DeploymentID *string `json:"deploymentID"`
// The description of the Stage resource.
Description *string `json:"description,omitempty"`
// The version of the associated API documentation.
DocumentationVersion *string `json:"documentationVersion,omitempty"`
// The string identifier of the associated RestApi.
// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
RestAPIID *string `json:"restAPIID,omitempty"`
RestAPIRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"restAPIRef,omitempty"`
// The name for the Stage resource. Stage names can only contain alphanumeric
// characters, hyphens, and underscores. Maximum length is 128 characters.
// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
// +kubebuilder:validation:Required
StageName *string `json:"stageName"`
// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
// The tag key can be up to 128 characters and must not start with aws:. The
// tag value can be up to 256 characters.
Tags map[string]*string `json:"tags,omitempty"`
// Specifies whether active tracing with X-ray is enabled for the Stage.
TracingEnabled *bool `json:"tracingEnabled,omitempty"`
// A map that defines the stage variables for the new Stage resource. Variable
// names can have alphanumeric and underscore characters, and the values must
// match [A-Za-z0-9-._~:/?#&=,]+.
Variables map[string]*string `json:"variables,omitempty"`
}

// StageStatus defines the observed state of Stage
type StageStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Settings for logging access in this stage.
	// +kubebuilder:validation:Optional
	AccessLogSettings *AccessLogSettings `json:"accessLogSettings,omitempty"`
	// The status of the cache cluster for the stage, if enabled.
	// +kubebuilder:validation:Optional
	CacheClusterStatus *string `json:"cacheClusterStatus,omitempty"`
	// The identifier of a client certificate for an API stage.
	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateID,omitempty"`
	// The timestamp when the stage was created.
	// +kubebuilder:validation:Optional
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// The timestamp when the stage last updated.
	// +kubebuilder:validation:Optional
	LastUpdatedDate *metav1.Time `json:"lastUpdatedDate,omitempty"`
	// A map that defines the method settings for a Stage resource. Keys (designated
// as /{method_setting_key below) are method paths defined as {resource_path}/{http_method}
// for an individual method override, or /\*/\* for overriding all methods in
// the stage.
	// +kubebuilder:validation:Optional
	MethodSettings map[string]*MethodSetting `json:"methodSettings,omitempty"`
	// The ARN of the WebAcl associated with the Stage.
	// +kubebuilder:validation:Optional
	WebACLARN *string `json:"webACLARN,omitempty"`
}

// Stage is the Schema for the Stages API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   StageSpec   `json:"spec,omitempty"`
	Status StageStatus `json:"status,omitempty"`
}

// StageList contains a list of Stage
// +kubebuilder:object:root=true
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []Stage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Stage{}, &StageList{})
}