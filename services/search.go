package services

import (
	"github.com/fpay/gopress"
)

const (
	// SearchServiceName is the identity of search service
	SearchServiceName = "search"
)

// SearchService type
type SearchService struct {
	// Uncomment this line if this service has dependence on other services in the container
	// c *gopress.Container
}

// NewSearchService returns instance of search service
func NewSearchService() *SearchService {
	return new(SearchService)
}

// ServiceName is used to implements gopress.Service
func (s *SearchService) ServiceName() string {
	return SearchServiceName
}

// RegisterContainer is used to implements gopress.Service
func (s *SearchService) RegisterContainer(c *gopress.Container) {
	// Uncomment this line if this service has dependence on other services in the container
	// s.c = c
}

func (s *SearchService) SampleMethod() {
}
