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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Protcol is the underlying protocol used by the Object Storage bucket
type Protocol string

const (
	// ProtocolS3V4 is the s3v4 API standard
	ProtocolS3V4 Protocol = "s3v4"
)

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Protocol is the underlying protocol used by the Object Storage bucket. eg. s3v4, gcs etc.
	Protocol Protocol `json:"protocol"`

	// BucketPrefix is the prefix path to the desired folder in a bucket
	// +optional
	BucketPrefix string `json:"bucketPrefix,omitempty"`

	// BucketClassName is the name of the BucketClass where this Bucket belongs
	// +optional
	BucketClassName string `json:"bucketClassName,omitempty"`

	// AccessSecretName is the name of the Secret object where the credentials for accessing this bucket is available
	AccessSecretName string `json:"accessSecretName"`
}

// Phase denotes the status of this Bucket binding to a particular BucketContent
type Phase string

const (
	// PhaseBound denotes that the Bucket is bound to a BucketContent
	PhaseBound Phase = "bound"

	// PhasePending denotes that the Bucket is waiting to be bound to a BucketContent
	PhasePending Phase = "pending"
)

// BucketStatus defines the observed state of Bucket
type BucketStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// BucketContentName is the name of the BucketContent resource created/bound by the COSI controller
	// corresponding to the requested Bucket resource
	// +optional
	BucketContentName string `json:"bucketContentName,omitempty"`

	// Phase denotes if the Bucket is bound to a BucketContent or not
	Phase Phase `json:"phase"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the buckets API
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec,omitempty"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Bucket
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
