/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Regions where AWS S3 API is available
type S3Region string

const (
	// S3RegionUSEast1 = us-east-1
	S3RegionUSEast1 S3Region = `json:"us-east-1"`
)

// S3ProtocolSpec is the specification for S3 and S3 compatible Object Storage services
type S3ProtocolSpec struct {
	// Endpoint is the URL for accessing the S3 endpoint
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// BucketName is the name of the S3 or S3 compatible bucket
	BucketName string `json:"bucket"`

	// Region is the Region where the bucket is situated
	// +optional
	Region S3Region `json:"region,omitempty"`

	// SignatureVersion is the version of the S3 API signature
	// +optional
	SignatureVersion string `json:"signatureVersion,omitempty"`

	// ServiceAccount is the service account for the BucketContent
	// +optional
	ServiceAccount string `json:"serviceAccount,omitempty"`
}

// ProtocolSpec defines the specification of the Object Storage protocol
type ProtocolSpec struct {
	// ProtocolType is the type of the underlying protocol for the associated Bucket
	ProtocolType Protocol `json:"type"`

	// S3ProtocolSpec is the specification for S3 and S3 compatible Object Storage services
	// +optional
	S3ProtocolSpec S3ProtocolSpec `json:"s3,omitempty"`

	// AzureProtocolSpec
	// GCSProtocolSpec
}

// ReleasePolicy defines the lifecycle of the bucket once the pod(s) using it have exited
type ReleasePolicy string

const (
	// ReleasePolicyRetain denotes that the bucket must be retained after it is used
	ReleasePolicyRetain ReleasePolicy = "retain"

	// ReleasePolicyDelete denotes that the bucket must be deleted after it is used
	ReleasePolicyDelete ReleasePolicy = "delete"
)

// BucketContentSpec defines the desired state of BucketContent
type BucketContentSpec struct {
	// Provisioner is the name of the provisioner as defined in the BucketClass of this BucketContent
	Provisioner string `json:"provisioner"`

	// ReleasePolicy defines the lifecycle of the bucket once the pod(s) using it have exited
	// +optional
	ReleasePolicy ReleasePolicy `json:"releasePolicy,omitempty"`

	// AccessMode defines the restriction of access allowed for this BucketContent
	// +optional
	AccessMode string `json:"accessMode,omitempty"`

	// BucketClassName defines the BucketClass where this Bucket belongs
	BucketClassName string `json:"bucketClassName"`

	// BucketRef is the reference to the Bucket that this BucketContent is bound to
	BucketRef string `json:"bucketRef"`

	// SecretRef is the reference to the Secret that contains the access and secret credentials for accessing the Object Storage bucket
	SecretRef *corev1.LocalObjectReference `json:"secretRef"`

	// Protocol is the specification of the protocol used by the underlying Object Storage bucket
	Protocol ProtocolSpec `json:"protocol"`

	// Parameters are any additional parameters to be passed in to the Object Storage provisioner
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

// BucketContentStatus defines the observed state of BucketContent
type BucketContentStatus struct {
}

// +kubebuilder:object:root=true

// BucketContent is the Schema for the bucketcontents API
type BucketContent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketContentSpec   `json:"spec,omitempty"`
	Status BucketContentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketContentList contains a list of BucketContent
type BucketContentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketContent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BucketContent{}, &BucketContentList{})
}
