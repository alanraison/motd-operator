package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MotdSourceSpec defines the desired state of MotdSource
type MotdSourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Hostname   string `json:"hostname"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	PrivateKey string `json:"privateKey"`
}

// MotdSourceStatus defines the observed state of MotdSource
type MotdSourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Updated int64  `json:"updated"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MotdSource is the Schema for the motdsources API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=motdsources,scope=Namespaced
type MotdSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MotdSourceSpec   `json:"spec,omitempty"`
	Status MotdSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MotdSourceList contains a list of MotdSource
type MotdSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MotdSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MotdSource{}, &MotdSourceList{})
}
