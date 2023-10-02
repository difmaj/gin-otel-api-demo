package todo

import (
	"context"
	"testing"

	"go.uber.org/goleak"

	"github.com/bluele/gcache"
	"github.com/stretchr/testify/suite"
)

type CacheRepositorySuite struct {
	suite.Suite
	cache gcache.Cache
	ctx   context.Context
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func (s *CacheRepositorySuite) SetupSuite() {
	s.cache = gcache.New(100).LRU().Build()
}

func (s *CacheRepositorySuite) TearDownSuite() {
	s.cache.Purge()
}
