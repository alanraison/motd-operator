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
	Address    string `json:"address"`
	Username   string `json:"username"`
	PrivateKey string `json:"privateKey"`
}

// MotdSourceStatus defines the observed state of MotdSource
type MotdSourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Updated      metav1.Time `json:"updated"`
	ShortMessage string      `json:"short_message"`
	FullMessage  string      `json:"full_message"`
	Error        string      `json:"error,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MotdSource is the Schema for the motdsources API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=motdsources,scope=Namespaced
// +kubebuilder:printcolumn:name="Address",type="string",JSONPath=".spec.address"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.short_message"
// +kubebuilder:printcolumn:name="Error",type="string",JSONPath=".status.error"
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
