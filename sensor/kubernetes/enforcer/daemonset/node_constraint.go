package daemonset

import (
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/retry"
	"github.com/stackrox/rox/sensor/kubernetes/enforcer/common"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// EnforceNodeConstraint reschedules the daemon set with unsatisfiable constraints.
func EnforceNodeConstraint(client *kubernetes.Clientset, deploymentInfo *central.DeploymentEnforcement) (err error) {
	// Load the current DaemonSet for the deployment.
	var ds *v1beta1.DaemonSet
	ds, err = client.ExtensionsV1beta1().DaemonSets(deploymentInfo.GetNamespace()).Get(deploymentInfo.GetDeploymentName(), metav1.GetOptions{})
	if err != nil {
		return retry.MakeRetryable(err)
	}

	// Apply the constraint modification.
	err = common.ApplyNodeConstraintToObj(ds, deploymentInfo.GetAlertId())
	if err != nil {
		return
	}

	// Post the new DaemonSet data.
	_, err = client.ExtensionsV1beta1().DaemonSets(deploymentInfo.GetNamespace()).Update(ds)
	if err != nil {
		return retry.MakeRetryable(err)
	}
	return
}
