package services

type KubernetesProvider interface {
	GetPods() []string
	GetDeployments() []string
	RolloutStatus(name string) string
}
