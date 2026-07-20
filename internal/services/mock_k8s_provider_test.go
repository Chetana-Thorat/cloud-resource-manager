package services

import "testing"

func TestMockKubernetesProvider(t *testing.T) {
	p := NewMockKubernetesProvider()

	pods := p.GetPods()
	if len(pods) == 0 {
		t.Fatalf("expected pods to be returned")
	}

	deployments := p.GetDeployments()
	if len(deployments) == 0 {
		t.Fatalf("expected deployments to be returned")
	}

	if status := p.RolloutStatus("api-service"); status != "rollout complete" {
		t.Fatalf("expected rollout complete, got %s", status)
	}

	if status := p.RolloutStatus("missing-service"); status != "deployment not found" {
		t.Fatalf("expected deployment not found, got %s", status)
	}
}
