package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/hris-api/mcp-server/config"
	"github.com/hris-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func EmployeesaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["raw"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("raw=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Employee
		
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
		url := fmt.Sprintf("%s/hris/employees%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.CreateEmployeeResponse
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

func CreateEmployeesaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_hris_employees",
		mcp.WithDescription("Create Employee"),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithBoolean("deleted", mcp.Description("Input parameter: Flag to indicate if the object is deleted.")),
		mcp.WithString("country_of_birth", mcp.Description("Input parameter: Country code according to ISO 3166-1 alpha-2.")),
		mcp.WithString("source", mcp.Description("Input parameter: When the employee is imported as a new hire, this field indicates what system (e.g. the name of the ATS) this employee was imported from.")),
		mcp.WithString("source_id", mcp.Description("Input parameter: Unique identifier of the employee in the system this employee was imported from (e.g. the ID in the ATS).")),
		mcp.WithString("photo_url", mcp.Description("Input parameter: The URL of the photo of a person.")),
		mcp.WithString("marital_status", mcp.Description("Input parameter: The marital status of the employee.")),
		mcp.WithArray("social_links", mcp.Description("")),
		mcp.WithArray("jobs", mcp.Description("")),
		mcp.WithArray("bank_accounts", mcp.Description("")),
		mcp.WithString("employment_status", mcp.Description("Input parameter: The employment status of the employee, indicating whether they are currently employed, inactive, terminated, or in another status.")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithString("employee_number", mcp.Description("Input parameter: An Employee Number, Employee ID or Employee Code, is a unique number that has been assigned to each individual staff member within a company.")),
		mcp.WithString("department", mcp.Description("Input parameter: The department the person is currently in. [Deprecated](https://developers.apideck.com/changelog) in favor of the dedicated department_id and department_name field.")),
		mcp.WithString("department_name", mcp.Description("Input parameter: Name of the department this employee belongs to.")),
		mcp.WithString("tax_id", mcp.Description("")),
		mcp.WithString("preferred_name", mcp.Description("Input parameter: The name the employee prefers to be addressed by, which may be different from their legal name.")),
		mcp.WithString("department_id", mcp.Description("Input parameter: Unique identifier of the department ID this employee belongs to.")),
		mcp.WithArray("direct_reports", mcp.Description("Input parameter: Direct reports is an array of ids that reflect the individuals in an organizational hierarchy who are directly supervised by this specific employee.")),
		mcp.WithString("pronouns", mcp.Description("Input parameter: The preferred pronouns of the person.")),
		mcp.WithString("division", mcp.Description("Input parameter: The division the person is currently in. Usually a collection of departments or teams or regions.")),
		mcp.WithString("salutation", mcp.Description("Input parameter: A formal salutation for the person. For example, 'Mr', 'Mrs'")),
		mcp.WithString("first_name", mcp.Description("Input parameter: The first name of the person.")),
		mcp.WithObject("employment_role", mcp.Description("")),
		mcp.WithString("title", mcp.Description("Input parameter: The job title of the person.")),
		mcp.WithBoolean("works_remote", mcp.Description("Input parameter: Indicates if the employee works from a remote location.")),
		mcp.WithString("dietary_preference", mcp.Description("Input parameter: Indicate the employee's dietary preference.")),
		mcp.WithString("deceased_on", mcp.Description("Input parameter: The date the person deceased.")),
		mcp.WithString("birthday", mcp.Description("Input parameter: The date of birth of the person.")),
		mcp.WithString("initials", mcp.Description("Input parameter: The initials of the person, usually derived from their first, middle, and last names.")),
		mcp.WithString("last_name", mcp.Description("Input parameter: The last name of the person.")),
		mcp.WithString("ethnicity", mcp.Description("Input parameter: The ethnicity of the employee")),
		mcp.WithArray("phone_numbers", mcp.Description("")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("leaving_reason", mcp.Description("Input parameter: The reason because the employment ended.")),
		mcp.WithString("tax_code", mcp.Description("")),
		mcp.WithArray("custom_fields", mcp.Description("")),
		mcp.WithObject("probation_period", mcp.Description("")),
		mcp.WithObject("manager", mcp.Description("")),
		mcp.WithArray("languages", mcp.Description("")),
		mcp.WithString("division_id", mcp.Description("Input parameter: Unique identifier of the division this employee belongs to.")),
		mcp.WithString("gender", mcp.Description("Input parameter: The gender represents the gender identity of a person.")),
		mcp.WithString("employment_start_date", mcp.Description("Input parameter: A Start Date is the date that the employee started working at the company")),
		mcp.WithString("preferred_language", mcp.Description("Input parameter: language code according to ISO 639-1. For the United States - EN")),
		mcp.WithArray("emails", mcp.Description("")),
		mcp.WithString("company_name", mcp.Description("Input parameter: The name of the company.")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithString("timezone", mcp.Description("Input parameter: The time zone related to the resource. The value is a string containing a standard time zone identifier, e.g. Europe/London.")),
		mcp.WithObject("team", mcp.Description("Input parameter: The team the person is currently in.")),
		mcp.WithString("middle_name", mcp.Description("Input parameter: Middle name of the person.")),
		mcp.WithString("company_id", mcp.Description("Input parameter: The unique identifier of the company.")),
		mcp.WithArray("addresses", mcp.Description("")),
		mcp.WithArray("food_allergies", mcp.Description("Input parameter: Indicate the employee's food allergies.")),
		mcp.WithString("record_url", mcp.Description("")),
		mcp.WithString("employment_end_date", mcp.Description("Input parameter: An End Date is the date that the employee ended working at the company")),
		mcp.WithString("display_name", mcp.Description("Input parameter: The name used to display the employee, often a combination of their first and last names.")),
		mcp.WithArray("nationalities", mcp.Description("")),
		mcp.WithString("social_security_number", mcp.Description("Input parameter: A unique identifier assigned by the government. This field is considered sensitive information and may be subject to special security and privacy restrictions.")),
		mcp.WithArray("tags", mcp.Description("")),
		mcp.WithArray("compensations", mcp.Description("")),
		mcp.WithString("row_version", mcp.Description("Input parameter: A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.")),
		mcp.WithString("description", mcp.Description("Input parameter: A description of the object.")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Input parameter: A unique identifier for an object.")),
		mcp.WithObject("partner", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    EmployeesaddHandler(cfg),
	}
}
