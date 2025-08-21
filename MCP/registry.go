package main

import (
	"github.com/1password-connect/mcp-server/config"
	"github.com/1password-connect/mcp-server/models"
	tools_items "github.com/1password-connect/mcp-server/tools/items"
	tools_metrics "github.com/1password-connect/mcp-server/tools/metrics"
	tools_vaults "github.com/1password-connect/mcp-server/tools/vaults"
	tools_files "github.com/1password-connect/mcp-server/tools/files"
	tools_activity "github.com/1password-connect/mcp-server/tools/activity"
	tools_health "github.com/1password-connect/mcp-server/tools/health"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_items.CreateDeletevaultitemTool(cfg),
		tools_items.CreateGetvaultitembyidTool(cfg),
		tools_items.CreatePatchvaultitemTool(cfg),
		tools_items.CreateUpdatevaultitemTool(cfg),
		tools_metrics.CreateGetprometheusmetricsTool(cfg),
		tools_vaults.CreateGetvaultsTool(cfg),
		tools_files.CreateGetitemfilesTool(cfg),
		tools_vaults.CreateGetvaultbyidTool(cfg),
		tools_items.CreateGetvaultitemsTool(cfg),
		tools_items.CreateCreatevaultitemTool(cfg),
		tools_files.CreateGetdetailsoffilebyidTool(cfg),
		tools_activity.CreateGetapiactivityTool(cfg),
		tools_health.CreateGetserverhealthTool(cfg),
		tools_health.CreateGetheartbeatTool(cfg),
	}
}
