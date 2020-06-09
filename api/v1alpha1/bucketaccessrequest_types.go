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

// BucketAccessRequestSpec defines the desired state of BucketAccessRequest
type BucketAccessRequestSpec struct {
	// ServiceAccountName is the name of the serviceAccount associated with cloud credentials such as AWS IAM and GCP equivalent
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// AccessSecretName is the name of the secret that the app will utilize to access the Bucket
	// +optional
	AccessSecretName *corev1.LocalObjectReference `json:"accessSecretName,omitempty"`

	// Bucket is the name of the Bucket that will be access by the app
	Bucket string `json:"bucket"`

	// BucketAccessClassName is the name of the BucketAccessClass associated with this BucketAccessRequest
	// +optional
	BucketAccessClassName string `json:"bucketAccessClassName,omitempty"`

	// BucketAccessName is the name of the BucketAccess associated with this BucketAccessRequest
	// +optional
	BucketAccessName string `json:"bucketAccessname,omitempty"`

	// Protocol is the specification of the protocol used by the underlying Object Storage bucket
	Protocol Protocol `json:"protocol"`
}

// BucketAccessRequestStatus defines the observed state of BucketAccessRequest
type BucketAccessRequestStatus struct {
}

// +kubebuilder:object:root=true

// BucketAccessRequest is the Schema for the bucketaccessrequests API
type BucketAccessRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketAccessRequestSpec   `json:"spec,omitempty"`
	Status BucketAccessRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccessRequestList contains a list of BucketAccessRequest
type BucketAccessRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccessRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BucketAccessRequest{}, &BucketAccessRequestList{})
}
