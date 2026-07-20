package services

type CloudProvider interface {
	CreateResource(name string) error
	ListResources() []string
	DeleteResource(name string) error
}
