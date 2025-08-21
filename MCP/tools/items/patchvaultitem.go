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

func PatchvaultitemHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		itemUuidVal, ok := args["itemUuid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: itemUuid"), nil
		}
		itemUuid, ok := itemUuidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: itemUuid"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody []map[string]interface{}
		
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
		url := fmt.Sprintf("%s/vaults/%s/items/%s", cfg.BaseURL, vaultUuid, itemUuid)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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

func CreatePatchvaultitemTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_vaults_vaultUuid_items_itemUuid",
		mcp.WithDescription("Update a subset of Item attributes"),
		mcp.WithString("vaultUuid", mcp.Required(), mcp.Description("The UUID of the Vault the item is in")),
		mcp.WithString("itemUuid", mcp.Required(), mcp.Description("The UUID of the Item to update")),
		mcp.WithArray("items", mcp.Required(), mcp.Description("Array of objects")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PatchvaultitemHandler(cfg),
	}
}
