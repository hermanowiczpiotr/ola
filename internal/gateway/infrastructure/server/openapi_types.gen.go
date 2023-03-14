// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package server

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Token string `json:"token"`
}

// RegisterRequest defines model for RegisterRequest.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse defines model for RegisterResponse.
type RegisterResponse struct {
	Email string `json:"email"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LoginRequest

// RegisterUserJSONRequestBody defines body for RegisterUser for application/json ContentType.
type RegisterUserJSONRequestBody = RegisterRequest
