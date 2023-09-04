package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"omdb-client/api/presenter"
	"omdb-client/entity"
	movies "omdb-client/usecase/movies"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// GetMovieByID handler
func GetMovieByID(service movies.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading Movies"
		vars := mux.Vars(r)
		ID := vars["id"]

		res, err := service.GetMovieByID(ID)

		if val, ok := entity.ErrCodeMapper[err]; ok {
			toJ := &presenter.AdditionalStatus{
				StatusCode:    val,
				StatusMessage: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(toJ)
			return
		}
		if err == entity.ErrUnauthorizedAccess {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(entity.ErrUnauthorizedAccess.Error()))
			return
		}

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Movie{
			Id:        res.Id,
			Title:     res.Title,
			Year:      res.Year,
			Rated:     res.Rated,
			Genre:     res.Genre,
			Plot:      res.Plot,
			Director:  res.Director,
			Actors:    res.Actors,
			Language:  res.Language,
			Country:   res.Country,
			Type:      res.Type,
			PosterUrl: res.PosterUrl,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(toJ)

	})
}

// SearchMovies handler
func SearchMovies(service movies.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading Movies"

		queryParams := r.URL.Query()

		// Access individual query parameters by their keys
		q := queryParams.Get("q")
		typeParam := queryParams.Get("type")
		page := queryParams.Get("page")

		// Convert the string to a uint64
		pageUint, err := strconv.ParseUint(page, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("page should be integer"))
			return
		}

		res, err := service.SearchMovies(q, typeParam, pageUint)

		if val, ok := entity.ErrCodeMapper[err]; ok {
			toJ := &presenter.AdditionalStatus{
				StatusCode:    val,
				StatusMessage: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(toJ)
			return
		}
		if err == entity.ErrUnauthorizedAccess {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(entity.ErrUnauthorizedAccess.Error()))
			return
		}

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)

	})
}

// MakeMoviesHandlers make url handlers
func MakeMoviesHandlers(r *mux.Router, n negroni.Negroni, service movies.UseCase) {
	r.Handle("/movies", n.With(
		negroni.Wrap(SearchMovies(service)),
	)).Methods("GET", "OPTIONS").Name("SearchMovies")

	r.Handle("/movies/{id}", n.With(
		negroni.Wrap(GetMovieByID(service)),
	)).Methods("GET", "OPTIONS").Name("GetMovieByID")

}
