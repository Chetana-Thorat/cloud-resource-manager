package services

type MockKubernetesProvider struct {
	pods        []string
	deployments []string
}

func NewMockKubernetesProvider() *MockKubernetesProvider {
	return &MockKubernetesProvider{
		pods:        []string{"pod-1", "pod-2"},
		deployments: []string{"api-service", "worker-service"},
	}
}

func (m *MockKubernetesProvider) GetPods() []string {
	return m.pods
}

func (m *MockKubernetesProvider) GetDeployments() []string {
	return m.deployments
}

func (m *MockKubernetesProvider) RolloutStatus(name string) string {
	for _, d := range m.deployments {
		if d == name {
			return "rollout complete"
		}
	}
	return "deployment not found"
}
