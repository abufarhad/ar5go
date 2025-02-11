// Package classification ar5go system API.
//
// the purpose of this service is to provide & store all user of a company and their location histories
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /api
//     Version: 1.0.0
//     License: None
//     Contact: Rezwanul-Haque<rezwanul.cse@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - base64
//
//     SecurityDefinitions:
//     base64:
//          type: apiKey
//          name: ar5go-app-key
//          in: header
// swagger:meta
package openapi

import (
	"ar5go/infra/errors"
)

// Generic error message
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body errors.RestErr `json:"error_response"`
}

type genericSuccessResponse struct {
	Message string `json:"message"`
}

// returns a message
// swagger:response genericSuccessResponse
type genericSuccessResponseWrapper struct {
	// in: body
	genericSuccessResponse `json:"message"`
}

// Payload for company
// swagger:parameters CreateCompany
type companyPayloadWrapper struct {
	// in:body
	Body interface{}
}

// List all the users of a company
// swagger:parameters UserQueryParameters
type usersQueryParametersWrapper struct {
	// in:query
	//example: 10
	Size int64 `json:"size"`
	//example: 2
	Page int64 `json:"page"`
	//example: created_at desc
	Sort string `json:"sort"`
	//example: rezwan
	QS string `json:"qs"`
	//example: user_id.(in, contains, equals, gt, gte, lt, lte)
	ColumnOperation string `json:"column:operation"`
}

// List all the users
// swagger:response UserResponse
type usersRespWrapper struct {
	// in:body
	Body interface{}
}
