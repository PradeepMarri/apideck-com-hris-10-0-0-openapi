package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hris-api/mcp-server/config"
	"github.com/hris-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func EmployeepayrollsoneHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		payroll_idVal, ok := args["payroll_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: payroll_id"), nil
		}
		payroll_id, ok := payroll_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: payroll_id"), nil
		}
		employee_idVal, ok := args["employee_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: employee_id"), nil
		}
		employee_id, ok := employee_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: employee_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["raw"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("raw=%v", val))
		}
		if val, ok := args["fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/hris/payrolls/employees/%s/payrolls/%s%s", cfg.BaseURL, payroll_id, employee_id, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-apideck-consumer-id"]; ok {
			req.Header.Set("x-apideck-consumer-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-app-id"]; ok {
			req.Header.Set("x-apideck-app-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-service-id"]; ok {
			req.Header.Set("x-apideck-service-id", fmt.Sprintf("%v", val))
		}

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
		var result models.GetEmployeePayrollResponse
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

func CreateEmployeepayrollsoneTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_hris_payrolls_employees_employee_id_payrolls_payroll_id",
		mcp.WithDescription("Get Employee Payroll"),
		mcp.WithString("payroll_id", mcp.Required(), mcp.Description("ID of the payroll you are acting upon.")),
		mcp.WithString("employee_id", mcp.Required(), mcp.Description("ID of the employee you are acting upon.")),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithString("fields", mcp.Description("The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields \"name\", \"email\" and \"addresses.city\". If any other fields are available, they will be excluded.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    EmployeepayrollsoneHandler(cfg),
	}
}
