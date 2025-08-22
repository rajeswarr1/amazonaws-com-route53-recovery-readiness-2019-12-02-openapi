package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetResourceSetRequest represents the GetResourceSetRequest schema from the OpenAPI specification
type GetResourceSetRequest struct {
}

// GetReadinessCheckStatusResponse represents the GetReadinessCheckStatusResponse schema from the OpenAPI specification
type GetReadinessCheckStatusResponse struct {
	Messages interface{} `json:"Messages,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Readiness interface{} `json:"Readiness,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
}

// GetResourceSetResponse represents the GetResourceSetResponse schema from the OpenAPI specification
type GetResourceSetResponse struct {
	Resourcesetarn interface{} `json:"ResourceSetArn,omitempty"`
	Resourcesetname interface{} `json:"ResourceSetName,omitempty"`
	Resourcesettype interface{} `json:"ResourceSetType,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// ResourceResult represents the ResourceResult schema from the OpenAPI specification
type ResourceResult struct {
	Componentid interface{} `json:"ComponentId,omitempty"`
	Lastcheckedtimestamp interface{} `json:"LastCheckedTimestamp"`
	Readiness interface{} `json:"Readiness"`
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
}

// UpdateReadinessCheckRequest represents the UpdateReadinessCheckRequest schema from the OpenAPI specification
type UpdateReadinessCheckRequest struct {
	Resourcesetname interface{} `json:"ResourceSetName"`
}

// ListRulesRequest represents the ListRulesRequest schema from the OpenAPI specification
type ListRulesRequest struct {
}

// R53ResourceRecord represents the R53ResourceRecord schema from the OpenAPI specification
type R53ResourceRecord struct {
	Domainname interface{} `json:"DomainName,omitempty"`
	Recordsetid interface{} `json:"RecordSetId,omitempty"`
}

// UpdateRecoveryGroupResponse represents the UpdateRecoveryGroupResponse schema from the OpenAPI specification
type UpdateRecoveryGroupResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
	Recoverygrouparn interface{} `json:"RecoveryGroupArn,omitempty"`
	Recoverygroupname interface{} `json:"RecoveryGroupName,omitempty"`
}

// CreateRecoveryGroupRequest represents the CreateRecoveryGroupRequest schema from the OpenAPI specification
type CreateRecoveryGroupRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
	Recoverygroupname interface{} `json:"RecoveryGroupName"`
}

// UpdateRecoveryGroupRequest represents the UpdateRecoveryGroupRequest schema from the OpenAPI specification
type UpdateRecoveryGroupRequest struct {
	Cells interface{} `json:"Cells"`
}

// CreateCellRequest represents the CreateCellRequest schema from the OpenAPI specification
type CreateCellRequest struct {
	Cellname interface{} `json:"CellName"`
	Cells interface{} `json:"Cells,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// Resource represents the Resource schema from the OpenAPI specification
type Resource struct {
	Dnstargetresource interface{} `json:"DnsTargetResource,omitempty"`
	Readinessscopes interface{} `json:"ReadinessScopes,omitempty"`
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
	Componentid interface{} `json:"ComponentId,omitempty"`
}

// NLBResource represents the NLBResource schema from the OpenAPI specification
type NLBResource struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// GetRecoveryGroupRequest represents the GetRecoveryGroupRequest schema from the OpenAPI specification
type GetRecoveryGroupRequest struct {
}

// CreateResourceSetRequest represents the CreateResourceSetRequest schema from the OpenAPI specification
type CreateResourceSetRequest struct {
	Resources interface{} `json:"Resources"`
	Tags interface{} `json:"Tags,omitempty"`
	Resourcesetname interface{} `json:"ResourceSetName"`
	Resourcesettype interface{} `json:"ResourceSetType"`
}

// ListReadinessChecksResponse represents the ListReadinessChecksResponse schema from the OpenAPI specification
type ListReadinessChecksResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Readinesschecks interface{} `json:"ReadinessChecks,omitempty"`
}

// CreateCrossAccountAuthorizationRequest represents the CreateCrossAccountAuthorizationRequest schema from the OpenAPI specification
type CreateCrossAccountAuthorizationRequest struct {
	Crossaccountauthorization interface{} `json:"CrossAccountAuthorization"`
}

// ReadinessCheckSummary represents the ReadinessCheckSummary schema from the OpenAPI specification
type ReadinessCheckSummary struct {
	Readiness interface{} `json:"Readiness,omitempty"`
	Readinesscheckname interface{} `json:"ReadinessCheckName,omitempty"`
}

// ListRecoveryGroupsResponse represents the ListRecoveryGroupsResponse schema from the OpenAPI specification
type ListRecoveryGroupsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Recoverygroups interface{} `json:"RecoveryGroups,omitempty"`
}

// GetRecoveryGroupReadinessSummaryResponse represents the GetRecoveryGroupReadinessSummaryResponse schema from the OpenAPI specification
type GetRecoveryGroupReadinessSummaryResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Readiness interface{} `json:"Readiness,omitempty"`
	Readinesschecks interface{} `json:"ReadinessChecks,omitempty"`
}

// ListRulesResponse represents the ListRulesResponse schema from the OpenAPI specification
type ListRulesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// TargetResource represents the TargetResource schema from the OpenAPI specification
type TargetResource struct {
	Nlbresource interface{} `json:"NLBResource,omitempty"`
	R53resource interface{} `json:"R53Resource,omitempty"`
}

// GetCellReadinessSummaryRequest represents the GetCellReadinessSummaryRequest schema from the OpenAPI specification
type GetCellReadinessSummaryRequest struct {
}

// ListCellsRequest represents the ListCellsRequest schema from the OpenAPI specification
type ListCellsRequest struct {
}

// GetArchitectureRecommendationsRequest represents the GetArchitectureRecommendationsRequest schema from the OpenAPI specification
type GetArchitectureRecommendationsRequest struct {
}

// DNSTargetResource represents the DNSTargetResource schema from the OpenAPI specification
type DNSTargetResource struct {
	Recordtype interface{} `json:"RecordType,omitempty"`
	Targetresource interface{} `json:"TargetResource,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Hostedzonearn interface{} `json:"HostedZoneArn,omitempty"`
	Recordsetid interface{} `json:"RecordSetId,omitempty"`
}

// ListRecoveryGroupsRequest represents the ListRecoveryGroupsRequest schema from the OpenAPI specification
type ListRecoveryGroupsRequest struct {
}

// GetReadinessCheckResourceStatusResponse represents the GetReadinessCheckResourceStatusResponse schema from the OpenAPI specification
type GetReadinessCheckResourceStatusResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Readiness interface{} `json:"Readiness,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
}

// Message represents the Message schema from the OpenAPI specification
type Message struct {
	Messagetext interface{} `json:"MessageText,omitempty"`
}

// ListReadinessChecksRequest represents the ListReadinessChecksRequest schema from the OpenAPI specification
type ListReadinessChecksRequest struct {
}

// ReadinessCheckOutput represents the ReadinessCheckOutput schema from the OpenAPI specification
type ReadinessCheckOutput struct {
	Tags interface{} `json:"Tags,omitempty"`
	Readinesscheckarn interface{} `json:"ReadinessCheckArn"`
	Readinesscheckname interface{} `json:"ReadinessCheckName,omitempty"`
	Resourceset interface{} `json:"ResourceSet"`
}

// UpdateResourceSetRequest represents the UpdateResourceSetRequest schema from the OpenAPI specification
type UpdateResourceSetRequest struct {
	Resourcesettype interface{} `json:"ResourceSetType"`
	Resources interface{} `json:"Resources"`
}

// UpdateResourceSetResponse represents the UpdateResourceSetResponse schema from the OpenAPI specification
type UpdateResourceSetResponse struct {
	Resourcesetname interface{} `json:"ResourceSetName,omitempty"`
	Resourcesettype interface{} `json:"ResourceSetType,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Resourcesetarn interface{} `json:"ResourceSetArn,omitempty"`
}

// ResourceSetOutput represents the ResourceSetOutput schema from the OpenAPI specification
type ResourceSetOutput struct {
	Tags interface{} `json:"Tags,omitempty"`
	Resourcesetarn interface{} `json:"ResourceSetArn"`
	Resourcesetname interface{} `json:"ResourceSetName"`
	Resourcesettype interface{} `json:"ResourceSetType"`
	Resources interface{} `json:"Resources"`
}

// CreateReadinessCheckRequest represents the CreateReadinessCheckRequest schema from the OpenAPI specification
type CreateReadinessCheckRequest struct {
	Readinesscheckname interface{} `json:"ReadinessCheckName"`
	Resourcesetname interface{} `json:"ResourceSetName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListResourceSetsRequest represents the ListResourceSetsRequest schema from the OpenAPI specification
type ListResourceSetsRequest struct {
}

// DeleteRecoveryGroupRequest represents the DeleteRecoveryGroupRequest schema from the OpenAPI specification
type DeleteRecoveryGroupRequest struct {
}

// GetRecoveryGroupResponse represents the GetRecoveryGroupResponse schema from the OpenAPI specification
type GetRecoveryGroupResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
	Recoverygrouparn interface{} `json:"RecoveryGroupArn,omitempty"`
	Recoverygroupname interface{} `json:"RecoveryGroupName,omitempty"`
}

// DeleteReadinessCheckRequest represents the DeleteReadinessCheckRequest schema from the OpenAPI specification
type DeleteReadinessCheckRequest struct {
}

// ListTagsForResourcesRequest represents the ListTagsForResourcesRequest schema from the OpenAPI specification
type ListTagsForResourcesRequest struct {
}

// DeleteCrossAccountAuthorizationRequest represents the DeleteCrossAccountAuthorizationRequest schema from the OpenAPI specification
type DeleteCrossAccountAuthorizationRequest struct {
}

// UpdateReadinessCheckResponse represents the UpdateReadinessCheckResponse schema from the OpenAPI specification
type UpdateReadinessCheckResponse struct {
	Readinesscheckarn interface{} `json:"ReadinessCheckArn,omitempty"`
	Readinesscheckname interface{} `json:"ReadinessCheckName,omitempty"`
	Resourceset interface{} `json:"ResourceSet,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListRulesOutput represents the ListRulesOutput schema from the OpenAPI specification
type ListRulesOutput struct {
	Ruleid interface{} `json:"RuleId"`
	Resourcetype interface{} `json:"ResourceType"`
	Ruledescription interface{} `json:"RuleDescription"`
}

// RecoveryGroupOutput represents the RecoveryGroupOutput schema from the OpenAPI specification
type RecoveryGroupOutput struct {
	Cells interface{} `json:"Cells"`
	Recoverygrouparn interface{} `json:"RecoveryGroupArn"`
	Recoverygroupname interface{} `json:"RecoveryGroupName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// RuleResult represents the RuleResult schema from the OpenAPI specification
type RuleResult struct {
	Lastcheckedtimestamp interface{} `json:"LastCheckedTimestamp"`
	Messages interface{} `json:"Messages"`
	Readiness interface{} `json:"Readiness"`
	Ruleid interface{} `json:"RuleId"`
}

// DeleteCellRequest represents the DeleteCellRequest schema from the OpenAPI specification
type DeleteCellRequest struct {
}

// GetCellRequest represents the GetCellRequest schema from the OpenAPI specification
type GetCellRequest struct {
}

// GetArchitectureRecommendationsResponse represents the GetArchitectureRecommendationsResponse schema from the OpenAPI specification
type GetArchitectureRecommendationsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Recommendations interface{} `json:"Recommendations,omitempty"`
	Lastaudittimestamp interface{} `json:"LastAuditTimestamp,omitempty"`
}

// CreateReadinessCheckResponse represents the CreateReadinessCheckResponse schema from the OpenAPI specification
type CreateReadinessCheckResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
	Readinesscheckarn interface{} `json:"ReadinessCheckArn,omitempty"`
	Readinesscheckname interface{} `json:"ReadinessCheckName,omitempty"`
	Resourceset interface{} `json:"ResourceSet,omitempty"`
}

// CreateCellResponse represents the CreateCellResponse schema from the OpenAPI specification
type CreateCellResponse struct {
	Parentreadinessscopes interface{} `json:"ParentReadinessScopes,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Cellarn interface{} `json:"CellArn,omitempty"`
	Cellname interface{} `json:"CellName,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
}

// GetCellReadinessSummaryResponse represents the GetCellReadinessSummaryResponse schema from the OpenAPI specification
type GetCellReadinessSummaryResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Readiness interface{} `json:"Readiness,omitempty"`
	Readinesschecks interface{} `json:"ReadinessChecks,omitempty"`
}

// ListCrossAccountAuthorizationsResponse represents the ListCrossAccountAuthorizationsResponse schema from the OpenAPI specification
type ListCrossAccountAuthorizationsResponse struct {
	Crossaccountauthorizations interface{} `json:"CrossAccountAuthorizations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListCellsResponse represents the ListCellsResponse schema from the OpenAPI specification
type ListCellsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
}

// UpdateCellRequest represents the UpdateCellRequest schema from the OpenAPI specification
type UpdateCellRequest struct {
	Cells interface{} `json:"Cells"`
}

// GetCellResponse represents the GetCellResponse schema from the OpenAPI specification
type GetCellResponse struct {
	Cellarn interface{} `json:"CellArn,omitempty"`
	Cellname interface{} `json:"CellName,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
	Parentreadinessscopes interface{} `json:"ParentReadinessScopes,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// GetReadinessCheckResponse represents the GetReadinessCheckResponse schema from the OpenAPI specification
type GetReadinessCheckResponse struct {
	Readinesscheckarn interface{} `json:"ReadinessCheckArn,omitempty"`
	Readinesscheckname interface{} `json:"ReadinessCheckName,omitempty"`
	Resourceset interface{} `json:"ResourceSet,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// GetRecoveryGroupReadinessSummaryRequest represents the GetRecoveryGroupReadinessSummaryRequest schema from the OpenAPI specification
type GetRecoveryGroupReadinessSummaryRequest struct {
}

// CreateResourceSetResponse represents the CreateResourceSetResponse schema from the OpenAPI specification
type CreateResourceSetResponse struct {
	Resourcesetarn interface{} `json:"ResourceSetArn,omitempty"`
	Resourcesetname interface{} `json:"ResourceSetName,omitempty"`
	Resourcesettype interface{} `json:"ResourceSetType,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateCrossAccountAuthorizationResponse represents the CreateCrossAccountAuthorizationResponse schema from the OpenAPI specification
type CreateCrossAccountAuthorizationResponse struct {
	Crossaccountauthorization interface{} `json:"CrossAccountAuthorization,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// ListTagsForResourcesResponse represents the ListTagsForResourcesResponse schema from the OpenAPI specification
type ListTagsForResourcesResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// Recommendation represents the Recommendation schema from the OpenAPI specification
type Recommendation struct {
	Recommendationtext interface{} `json:"RecommendationText"`
}

// CreateRecoveryGroupResponse represents the CreateRecoveryGroupResponse schema from the OpenAPI specification
type CreateRecoveryGroupResponse struct {
	Cells interface{} `json:"Cells,omitempty"`
	Recoverygrouparn interface{} `json:"RecoveryGroupArn,omitempty"`
	Recoverygroupname interface{} `json:"RecoveryGroupName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// GetReadinessCheckResourceStatusRequest represents the GetReadinessCheckResourceStatusRequest schema from the OpenAPI specification
type GetReadinessCheckResourceStatusRequest struct {
}

// GetReadinessCheckStatusRequest represents the GetReadinessCheckStatusRequest schema from the OpenAPI specification
type GetReadinessCheckStatusRequest struct {
}

// CellOutput represents the CellOutput schema from the OpenAPI specification
type CellOutput struct {
	Cellarn interface{} `json:"CellArn"`
	Cellname interface{} `json:"CellName"`
	Cells interface{} `json:"Cells"`
	Parentreadinessscopes interface{} `json:"ParentReadinessScopes"`
	Tags interface{} `json:"Tags,omitempty"`
}

// Tags represents the Tags schema from the OpenAPI specification
type Tags struct {
}

// ListResourceSetsResponse represents the ListResourceSetsResponse schema from the OpenAPI specification
type ListResourceSetsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcesets interface{} `json:"ResourceSets,omitempty"`
}

// ListCrossAccountAuthorizationsRequest represents the ListCrossAccountAuthorizationsRequest schema from the OpenAPI specification
type ListCrossAccountAuthorizationsRequest struct {
}

// DeleteCrossAccountAuthorizationResponse represents the DeleteCrossAccountAuthorizationResponse schema from the OpenAPI specification
type DeleteCrossAccountAuthorizationResponse struct {
}

// GetReadinessCheckRequest represents the GetReadinessCheckRequest schema from the OpenAPI specification
type GetReadinessCheckRequest struct {
}

// UpdateCellResponse represents the UpdateCellResponse schema from the OpenAPI specification
type UpdateCellResponse struct {
	Cellarn interface{} `json:"CellArn,omitempty"`
	Cellname interface{} `json:"CellName,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
	Parentreadinessscopes interface{} `json:"ParentReadinessScopes,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DeleteResourceSetRequest represents the DeleteResourceSetRequest schema from the OpenAPI specification
type DeleteResourceSetRequest struct {
}
