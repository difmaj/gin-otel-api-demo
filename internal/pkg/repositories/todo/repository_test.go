package todo

import (
	"context"
	"testing"

	"github.com/difmaj/gin-otel-api-demo/internal/pkg/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ToDoCacheRepositorySuite struct {
	CacheRepositorySuite
}

func TestToDoCacheRepositorySuite(t *testing.T) {
	suite.Run(t, new(ToDoCacheRepositorySuite))
}

func (s *ToDoCacheRepositorySuite) SetupTest() {
	s.ctx = context.TODO()

	// Add fixtures
	s.cache.Purge()
	s.cache.Set(
		uuid.MustParse("5572043c-6c66-4b9b-86d4-1fb6052a90a8"),
		[]byte(`{"id":"5572043c-6c66-4b9b-86d4-1fb6052a90a8","title":"Test","completed":false}`),
	)
	s.cache.Set(
		uuid.MustParse("8d093e54-373f-40eb-8c63-c2975170cef6"),
		[]byte(`{"id":"8d093e54-373f-40eb-8c63-c2975170cef6","title":"Test Number 2","completed":true}`),
	)
	s.cache.Set(
		uuid.MustParse("071afcbf-feaf-40bc-91ae-871f8e2875c9"),
		[]byte(`{"id":"071afcbf-feaf-40bc-91ae-871f8e2875c9","title":"Another Test","completed":false}`),
	)
}

func (s *ToDoCacheRepositorySuite) TestCreate() {
	s.Run("success", func() {
		repo := NewCacheRepository(s.cache)

		request := &entities.CreateToDoRequest{
			Title:     "Test123",
			Completed: false,
		}
		response, err := repo.Create(s.ctx, request)
		s.Require().NoError(err)
		s.Require().NotNil(response)
		s.Require().NotNil(response.Data)
		s.Require().NotEmpty(response.Data.ID)
		s.Require().Equal(request.Title, response.Data.Title)
		s.Require().False(response.Data.Completed)
	})
	s.Run("failure", func() {
		repo := NewCacheRepository(s.cache)

		request := &entities.CreateToDoRequest{}
		response, err := repo.Create(s.ctx, request)
		s.Require().Error(err)
		s.Require().Nil(response)
		s.Require().ErrorIs(entities.ErrMissingToDoTitle, err)
	})
}

func (s *ToDoCacheRepositorySuite) TestGet() {
	s.Run("success", func() {
		repo := NewCacheRepository(s.cache)
		responseGet, err := repo.Get(s.ctx, &entities.GetToDoRequest{
			ID: uuid.MustParse("5572043c-6c66-4b9b-86d4-1fb6052a90a8"),
		})
		s.Require().NoError(err)
		s.Require().NotNil(responseGet)
		s.Require().NotNil(responseGet.Data)
		s.Require().False(responseGet.Data.Completed)
	})

	s.Run("failure", func() {
		repo := NewCacheRepository(s.cache)
		response, err := repo.Get(s.ctx, &entities.GetToDoRequest{
			ID: uuid.New(),
		})

		s.Require().Error(err)
		s.Require().Nil(response)
	})
}

func (s *ToDoCacheRepositorySuite) TestList() {
	s.Run("success", func() {
		repo := NewCacheRepository(s.cache)

		response, err := repo.List(s.ctx, nil)
		s.Require().NoError(err)
		s.Require().NotNil(response)
		s.Require().NotNil(response.Data)
		s.Require().Len(response.Data, 3)
	})

	s.Run("success_with_completed_filter", func() {
		repo := NewCacheRepository(s.cache)

		request := &entities.ListToDoRequest{
			Filters: &entities.ListToDoFilters{
				Completed: new(bool),
			},
		}

		response, err := repo.List(s.ctx, request)
		s.Require().NoError(err)
		s.Require().NotNil(response)
		s.Require().NotNil(response.Data)
		s.Require().Len(response.Data, 2)
	})

	s.Run("failure", func() {
		s.cache.Purge()
		repo := NewCacheRepository(s.cache)

		_, err := repo.List(s.ctx, nil)
		s.Require().NoError(err)
	})
}

func (s *ToDoCacheRepositorySuite) TestUpdate() {
	s.Run("success", func() {
		repo := NewCacheRepository(s.cache)

		responseUpdate, err := repo.Update(s.ctx, &entities.UpdateToDoRequest{
			ID:        uuid.MustParse("5572043c-6c66-4b9b-86d4-1fb6052a90a8"),
			Title:     "Test1234",
			Completed: true,
		})
		s.Require().NoError(err)
		s.Require().NotNil(responseUpdate)
		s.Require().NotNil(responseUpdate.Data)
		s.Require().Equal("Test1234", responseUpdate.Data.Title)
		s.Require().True(responseUpdate.Data.Completed)
	})

	s.Run("failure", func() {
		repo := NewCacheRepository(s.cache)

		response, err := repo.Update(s.ctx, &entities.UpdateToDoRequest{
			ID:        uuid.New(),
			Title:     "Test1234",
			Completed: true,
		})
		s.Require().Error(err)
		s.Require().Nil(response)
	})
}

func (s *ToDoCacheRepositorySuite) TestDelete() {
	s.Run("success", func() {
		repo := NewCacheRepository(s.cache)
		responseDelete, err := repo.Delete(s.ctx, &entities.DeleteToDoRequest{
			ID: uuid.MustParse("5572043c-6c66-4b9b-86d4-1fb6052a90a8"),
		})
		s.Require().NoError(err)
		s.Require().NotNil(responseDelete)
		s.Require().True(responseDelete.Success)
	})

	s.Run("failure", func() {
		repo := NewCacheRepository(s.cache)

		response, err := repo.Delete(s.ctx, &entities.DeleteToDoRequest{
			ID: uuid.New(),
		})
		s.Require().Error(err)
		s.Require().Nil(response)
	})
}
