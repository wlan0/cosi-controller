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

// BucketAccessSpec defines the desired state of BucketAccess
type BucketAccessSpec struct {
	// Protocol is the specification of the protocol used by the underlying Object Storage bucket
	Protocol ProtocolSpec `json:"protocol"`

	// ServiceAccountName is the name of the serviceAccount that authorizes access to keySecrets
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// ServiceAccountNamespace is namespace of the serviceAccount that authorizes access to keySecrets
	// +optional
	ServiceAccountNamespace string `json:"serviceAccountNamespace,omitempty"`

	// BucketAccessClassName is the name of the BucketClass associated with this BucketAccessSpec
	// +optional
	BucketAccessClassName string `json:"bucketAccessClassName,omitempty"`
}

// BucketAccessStatus defines the observed state of BucketAccess
type BucketAccessStatus struct {
	// KeySecretName is the driver generated secret for access object storage
	KeySecretName string `json:"keySecretName,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccess is the Schema for the bucketaccesses API
type BucketAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketAccessSpec   `json:"spec,omitempty"`
	Status BucketAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccessList contains a list of BucketAccess
type BucketAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccess `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BucketAccess{}, &BucketAccessList{})
}
