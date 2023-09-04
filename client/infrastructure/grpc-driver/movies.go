package grpcdriver

import (
	"context"
	"log"
	"time"

	pb_client "omdb-client/infrastructure/grpc-driver/pb-file"

	"omdb-client/config"
	entity "omdb-client/entity"

	"google.golang.org/grpc"
)

// Token grpc client
type omdbGRPC struct{}

// NewOmdbGRPC create new repository
func NewOmdbGRPC() *omdbGRPC {
	return &omdbGRPC{}
}

func (r *omdbGRPC) GetMovieByID(Id string) (movie *entity.Movie, err error) {

	conn, err := grpc.Dial(config.GRPC_URL, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	a := pb_client.NewOMDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := a.GetMovieByID(ctx, &pb_client.GetMovieByIDRequest{
		Id: Id,
	})
	if err != nil {
		log.Printf("Could not get movie :%v", err)
		return
	}

	result := &entity.Movie{
		Id:        resp.Id,
		Title:     resp.Title,
		Year:      resp.Year,
		Rated:     resp.Rated,
		Genre:     resp.Genre,
		Plot:      resp.Plot,
		Director:  resp.Director,
		Actors:    resp.Actors,
		Language:  resp.Language,
		Country:   resp.Country,
		Type:      resp.Type,
		PosterUrl: resp.PosterUrl,
	}

	log.Printf("Successfully Get Movie By ID")

	return result, err
}

func (r *omdbGRPC) SearchMovies(q string, typeParam string, page uint64) (movie *entity.SearchMoviesResponse, err error) {

	conn, err := grpc.Dial(config.GRPC_URL, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	a := pb_client.NewOMDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := a.SearchMovies(ctx, &pb_client.SearchMoviesRequest{
		Query: q,
		Type:  typeParam,
		Page:  page,
	})
	if err != nil {
		log.Printf("Could not search movie :%v", err)
		return
	}

	result := &entity.SearchMoviesResponse{
		TotalResults: resp.TotalResults,
	}

	for _, searchResult := range resp.Movies {
		movie := &entity.MovieResult{
			Id:        searchResult.Id,
			Title:     searchResult.Title,
			Year:      searchResult.Year,
			Type:      searchResult.Type,
			PosterURL: searchResult.PosterUrl,
		}
		result.Movies = append(result.Movies, movie)
	}

	log.Printf("Successfully Search Movies")

	return result, err
}
