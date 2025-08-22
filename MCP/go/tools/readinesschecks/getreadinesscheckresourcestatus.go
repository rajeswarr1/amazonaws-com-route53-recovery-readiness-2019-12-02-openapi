package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws-route53-recovery-readiness/mcp-server/config"
	"github.com/aws-route53-recovery-readiness/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetreadinesscheckresourcestatusHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		readinessCheckNameVal, ok := args["readinessCheckName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: readinessCheckName"), nil
		}
		readinessCheckName, ok := readinessCheckNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: readinessCheckName"), nil
		}
		resourceIdentifierVal, ok := args["resourceIdentifier"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceIdentifier"), nil
		}
		resourceIdentifier, ok := resourceIdentifierVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceIdentifier"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/readinesschecks/%s/resource/%s/status%s", cfg.BaseURL, readinessCheckName, resourceIdentifier, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
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
		var result models.GetReadinessCheckResourceStatusResponse
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

func CreateGetreadinesscheckresourcestatusTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_readinesschecks_readinessCheckName_resource_resourceIdentifier_status",
		mcp.WithDescription("Gets individual readiness status for a readiness check. To see the overall readiness status for a recovery group, that considers the readiness status for all the readiness checks in the recovery group, use GetRecoveryGroupReadinessSummary."),
		mcp.WithNumber("maxResults", mcp.Description("The number of objects that you want to return with this call.")),
		mcp.WithString("nextToken", mcp.Description("The token that identifies which batch of results you want to see.")),
		mcp.WithString("readinessCheckName", mcp.Required(), mcp.Description("Name of a readiness check.")),
		mcp.WithString("resourceIdentifier", mcp.Required(), mcp.Description("The resource identifier, which is the Amazon Resource Name (ARN) or the identifier generated for the resource by Application Recovery Controller (for example, for a DNS target resource).")),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetreadinesscheckresourcestatusHandler(cfg),
	}
}
