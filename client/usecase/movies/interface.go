package contributor

import entity "omdb-client/entity"

// GRPCDriver interface
type GRPCDriver interface {
	GetMovieByID(Id string) (movie *entity.Movie, err error)
	SearchMovies(q string, typeParam string, page uint64) (movie *entity.SearchMoviesResponse, err error)
}

// UseCase interface
type UseCase interface {
	GetMovieByID(Id string) (movie *entity.Movie, err error)
	SearchMovies(q string, typeParam string, page uint64) (movie *entity.SearchMoviesResponse, err error)
}
