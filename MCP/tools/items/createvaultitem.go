package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/1password-connect/mcp-server/config"
	"github.com/1password-connect/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatevaultitemHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		vaultUuidVal, ok := args["vaultUuid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: vaultUuid"), nil
		}
		vaultUuid, ok := vaultUuidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: vaultUuid"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.FullItem
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/vaults/%s/items", cfg.BaseURL, vaultUuid)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.FullItem
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatevaultitemTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_vaults_vaultUuid_items",
		mcp.WithDescription("Create a new Item"),
		mcp.WithString("vaultUuid", mcp.Required(), mcp.Description("The UUID of the Vault to create an Item in")),
		mcp.WithString("createdAt", mcp.Description("")),
		mcp.WithString("id", mcp.Description("")),
		mcp.WithObject("vault", mcp.Description("")),
		mcp.WithNumber("version", mcp.Description("")),
		mcp.WithString("lastEditedBy", mcp.Description("")),
		mcp.WithString("updatedAt", mcp.Description("")),
		mcp.WithString("category", mcp.Description("")),
		mcp.WithString("state", mcp.Description("")),
		mcp.WithArray("tags", mcp.Description("")),
		mcp.WithString("title", mcp.Description("")),
		mcp.WithBoolean("favorite", mcp.Description("")),
		mcp.WithArray("urls", mcp.Description("")),
		mcp.WithArray("fields", mcp.Description("")),
		mcp.WithArray("files", mcp.Description("")),
		mcp.WithArray("sections", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatevaultitemHandler(cfg),
	}
}
