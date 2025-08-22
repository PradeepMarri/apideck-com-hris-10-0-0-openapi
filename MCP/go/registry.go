package main

import (
	"github.com/hris-api/mcp-server/config"
	"github.com/hris-api/mcp-server/models"
	tools_companies "github.com/hris-api/mcp-server/tools/companies"
	tools_departments "github.com/hris-api/mcp-server/tools/departments"
	tools_time_off_requests "github.com/hris-api/mcp-server/tools/time_off_requests"
	tools_employees "github.com/hris-api/mcp-server/tools/employees"
	tools_employee_schedules "github.com/hris-api/mcp-server/tools/employee_schedules"
	tools_payrolls "github.com/hris-api/mcp-server/tools/payrolls"
	tools_employee_payrolls "github.com/hris-api/mcp-server/tools/employee_payrolls"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_companies.CreateCompaniesallTool(cfg),
		tools_companies.CreateCompaniesaddTool(cfg),
		tools_departments.CreateDepartmentsdeleteTool(cfg),
		tools_departments.CreateDepartmentsoneTool(cfg),
		tools_departments.CreateDepartmentsupdateTool(cfg),
		tools_time_off_requests.CreateTimeoffrequestsdeleteTool(cfg),
		tools_time_off_requests.CreateTimeoffrequestsoneTool(cfg),
		tools_time_off_requests.CreateTimeoffrequestsupdateTool(cfg),
		tools_departments.CreateDepartmentsallTool(cfg),
		tools_departments.CreateDepartmentsaddTool(cfg),
		tools_employees.CreateEmployeesdeleteTool(cfg),
		tools_employees.CreateEmployeesoneTool(cfg),
		tools_employees.CreateEmployeesupdateTool(cfg),
		tools_companies.CreateCompaniesdeleteTool(cfg),
		tools_companies.CreateCompaniesoneTool(cfg),
		tools_companies.CreateCompaniesupdateTool(cfg),
		tools_employees.CreateEmployeesallTool(cfg),
		tools_employees.CreateEmployeesaddTool(cfg),
		tools_employee_schedules.CreateEmployeeschedulesallTool(cfg),
		tools_payrolls.CreatePayrollsallTool(cfg),
		tools_employee_payrolls.CreateEmployeepayrollsallTool(cfg),
		tools_employee_payrolls.CreateEmployeepayrollsoneTool(cfg),
		tools_payrolls.CreatePayrollsoneTool(cfg),
		tools_time_off_requests.CreateTimeoffrequestsallTool(cfg),
		tools_time_off_requests.CreateTimeoffrequestsaddTool(cfg),
	}
}
