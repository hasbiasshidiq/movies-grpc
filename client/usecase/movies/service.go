package contributor

import (
	entity "omdb-client/entity"
)

// Service for Token usecase
type Service struct {
	grpc GRPCDriver
}

// NewService create new service
func NewService(g GRPCDriver) *Service {
	return &Service{
		grpc: g,
	}
}

// GetMovieByID
func (s *Service) GetMovieByID(Id string) (*entity.Movie, error) {

	movie, err := s.grpc.GetMovieByID(Id)

	return movie, err
}

// SearchMovies
func (s *Service) SearchMovies(q string, typeParam string, page uint64) (movie *entity.SearchMoviesResponse, err error) {

	movies, err := s.grpc.SearchMovies(q, typeParam, page)

	return movies, err
}
