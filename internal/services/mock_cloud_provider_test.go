package services

import "testing"

func TestMockCloudProvider(t *testing.T) {
	p := NewMockCloudProvider()

	if err := p.CreateResource("demo"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	resources := p.ListResources()
	if len(resources) != 1 || resources[0] != "demo" {
		t.Fatalf("expected [demo], got %v", resources)
	}

	if err := p.DeleteResource("demo"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(p.ListResources()) != 0 {
		t.Fatalf("expected no resources after delete")
	}
}

func TestMockCloudProvider_CreateEmptyName(t *testing.T) {
	p := NewMockCloudProvider()

	if err := p.CreateResource(""); err == nil {
		t.Fatalf("expected error for empty resource name")
	}
}

func TestMockCloudProvider_DeleteMissingResource(t *testing.T) {
	p := NewMockCloudProvider()

	if err := p.DeleteResource("missing"); err == nil {
		t.Fatalf("expected error for missing resource")
	}
}
