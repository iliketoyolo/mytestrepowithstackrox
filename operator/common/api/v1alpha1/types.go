//+kubebuilder:object:generate=true

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make generate" in the common directory to regenerate code,
// and run "make manifests" in the "central" and "securedcluster" directories
// to regenerate manifests, after modifying this file
// TODO(ROX-7110): prevent merging PRs if manifests are not up to date.

// CustomizeSpec defines customizations to apply.
type CustomizeSpec struct {
	// Custom labels to set on all objects apart from Pods.
	Labels map[string]string `json:"labels,omitempty"`
	// Custom annotations to set on all objects apart from Pods.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Custom labels to set on Pods.
	PodLabels map[string]string `json:"podLabels,omitempty"`
	// Custom annotations to set on Pods.
	PodAnnotations map[string]string `json:"podAnnotations,omitempty"`
	// Custom environment variables to set on pods' containers.
	EnvVars map[string]string `json:"envVars,omitempty"`
}

// DeploymentSpec defines settings that affect a deployment.
type DeploymentSpec struct {
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	Resources    *Resources        `json:"resources,omitempty"`
	// Secret that contains cert and key. Omit means: autogenerate.
	ServiceTLS *corev1.LocalObjectReference `json:"serviceTLS,omitempty"`
	// Customizations to apply on this deployment.
	Customize *CustomizeSpec `json:"customize,omitempty"`
	// TODO(ROX-7150): We do not support setting image in the CRs because they are determined by
	// the operator version whose lifecycle is orthogonal to that of the CR.
}

// Resources define resource-related settings.
type Resources struct {
	// Override allow overriding resource requirements for pods of this deployment,
	// the defaults are computed by the operator.
	// Currently that computation simply means using the defaults from the Helm charts.
	Override *corev1.ResourceRequirements `json:"override,omitempty"`
	// TODO(ROX-7146): potentially add a Cap field once we support vertical autoscaling.
}

// StackRoxCondition defines a condition for a StackRox custom resource.
type StackRoxCondition struct {
	Type    ConditionType   `json:"type"`
	Status  ConditionStatus `json:"status"`
	Reason  ConditionReason `json:"reason,omitempty"`
	Message string          `json:"message,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// ConditionType is a type of values of condition type.
type ConditionType string

// ConditionStatus is a type of values of condition status.
type ConditionStatus string

// ConditionReason is a type of values of condition reason.
type ConditionReason string

// These are the allowed values for StackRoxCondition fields.
const (
	ConditionInitialized    ConditionType = "Initialized"
	ConditionDeployed       ConditionType = "Deployed"
	ConditionReleaseFailed  ConditionType = "ReleaseFailed"
	ConditionIrreconcilable ConditionType = "Irreconcilable"

	StatusTrue    ConditionStatus = "True"
	StatusFalse   ConditionStatus = "False"
	StatusUnknown ConditionStatus = "Unknown"

	ReasonInstallSuccessful   ConditionReason = "InstallSuccessful"
	ReasonUpgradeSuccessful   ConditionReason = "UpgradeSuccessful"
	ReasonUninstallSuccessful ConditionReason = "UninstallSuccessful"
	ReasonInstallError        ConditionReason = "InstallError"
	ReasonUpgradeError        ConditionReason = "UpgradeError"
	ReasonReconcileError      ConditionReason = "ReconcileError"
	ReasonUninstallError      ConditionReason = "UninstallError"
)

// StackRoxRelease describes the Helm "release" that was most recently applied.
type StackRoxRelease struct {
	Version string `json:"version,omitempty"`
}

// AdditionalCA defines a certificate for an additional Certificate Authority.
type AdditionalCA struct {
	// Must be a valid file basename
	Name string `json:"name"`
	// PEM format
	Content string `json:"content"`
}

// TLSConfig defines common TLS-related settings for all components.
type TLSConfig struct {
	CASecret      *corev1.LocalObjectReference `json:"caSecret,omitempty"` // empty means: please autogenerate
	AdditionalCAs []AdditionalCA               `json:"additionalCAs,omitempty"`
}