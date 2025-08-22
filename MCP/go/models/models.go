package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}

// SpellCheck represents the SpellCheck schema from the OpenAPI specification
type SpellCheck struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// SpellingTokenSuggestion represents the SpellingTokenSuggestion schema from the OpenAPI specification
type SpellingTokenSuggestion struct {
	Pingurlsuffix string `json:"pingUrlSuffix,omitempty"`
	Score float64 `json:"score,omitempty"`
	Suggestion string `json:"suggestion"`
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// SpellingFlaggedToken represents the SpellingFlaggedToken schema from the OpenAPI specification
type SpellingFlaggedToken struct {
	Token string `json:"token"`
	TypeField string `json:"type"`
	Offset int `json:"offset"`
	Pingurlsuffix string `json:"pingUrlSuffix,omitempty"`
	Suggestions []SpellingTokenSuggestion `json:"suggestions,omitempty"`
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}
