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

// BucketAccessPolicy is the set of authorization rules for accessing buckets
type BucketAccessPolicy struct {
	// Allow is the list of allowed operations
	Allow []string `json:"allow,omitempty"`

	// Deny is the list of restricted operations
	Deny []string `json:"deny,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccessClass is the Schema for the bucketaccessclasses API
type BucketAccessClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Provisioner is the name of the provisioner for the class of BucketAccess
	Provisioner string `json:"provisioner,omitempty"`

	// SupportedProtocols is a list of protocols that this BucketAccessClass supports
	// +optional
	SupportedProtocols []Protocol `json:"supportedProtocols,omitempty"`

	// AccessPolicy is the authorization rules for accessing buckets
	AccessPolicy BucketAccessPolicy `json:"accessPolicy,omitempty"`

	// Parameters are additional parameters to support custom fields
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccessClassList contains a list of BucketAccessClass
type BucketAccessClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccessClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BucketAccessClass{}, &BucketAccessClassList{})
}
