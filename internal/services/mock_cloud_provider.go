package services

import "fmt"

type MockCloudProvider struct {
	resources map[string]bool
}

func NewMockCloudProvider() *MockCloudProvider {
	return &MockCloudProvider{
		resources: make(map[string]bool),
	}
}

func (m *MockCloudProvider) CreateResource(name string) error {
	if name == "" {
		return fmt.Errorf("resource name cannot be empty")
	}
	m.resources[name] = true
	return nil
}

func (m *MockCloudProvider) ListResources() []string {
	var result []string
	for name := range m.resources {
		result = append(result, name)
	}
	return result
}

func (m *MockCloudProvider) DeleteResource(name string) error {
	if _, exists := m.resources[name]; !exists {
		return fmt.Errorf("resource not found: %s", name)
	}
	delete(m.resources, name)
	return nil
}
