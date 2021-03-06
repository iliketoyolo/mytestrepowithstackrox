package check433

const interpretationText = `StackRox has visibility into the effects of Kubernetes Network Policies on your deployments.
Network Policies restrict inbound and outbound traffic. Therefore, if every deployment has both inbound
(ingress) and outbound (egress) Network Policies that apply to it, it can be considered compliant, so
long as none of those deployments are using the host namespace, which allows circumvention of Network Policies.`
