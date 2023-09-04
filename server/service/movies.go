package token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	pb "omdb-server/pb-file"

	"golang.org/x/net/context"
)

// Server interface for our service methods
type Server struct {
	pb.UnimplementedOMDBServiceServer
}

type Movie struct {
	Title     string `json:"Title"`
	Year      string `json:"Year"`
	Rated     string `json:"Rated"`
	Genre     string `json:"Genre"`
	Plot      string `json:"Plot"`
	Director  string `json:"Director"`
	Actors    string `json:"Actors"`
	Language  string `json:"Language"`
	Country   string `json:"Country"`
	Type      string `json:"Type"`
	PosterUrl string `json:"PosterUrl"`
}

type MovieResult struct {
	Search       []SearchResult `json:"Search"`
	TotalResults string         `json:"totalResults"`
	Response     string         `json:"Response"`
}

type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func (s *Server) GetMovieByID(ctx context.Context, req *pb.GetMovieByIDRequest) (resp *pb.GetMovieByIDResponse, err error) {

	apiKey := os.Getenv("API_KEY")

	apiURL := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&i=%s", apiKey, req.Id)

	// Make an HTTP GET request to the API
	httpResp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer httpResp.Body.Close()

	// Check the response status code
	if httpResp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", httpResp.StatusCode)
		return
	}

	// Decode the JSON response into a Movie struct
	var movie Movie
	decoder := json.NewDecoder(httpResp.Body)
	if err = decoder.Decode(&movie); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	ActorArray := strings.Split(movie.Actors, ",")

	resp = &pb.GetMovieByIDResponse{
		Id:        req.Id,
		Title:     movie.Title,
		Year:      movie.Year,
		Rated:     movie.Rated,
		Genre:     movie.Genre,
		Plot:      movie.Plot,
		Director:  movie.Director,
		Actors:    ActorArray,
		Language:  movie.Language,
		Country:   movie.Country,
		Type:      movie.Type,
		PosterUrl: movie.PosterUrl,
	}

	return
}

func (s *Server) SearchMovies(ctx context.Context, req *pb.SearchMoviesRequest) (resp *pb.SearchMoviesResponse, err error) {

	apiKey := os.Getenv("API_KEY")
	apiURL := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&type=%s&s=%s&page=%d", apiKey, req.Type, req.Query, req.Page)

	// Make an HTTP GET request to the API
	httpResp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer httpResp.Body.Close()

	// Check the response status code
	if httpResp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", httpResp.StatusCode)
		return
	}

	// Decode the JSON response into a Movie struct
	var movieResult *MovieResult
	decoder := json.NewDecoder(httpResp.Body)
	if err = decoder.Decode(&movieResult); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	totalResults, err := strconv.ParseUint(movieResult.TotalResults, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := &pb.SearchMoviesResponse{
		TotalResults: totalResults,
	}

	for _, searchResult := range movieResult.Search {
		movie := &pb.MovieResult{
			Id:        searchResult.IMDbID,
			Title:     searchResult.Title,
			Year:      searchResult.Year,
			Type:      searchResult.Type,
			PosterUrl: searchResult.Poster,
		}
		response.Movies = append(response.Movies, movie)
	}

	return response, err
}
