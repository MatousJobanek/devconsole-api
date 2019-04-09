package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GitSourceSpec defines the desired state of GitSource
// +k8s:openapi-gen=true
type GitSourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// URL of the git repo
	URL string `json:"url"`

	// Ref is a git reference. Optional. "master" is used by default.
	Ref string `json:"ref,omitempty"`

	// ContextDir is a path to subfolder in the repo. Optional.
	ContextDir string `json:"contextDir,omitempty"`

	// HttpProxy is optional.
	HttpProxy string `json:"httpProxy,omitempty"`

	// HttpsProxy is optional.
	HttpsProxy string `json:"httpsProxy,omitempty"`

	// NoProxy can be used to specify domains for which no proxying should be performed. Optional.
	NoProxy string `json:"noProxy,omitempty"`

	// SecretRef refers to the secret that contains credentials to access the git repo. Optional.
	SecretRef *SecretRef `json:"secretRef,omitempty"`

	// Flavor of the git provider like github, gitlab, bitbucket, generic, etc. Optional.
	Flavor string `json:"flavor,omitempty"`
}

// SecretRef holds information about the secret that contains credentials to access the git repo
type SecretRef struct {
	// Name is the name of the secret that contains credentials to access the git repo
	Name string `json:"name"`
}

// GitSourceStatus defines the observed state of GitSource
// +k8s:openapi-gen=true
type GitSourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// State represents current state of the GitSource, can be either initializing or ready
	State State `json:"state"`

	// Connection holds information whether the last attempt to reach the git source was successful or not. Optional
	Connection Connection `json:"state,omitempty"`
}

// State represents current state of the GitSource,
type State string

// Initializing represents a state of GitSource whose creation wasn't completed and is not ready to use
const Initializing State = "initializing"

// Ready represents a state of GitSource whose creation was completed and is ready to use
var Ready State = "ready"

// Connection holds information whether the last attempt to reach the git source was successful or not
type Connection struct {
	// State is the result of the attempt to reach a GitSource. Can be either Failed or OK
	State ConnectionState `json:"state"`

	// Error has the error message if the attempt to reach a GitSource failed
	Error string `json:"state,omitempty"`
}

// ConnectionState is the result of the attempt to reach a GitSource.
type ConnectionState string

// Failed is the state of Connection when an attempt to reach a GitSource failed
var Failed ConnectionState = "failed"

// OK is the state of Connection when an attempt to reach a GitSource was successful
var OK ConnectionState = "ok"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitSource is the Schema for the gitsources API
// +k8s:openapi-gen=true
type GitSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitSourceSpec   `json:"spec,omitempty"`
	Status GitSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitSourceList contains a list of GitSource
type GitSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitSource{}, &GitSourceList{})
}
