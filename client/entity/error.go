package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

// ErrApplAlreadyExist application already exists in the database
var ErrAlreadyExist = errors.New("already exists")

// ErrInvalidEntity invalid entity
var ErrExpirateToken = errors.New("token is expirate")

// ErrInvalidTokenEntity invalid token
var ErrInvalidToken = errors.New("invalid token")

// ErrInvalidExpirationTime invalid expiation time / expiration time before now
var ErrInvalidExpirationTime = errors.New("invalid expiration time")

// ErrInvalidCredentials invalid login credentials
var ErrInvalidCredentials = errors.New("invalid credentials")

// ErrUnauthorizedAccess invalid entity
var ErrUnauthorizedAccess = errors.New("unauthorized access")

var ErrCodeMapper = map[error]int{
	ErrNotFound:              10,
	ErrInvalidEntity:         11,
	ErrAlreadyExist:          12,
	ErrInvalidToken:          20,
	ErrExpirateToken:         21,
	ErrInvalidExpirationTime: 30,
	ErrInvalidCredentials:    40,
}
