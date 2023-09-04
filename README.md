## Setup Server 

### move into server directory
```
cd server
```
### build server
```
go build .
```
### run server
```
./omdb-server
```


## Setup Client

### move into server directory
```
cd client
```
### build client
```
go build api/main.go
```
### run client
```
./main
```

## Example request for GetMovieByID
### Curl Request
```
curl --location 'http://127.0.0.1:8000/movies/tt4853102'
```
### Response
```
{
    "Id": "tt4853102",
    "Title": "Batman: The Killing Joke",
    "Year": "2016",
    "Rated": "R",
    "Genre": "Animation, Action, Crime",
    "Plot": "As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.",
    "Director": "Sam Liu",
    "Actors": [
        "Kevin Conroy",
        " Mark Hamill",
        " Tara Strong"
    ],
    "Language": "English",
    "Country": "United States",
    "Type": "movie",
    "PosterUrl": ""
}
```

## Example request for SearchMovies
### Curl Request
```
curl --location 'http://127.0.0.1:8000/movies?q=Batman&page=2'
```
### Response
```
{
    "movies": [
        {
            "id": "tt4853102",
            "title": "Batman: The Killing Joke",
            "year": "2016",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
        },
        {
            "id": "tt1569923",
            "title": "Batman: Under the Red Hood",
            "year": "2010",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
        },
        {
            "id": "tt2313197",
            "title": "Batman: The Dark Knight Returns, Part 1",
            "year": "2012",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"
        },
        {
            "id": "tt0106364",
            "title": "Batman: Mask of the Phantasm",
            "year": "1993",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BYTRiMWM3MGItNjAxZC00M2E3LThhODgtM2QwOGNmZGU4OWZhXkEyXkFqcGdeQXVyNjExODE1MDc@._V1_SX300.jpg"
        },
        {
            "id": "tt2166834",
            "title": "Batman: The Dark Knight Returns, Part 2",
            "year": "2013",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BYTEzMmE0ZDYtYWNmYi00ZWM4LWJjOTUtYTE0ZmQyYWM3ZjA0XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg"
        },
        {
            "id": "tt0060153",
            "title": "Batman: The Movie",
            "year": "1966",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BMmM1OGIzM2UtNThhZS00ZGNlLWI4NzEtZjlhOTNhNmYxZGQ0XkEyXkFqcGdeQXVyNTkxMzEwMzU@._V1_SX300.jpg"
        },
        {
            "id": "tt1672723",
            "title": "Batman: Year One",
            "year": "2011",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BNTJjMmVkZjctNjNjMS00ZmI2LTlmYWEtOWNiYmQxYjY0YWVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
        },
        {
            "id": "tt3139086",
            "title": "Batman: Assault on Arkham",
            "year": "2014",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BZDU1ZGRiY2YtYmZjMi00ZDQwLWJjMWMtNzUwNDMwYjQ4ZTVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
        },
        {
            "id": "tt0147746",
            "title": "Batman Beyond",
            "year": "1999–2001",
            "type": "series",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BZWJhNjA4YTAtODBlMS00NjIzLThhZWUtOGYxMGM3OTRmNDZmXkEyXkFqcGdeQXVyNjk1Njg5NTA@._V1_SX300.jpg"
        },
        {
            "id": "tt1117563",
            "title": "Batman: Gotham Knight",
            "year": "2008",
            "type": "movie",
            "poster_url": "https://m.media-amazon.com/images/M/MV5BM2I0YTFjOTUtMWYzNC00ZTgyLTk2NWEtMmE3N2VlYjEwN2JlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
        }
    ],
    "total_results": 544
}
```
