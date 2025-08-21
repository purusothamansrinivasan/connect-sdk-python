package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// File represents the File schema from the OpenAPI specification
type File struct {
	Size int `json:"size,omitempty"` // Size in bytes of the file
	Content string `json:"content,omitempty"` // Base64-encoded contents of the file. Only set if size <= OP_MAX_INLINE_FILE_SIZE_KB kb and `inline_files` is set to `true`.
	Content_path string `json:"content_path,omitempty"` // Path of the Connect API that can be used to download the contents of this file.
	Id string `json:"id,omitempty"` // ID of the file
	Name string `json:"name,omitempty"` // Name of the file
	Section map[string]interface{} `json:"section,omitempty"` // For files that are in a section, this field describes the section.
}

// APIRequest represents the APIRequest schema from the OpenAPI specification
type APIRequest struct {
	Requestid string `json:"requestId,omitempty"` // The unique id used to identify a single request.
	Resource map[string]interface{} `json:"resource,omitempty"`
	Result string `json:"result,omitempty"`
	Timestamp string `json:"timestamp,omitempty"` // The time at which the request was processed by the server.
	Action string `json:"action,omitempty"`
	Actor map[string]interface{} `json:"actor,omitempty"`
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Message string `json:"message,omitempty"` // A message detailing the error
	Status int `json:"status,omitempty"` // HTTP Status Code
}

// FullItem represents the FullItem schema from the OpenAPI specification
type FullItem struct {
	Updatedat string `json:"updatedAt,omitempty"`
	Category string `json:"category"`
	State string `json:"state,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Title string `json:"title,omitempty"`
	Favorite bool `json:"favorite,omitempty"`
	Urls []map[string]interface{} `json:"urls,omitempty"`
	Createdat string `json:"createdAt,omitempty"`
	Id string `json:"id,omitempty"`
	Vault map[string]interface{} `json:"vault"`
	Version int `json:"version,omitempty"`
	Lasteditedby string `json:"lastEditedBy,omitempty"`
	Sections []map[string]interface{} `json:"sections,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	Files []File `json:"files,omitempty"`
}

// Item represents the Item schema from the OpenAPI specification
type Item struct {
	Createdat string `json:"createdAt,omitempty"`
	Id string `json:"id,omitempty"`
	Vault map[string]interface{} `json:"vault"`
	Version int `json:"version,omitempty"`
	Lasteditedby string `json:"lastEditedBy,omitempty"`
	Updatedat string `json:"updatedAt,omitempty"`
	Category string `json:"category"`
	State string `json:"state,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Title string `json:"title,omitempty"`
	Favorite bool `json:"favorite,omitempty"`
	Urls []map[string]interface{} `json:"urls,omitempty"`
}

// Vault represents the Vault schema from the OpenAPI specification
type Vault struct {
	Description string `json:"description,omitempty"`
	Items int `json:"items,omitempty"` // Number of active items in the vault
	Updatedat string `json:"updatedAt,omitempty"`
	Attributeversion int `json:"attributeVersion,omitempty"` // The vault version
	Createdat string `json:"createdAt,omitempty"`
	Id string `json:"id,omitempty"`
	Contentversion int `json:"contentVersion,omitempty"` // The version of the vault contents
	Name string `json:"name,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// GeneratorRecipe represents the GeneratorRecipe schema from the OpenAPI specification
type GeneratorRecipe struct {
	Charactersets []string `json:"characterSets,omitempty"`
	Excludecharacters string `json:"excludeCharacters,omitempty"` // List of all characters that should be excluded from generated passwords.
	Length int `json:"length,omitempty"` // Length of the generated value
}

// ServiceDependency represents the ServiceDependency schema from the OpenAPI specification
type ServiceDependency struct {
	Message string `json:"message,omitempty"` // Human-readable message for explaining the current state.
	Service string `json:"service,omitempty"`
	Status string `json:"status,omitempty"`
}

// Field represents the Field schema from the OpenAPI specification
type Field struct {
	Purpose string `json:"purpose,omitempty"` // Some item types, Login and Password, have fields used for autofill. This property indicates that purpose and is required for some item types.
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
	TypeField string `json:"type"`
	Generate bool `json:"generate,omitempty"` // If value is not present then a new value should be generated for this field
	Recipe GeneratorRecipe `json:"recipe,omitempty"` // The recipe is used in conjunction with the "generate" property to set the character set used to generate a new secure value
	Section map[string]interface{} `json:"section,omitempty"`
	Entropy float64 `json:"entropy,omitempty"` // For fields with a purpose of `PASSWORD` this is the entropy of the value
	Id string `json:"id"`
}
