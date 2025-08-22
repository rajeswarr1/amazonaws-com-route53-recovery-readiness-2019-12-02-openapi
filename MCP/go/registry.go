package main

import (
	"github.com/aws-route53-recovery-readiness/mcp-server/config"
	"github.com/aws-route53-recovery-readiness/mcp-server/models"
	tools_cells "github.com/aws-route53-recovery-readiness/mcp-server/tools/cells"
	tools_readinesschecks "github.com/aws-route53-recovery-readiness/mcp-server/tools/readinesschecks"
	tools_recoverygroups "github.com/aws-route53-recovery-readiness/mcp-server/tools/recoverygroups"
	tools_rules "github.com/aws-route53-recovery-readiness/mcp-server/tools/rules"
	tools_crossaccountauthorizations "github.com/aws-route53-recovery-readiness/mcp-server/tools/crossaccountauthorizations"
	tools_resourcesets "github.com/aws-route53-recovery-readiness/mcp-server/tools/resourcesets"
	tools_tags "github.com/aws-route53-recovery-readiness/mcp-server/tools/tags"
	tools_cellreadiness "github.com/aws-route53-recovery-readiness/mcp-server/tools/cellreadiness"
	tools_recoverygroupreadiness "github.com/aws-route53-recovery-readiness/mcp-server/tools/recoverygroupreadiness"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_cells.CreateListcellsTool(cfg),
		tools_cells.CreateCreatecellTool(cfg),
		tools_readinesschecks.CreateGetreadinesscheckresourcestatusTool(cfg),
		tools_readinesschecks.CreateListreadinesschecksTool(cfg),
		tools_readinesschecks.CreateCreatereadinesscheckTool(cfg),
		tools_recoverygroups.CreateGetarchitecturerecommendationsTool(cfg),
		tools_rules.CreateListrulesTool(cfg),
		tools_recoverygroups.CreateListrecoverygroupsTool(cfg),
		tools_recoverygroups.CreateCreaterecoverygroupTool(cfg),
		tools_crossaccountauthorizations.CreateListcrossaccountauthorizationsTool(cfg),
		tools_crossaccountauthorizations.CreateCreatecrossaccountauthorizationTool(cfg),
		tools_readinesschecks.CreateGetreadinesscheckstatusTool(cfg),
		tools_resourcesets.CreateCreateresourcesetTool(cfg),
		tools_resourcesets.CreateListresourcesetsTool(cfg),
		tools_resourcesets.CreateDeleteresourcesetTool(cfg),
		tools_resourcesets.CreateGetresourcesetTool(cfg),
		tools_resourcesets.CreateUpdateresourcesetTool(cfg),
		tools_tags.CreateListtagsforresourcesTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_cellreadiness.CreateGetcellreadinesssummaryTool(cfg),
		tools_recoverygroupreadiness.CreateGetrecoverygroupreadinesssummaryTool(cfg),
		tools_cells.CreateGetcellTool(cfg),
		tools_cells.CreateUpdatecellTool(cfg),
		tools_cells.CreateDeletecellTool(cfg),
		tools_crossaccountauthorizations.CreateDeletecrossaccountauthorizationTool(cfg),
		tools_readinesschecks.CreateDeletereadinesscheckTool(cfg),
		tools_readinesschecks.CreateGetreadinesscheckTool(cfg),
		tools_readinesschecks.CreateUpdatereadinesscheckTool(cfg),
		tools_recoverygroups.CreateDeleterecoverygroupTool(cfg),
		tools_recoverygroups.CreateGetrecoverygroupTool(cfg),
		tools_recoverygroups.CreateUpdaterecoverygroupTool(cfg),
	}
}
