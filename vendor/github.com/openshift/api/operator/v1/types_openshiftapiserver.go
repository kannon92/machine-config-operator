package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=openshiftapiservers,scope=Cluster,categories=coreoperators
// +kubebuilder:subresource:status
// +openshift:api-approved.openshift.io=https://github.com/openshift/api/pull/475
// +openshift:file-pattern=0000_30_openshift-apiserver-operator_01_configMARKERS.crd.yaml

// OpenShiftAPIServer provides information to configure an operator to manage openshift-apiserver.
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type OpenShiftAPIServer struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	// spec is the specification of the desired behavior of the OpenShift API Server.
	// +kubebuilder:validation:Required
	// +required
	Spec OpenShiftAPIServerSpec `json:"spec"`

	// status defines the observed status of the OpenShift API Server.
	// +optional
	Status OpenShiftAPIServerStatus `json:"status"`
}

type OpenShiftAPIServerSpec struct {
	OperatorSpec `json:",inline"`
}

type OpenShiftAPIServerStatus struct {
	OperatorStatus `json:",inline"`

	// latestAvailableRevision is the latest revision used as suffix of revisioned
	// secrets like encryption-config. A new revision causes a new deployment of
	// pods.
	// +optional
	// +kubebuilder:validation:Minimum=0
	LatestAvailableRevision int32 `json:"latestAvailableRevision,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenShiftAPIServerList is a collection of items
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type OpenShiftAPIServerList struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard list's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items contains the items
	Items []OpenShiftAPIServer `json:"items"`
}
