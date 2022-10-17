/*
Copyright 2022.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HarborRegistryConfigurationSpec defines the desired state of HarborRegistryConfiguration
type HarborRegistryConfigurationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of HarborRegistryConfiguration. Edit harborregistryconfiguration_types.go to remove/update

	HarborTarget    HarborTarget    `json:"harborTarget"`
	RegistryOptions RegistryOptions `json:"registryOptions"`
}

// HarborRegistryConfigurationStatus defines the observed state of HarborRegistryConfiguration
type HarborRegistryConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HarborRegistryConfiguration is the Schema for the harborregistryconfigurations API
type HarborRegistryConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HarborRegistryConfigurationSpec   `json:"spec,omitempty"`
	Status HarborRegistryConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HarborRegistryConfigurationList contains a list of HarborRegistryConfiguration
type HarborRegistryConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HarborRegistryConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HarborRegistryConfiguration{}, &HarborRegistryConfigurationList{})
}

type RegistryOptions struct {
	Name              string              `json:"name"`
	Type              string              `json:"type"`
	TargetRegistryUrl string              `json:"targetRegistryUrl"`
	Description       string              `json:"description,omitempty"`
	Credential        *RegistryCredential `json:"credential,omitempty"`
}

type HarborTarget struct {
	ApiUrl   string `json:"apiUrl"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistryCredential struct {

	// Access key, e.g. user name when credential type is 'basic'.
	AccessKey string `json:"access_key,omitempty"`

	// Access secret, e.g. password when credential type is 'basic'.
	AccessSecret string `json:"access_secret,omitempty"`

	// Credential type, such as 'basic', 'oauth'.
	Type string `json:"type,omitempty"`
}
