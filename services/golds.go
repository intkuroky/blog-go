package services

import (
	"github.com/fpay/gopress"
)

const (
	// GoldsServiceName is the identity of golds service
	GoldsServiceName = "golds"
)

// GoldsService type
type GoldsService struct {
	// Uncomment this line if this service has dependence on other services in the container
	// c *gopress.Container
}

// NewGoldsService returns instance of golds service
func NewGoldsService() *GoldsService {
	return new(GoldsService)
}

// ServiceName is used to implements gopress.Service
func (s *GoldsService) ServiceName() string {
	return GoldsServiceName
}

// RegisterContainer is used to implements gopress.Service
func (s *GoldsService) RegisterContainer(c *gopress.Container) {
	// Uncomment this line if this service has dependence on other services in the container
	// s.c = c
}

func (s *GoldsService) SampleMethod() {
}
