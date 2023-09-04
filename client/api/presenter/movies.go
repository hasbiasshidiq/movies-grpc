package presenter

type Movie struct {
	Id        string   `json:"Id"`
	Title     string   `json:"Title"`
	Year      string   `json:"Year"`
	Rated     string   `json:"Rated"`
	Genre     string   `json:"Genre"`
	Plot      string   `json:"Plot"`
	Director  string   `json:"Director"`
	Actors    []string `json:"Actors"`
	Language  string   `json:"Language"`
	Country   string   `json:"Country"`
	Type      string   `json:"Type"`
	PosterUrl string   `json:"PosterUrl"`
}

type MovieResult struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Year      string `json:"year"`
	Type      string `json:"type"`
	PosterURL string `json:"poster_url"`
}

type SearchMoviesResponse struct {
	Movies       []MovieResult `json:"movies"`
	TotalResults uint64        `json:"total_results"`
}
