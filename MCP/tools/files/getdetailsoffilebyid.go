package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/1password-connect/mcp-server/config"
	"github.com/1password-connect/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetdetailsoffilebyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		fileUuidVal, ok := args["fileUuid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fileUuid"), nil
		}
		fileUuid, ok := fileUuidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fileUuid"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["inline_files"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("inline_files=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/vaults/%s/items/%s/files/%s%s", cfg.BaseURL, vaultUuid, itemUuid, fileUuid, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.File
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

func CreateGetdetailsoffilebyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_vaults_vaultUuid_items_itemUuid_files_fileUuid",
		mcp.WithDescription("Get the details of a File"),
		mcp.WithString("vaultUuid", mcp.Required(), mcp.Description("The UUID of the Vault to fetch Item from")),
		mcp.WithString("itemUuid", mcp.Required(), mcp.Description("The UUID of the Item to fetch File from")),
		mcp.WithString("fileUuid", mcp.Required(), mcp.Description("The UUID of the File to fetch")),
		mcp.WithBoolean("inline_files", mcp.Description("Tells server to return the base64-encoded file contents in the response.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetdetailsoffilebyidHandler(cfg),
	}
}
