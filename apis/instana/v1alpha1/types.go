/*
Copyright 2020 The Crossplane Authors.

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

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// MyTypeParameters are the configurable fields of a MyType.
type InstanaParameters struct {
	InstanaVersion    string            `json:"instanaVersion"`
	NFSServerHost     string            `json:"nfsServerHost"`
	SettingsConfigmap SettingsConfigmap `json:"settingsConfigmap"`
}

type SettingsConfigmap struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// MyTypeObservation are the observable fields of a MyType.
type InstanaObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A MyTypeSpec defines the desired state of a MyType.
type InstanaSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InstanaParameters `json:"forProvider"`
}

// A MyTypeStatus represents the observed state of a MyType.
type InstanaStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InstanaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instana describes how to install operator based instana on k8s.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// Please replace `PROVIDER-NAME` with your actual provider name, like `aws`, `azure`, `gcp`, `alibaba`
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,instana}
type Instana struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanaSpec   `json:"spec"`
	Status InstanaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InfraList contains a list of Infra
type InstanaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instana `json:"items"`
}

