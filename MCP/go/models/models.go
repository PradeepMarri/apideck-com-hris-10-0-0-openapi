package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteDepartmentResponse represents the DeleteDepartmentResponse schema from the OpenAPI specification
type DeleteDepartmentResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// TimeOffRequest represents the TimeOffRequest schema from the OpenAPI specification
type TimeOffRequest struct {
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Request_date string `json:"request_date,omitempty"` // The date the request was made.
	Start_date string `json:"start_date,omitempty"` // The start date of the time off request.
	Approval_date string `json:"approval_date,omitempty"` // The date the request was approved
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Status string `json:"status,omitempty"` // The status of the time off request.
	End_date string `json:"end_date,omitempty"` // The end date of the time off request.
	Amount float64 `json:"amount,omitempty"` // The amount of time off requested.
	Units string `json:"units,omitempty"` // The unit of time off requested. Possible values include: `hours`, `days`, or `other`.
	Employee_id string `json:"employee_id,omitempty"` // ID of the employee
	Notes map[string]interface{} `json:"notes,omitempty"`
	Request_type string `json:"request_type,omitempty"` // The type of request
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Description string `json:"description,omitempty"` // Description of the time off request.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Policy_id string `json:"policy_id,omitempty"` // ID of the policy
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// HrisCompany represents the HrisCompany schema from the OpenAPI specification
type HrisCompany struct {
	Company_number string `json:"company_number,omitempty"` // An Company Number, Company ID or Company Code, is a unique number that has been assigned to each company.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Display_name string `json:"display_name,omitempty"`
	Status string `json:"status,omitempty"`
	Subdomain string `json:"subdomain,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Debtor_id string `json:"debtor_id,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Emails []Email `json:"emails,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Websites []Website `json:"websites,omitempty"`
	Legal_name string `json:"legal_name"`
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// GetTimeOffRequestResponse represents the GetTimeOffRequestResponse schema from the OpenAPI specification
type GetTimeOffRequestResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data TimeOffRequest `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// EmployeesSort represents the EmployeesSort schema from the OpenAPI specification
type EmployeesSort struct {
	By string `json:"by,omitempty"` // The field on which to sort the Employees
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
}

// Schedule represents the Schedule schema from the OpenAPI specification
type Schedule struct {
	Id string `json:"id"` // A unique identifier for an object.
	Start_date string `json:"start_date"` // The start date, inclusive, of the schedule period.
	Work_pattern map[string]interface{} `json:"work_pattern"`
	End_date string `json:"end_date"` // The end date, inclusive, of the schedule period.
}

// EmployeeJob represents the EmployeeJob schema from the OpenAPI specification
type EmployeeJob struct {
	Hired_at string `json:"hired_at,omitempty"` // The date on which the employee was hired by the organization
	Role string `json:"role,omitempty"` // The position and responsibilities of the person within the organization.
	Start_date string `json:"start_date,omitempty"` // The date on which the employee starts working in their current job role.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Location Address `json:"location,omitempty"`
	Payment_unit string `json:"payment_unit,omitempty"` // Unit of measurement for employee compensation.
	Is_primary bool `json:"is_primary,omitempty"` // Indicates whether this the employee's primary job.
	Title string `json:"title,omitempty"` // The job title of the person.
	Compensation_rate float64 `json:"compensation_rate,omitempty"` // The rate of pay for the employee in their current job role.
	Employee_id string `json:"employee_id,omitempty"` // A unique identifier for an object.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	End_date string `json:"end_date,omitempty"` // The date on which the employee leaves or is expected to leave their current job role.
}

// Compensation represents the Compensation schema from the OpenAPI specification
type Compensation struct {
	Net_pay float64 `json:"net_pay,omitempty"` // The employee's net pay. Only available when payroll has been processed
	Taxes []Tax `json:"taxes,omitempty"` // An array of employer and employee taxes for the pay period.
	Benefits []Benefit `json:"benefits,omitempty"` // An array of employee benefits for the pay period.
	Deductions []Deduction `json:"deductions,omitempty"` // An array of employee deductions for the pay period.
	Employee_id string `json:"employee_id"` // A unique identifier for an object.
	Gross_pay float64 `json:"gross_pay,omitempty"` // The employee's gross pay. Only available when payroll has been processed
}

// Payroll represents the Payroll schema from the OpenAPI specification
type Payroll struct {
	Company_id string `json:"company_id,omitempty"` // The unique identifier of the company.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Processed bool `json:"processed"` // Whether or not the payroll has been successfully processed. Note that processed payrolls cannot be updated.
	Start_date string `json:"start_date"` // The start date, inclusive, of the pay period.
	Compensations []Compensation `json:"compensations,omitempty"` // An array of compensations for the payroll.
	End_date string `json:"end_date"` // The end date, inclusive, of the pay period.
	Processed_date string `json:"processed_date,omitempty"` // The date the payroll was processed.
	Id string `json:"id"` // A unique identifier for an object.
	Check_date string `json:"check_date"` // The date on which employees will be paid for the payroll.
	Totals PayrollTotals `json:"totals,omitempty"`
}

// GetHrisJobResponse represents the GetHrisJobResponse schema from the OpenAPI specification
type GetHrisJobResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data HrisJob `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// TimeOffRequestsFilter represents the TimeOffRequestsFilter schema from the OpenAPI specification
type TimeOffRequestsFilter struct {
	End_date string `json:"end_date,omitempty"` // End date
	Start_date string `json:"start_date,omitempty"` // Start date
	Time_off_request_status string `json:"time_off_request_status,omitempty"` // Time off request status to filter on
	Employee_id string `json:"employee_id,omitempty"` // Employee ID
}

// GetEmployeePayrollResponse represents the GetEmployeePayrollResponse schema from the OpenAPI specification
type GetEmployeePayrollResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data EmployeePayroll `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// EmployeeCompensation represents the EmployeeCompensation schema from the OpenAPI specification
type EmployeeCompensation struct {
	Rate float64 `json:"rate,omitempty"` // The amount paid per payment unit.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Effective_date string `json:"effective_date,omitempty"` // The date on which a change to an employee's compensation takes effect.
	Flsa_status string `json:"flsa_status,omitempty"` // The FLSA status for this compensation.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Job_id string `json:"job_id,omitempty"` // The ID of the job to which the compensation belongs.
	Payment_frequency string `json:"payment_frequency,omitempty"` // Frequency of employee compensation.
	Payment_unit string `json:"payment_unit,omitempty"` // Unit of measurement for employee compensation.
}

// GetTimeOffRequestsResponse represents the GetTimeOffRequestsResponse schema from the OpenAPI specification
type GetTimeOffRequestsResponse struct {
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []TimeOffRequest `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
}

// Tax represents the Tax schema from the OpenAPI specification
type Tax struct {
	Name string `json:"name,omitempty"` // The name of the tax.
	Amount float64 `json:"amount,omitempty"` // The amount of the tax.
	Employer bool `json:"employer,omitempty"` // Paid by employer.
}

// HrisWebhookEvent represents the HrisWebhookEvent schema from the OpenAPI specification
type HrisWebhookEvent struct {
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Event_type string `json:"event_type,omitempty"`
}

// PayrollTotals represents the PayrollTotals schema from the OpenAPI specification
type PayrollTotals struct {
	Check_amount float64 `json:"check_amount,omitempty"` // The total check amount for the payroll.
	Employee_taxes float64 `json:"employee_taxes,omitempty"` // The total amount of employee paid taxes for the payroll.
	Employer_benefit_contributions float64 `json:"employer_benefit_contributions,omitempty"` // The total amount of company contributed benefits for the payroll.
	Employer_taxes float64 `json:"employer_taxes,omitempty"` // The total amount of employer paid taxes for the payroll.
	Gross_pay float64 `json:"gross_pay,omitempty"` // The gross pay amount for the payroll.
	Company_debit float64 `json:"company_debit,omitempty"` // The total company debit for the payroll.
	Employee_benefit_deductions float64 `json:"employee_benefit_deductions,omitempty"` // The total amount of employee deducted benefits for the payroll.
	Net_pay float64 `json:"net_pay,omitempty"` // The net pay amount for the payroll.
	Tax_debit float64 `json:"tax_debit,omitempty"` // The total tax debit for the payroll.
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	TypeField string `json:"type,omitempty"` // The type of address.
	Name string `json:"name,omitempty"` // The name of the address.
	Longitude string `json:"longitude,omitempty"` // Longitude of the address
	StringField string `json:"string,omitempty"` // The address string. Some APIs don't provide structured address data.
	State string `json:"state,omitempty"` // Name of state
	Line2 string `json:"line2,omitempty"` // Line 2 of the address
	Phone_number string `json:"phone_number,omitempty"` // Phone number of the address
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	City string `json:"city,omitempty"` // Name of city.
	Notes string `json:"notes,omitempty"` // Additional notes
	County string `json:"county,omitempty"` // Address field that holds a sublocality, such as a county
	Line3 string `json:"line3,omitempty"` // Line 3 of the address
	Website string `json:"website,omitempty"` // Website of the address
	Line1 string `json:"line1,omitempty"` // Line 1 of the address e.g. number, street, suite, apt #, etc.
	Salutation string `json:"salutation,omitempty"` // Salutation of the contact person at the address
	Email string `json:"email,omitempty"` // Email address of the address
	Street_number string `json:"street_number,omitempty"` // Street number
	Id string `json:"id,omitempty"` // Unique identifier for the address.
	Line4 string `json:"line4,omitempty"` // Line 4 of the address
	Postal_code string `json:"postal_code,omitempty"` // Zip code or equivalent.
	Contact_name string `json:"contact_name,omitempty"` // Name of the contact person at the address
	Country string `json:"country,omitempty"` // country code according to ISO 3166-1 alpha-2.
	Fax string `json:"fax,omitempty"` // Fax number of the address
	Latitude string `json:"latitude,omitempty"` // Latitude of the address
}

// EmployeePayroll represents the EmployeePayroll schema from the OpenAPI specification
type EmployeePayroll struct {
	Processed_date string `json:"processed_date,omitempty"` // The date the payroll was processed.
	Employee_id string `json:"employee_id,omitempty"` // ID of the employee
	Id string `json:"id"` // A unique identifier for an object.
	Processed bool `json:"processed"` // Whether or not the payroll has been successfully processed. Note that processed payrolls cannot be updated.
	Start_date string `json:"start_date"` // The start date, inclusive, of the pay period.
	Company_id string `json:"company_id,omitempty"` // The unique identifier of the company.
	Compensations []Compensation `json:"compensations,omitempty"` // An array of compensations for the payroll.
	Totals PayrollTotals `json:"totals,omitempty"`
	Check_date string `json:"check_date"` // The date on which employees will be paid for the payroll.
	End_date string `json:"end_date"` // The end date, inclusive, of the pay period.
}

// GetHrisCompanyResponse represents the GetHrisCompanyResponse schema from the OpenAPI specification
type GetHrisCompanyResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data HrisCompany `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// DeleteEmployeeResponse represents the DeleteEmployeeResponse schema from the OpenAPI specification
type DeleteEmployeeResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// WebhookEvent represents the WebhookEvent schema from the OpenAPI specification
type WebhookEvent struct {
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// DeleteHrisCompanyResponse represents the DeleteHrisCompanyResponse schema from the OpenAPI specification
type DeleteHrisCompanyResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// PayrollsFilter represents the PayrollsFilter schema from the OpenAPI specification
type PayrollsFilter struct {
	Start_date string `json:"start_date,omitempty"` // Return payrolls whose pay period is after the start date
	End_date string `json:"end_date,omitempty"` // Return payrolls whose pay period is before the end date
}

// DeleteTimeOffRequestResponse represents the DeleteTimeOffRequestResponse schema from the OpenAPI specification
type DeleteTimeOffRequestResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// GetPayrollResponse represents the GetPayrollResponse schema from the OpenAPI specification
type GetPayrollResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Payroll `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// Website represents the Website schema from the OpenAPI specification
type Website struct {
	Url string `json:"url"` // The website URL
	Id string `json:"id,omitempty"` // Unique identifier for the website
	TypeField string `json:"type,omitempty"` // The type of website
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// Employee represents the Employee schema from the OpenAPI specification
type Employee struct {
	Emails []Email `json:"emails,omitempty"`
	Company_name string `json:"company_name,omitempty"` // The name of the company.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Timezone string `json:"timezone,omitempty"` // The time zone related to the resource. The value is a string containing a standard time zone identifier, e.g. Europe/London.
	Team Team `json:"team,omitempty"` // The team the person is currently in.
	Middle_name string `json:"middle_name,omitempty"` // Middle name of the person.
	Company_id string `json:"company_id,omitempty"` // The unique identifier of the company.
	Addresses []Address `json:"addresses,omitempty"`
	Food_allergies []string `json:"food_allergies,omitempty"` // Indicate the employee's food allergies.
	Record_url string `json:"record_url,omitempty"`
	Employment_end_date string `json:"employment_end_date,omitempty"` // An End Date is the date that the employee ended working at the company
	Display_name string `json:"display_name,omitempty"` // The name used to display the employee, often a combination of their first and last names.
	Nationalities []string `json:"nationalities,omitempty"`
	Social_security_number string `json:"social_security_number,omitempty"` // A unique identifier assigned by the government. This field is considered sensitive information and may be subject to special security and privacy restrictions.
	Tags []string `json:"tags,omitempty"`
	Compensations []EmployeeCompensation `json:"compensations,omitempty"`
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Description string `json:"description,omitempty"` // A description of the object.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Id string `json:"id"` // A unique identifier for an object.
	Partner Person `json:"partner,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Deleted bool `json:"deleted,omitempty"` // Flag to indicate if the object is deleted.
	Country_of_birth string `json:"country_of_birth,omitempty"` // Country code according to ISO 3166-1 alpha-2.
	Source string `json:"source,omitempty"` // When the employee is imported as a new hire, this field indicates what system (e.g. the name of the ATS) this employee was imported from.
	Source_id string `json:"source_id,omitempty"` // Unique identifier of the employee in the system this employee was imported from (e.g. the ID in the ATS).
	Photo_url string `json:"photo_url,omitempty"` // The URL of the photo of a person.
	Marital_status string `json:"marital_status,omitempty"` // The marital status of the employee.
	Social_links []SocialLink `json:"social_links,omitempty"`
	Jobs []EmployeeJob `json:"jobs,omitempty"`
	Bank_accounts []BankAccount `json:"bank_accounts,omitempty"`
	Employment_status string `json:"employment_status,omitempty"` // The employment status of the employee, indicating whether they are currently employed, inactive, terminated, or in another status.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Employee_number string `json:"employee_number,omitempty"` // An Employee Number, Employee ID or Employee Code, is a unique number that has been assigned to each individual staff member within a company.
	Department string `json:"department,omitempty"` // The department the person is currently in. [Deprecated](https://developers.apideck.com/changelog) in favor of the dedicated department_id and department_name field.
	Department_name string `json:"department_name,omitempty"` // Name of the department this employee belongs to.
	Tax_id string `json:"tax_id,omitempty"`
	Preferred_name string `json:"preferred_name,omitempty"` // The name the employee prefers to be addressed by, which may be different from their legal name.
	Department_id string `json:"department_id,omitempty"` // Unique identifier of the department ID this employee belongs to.
	Direct_reports []string `json:"direct_reports,omitempty"` // Direct reports is an array of ids that reflect the individuals in an organizational hierarchy who are directly supervised by this specific employee.
	Pronouns string `json:"pronouns,omitempty"` // The preferred pronouns of the person.
	Division string `json:"division,omitempty"` // The division the person is currently in. Usually a collection of departments or teams or regions.
	Salutation string `json:"salutation,omitempty"` // A formal salutation for the person. For example, 'Mr', 'Mrs'
	First_name string `json:"first_name,omitempty"` // The first name of the person.
	Employment_role map[string]interface{} `json:"employment_role,omitempty"`
	Title string `json:"title,omitempty"` // The job title of the person.
	Works_remote bool `json:"works_remote,omitempty"` // Indicates if the employee works from a remote location.
	Dietary_preference string `json:"dietary_preference,omitempty"` // Indicate the employee's dietary preference.
	Deceased_on string `json:"deceased_on,omitempty"` // The date the person deceased.
	Birthday string `json:"birthday,omitempty"` // The date of birth of the person.
	Initials string `json:"initials,omitempty"` // The initials of the person, usually derived from their first, middle, and last names.
	Last_name string `json:"last_name,omitempty"` // The last name of the person.
	Ethnicity string `json:"ethnicity,omitempty"` // The ethnicity of the employee
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Leaving_reason string `json:"leaving_reason,omitempty"` // The reason because the employment ended.
	Tax_code string `json:"tax_code,omitempty"`
	Custom_fields []CustomField `json:"custom_fields,omitempty"`
	Probation_period map[string]interface{} `json:"probation_period,omitempty"`
	Manager map[string]interface{} `json:"manager,omitempty"`
	Languages []string `json:"languages,omitempty"`
	Division_id string `json:"division_id,omitempty"` // Unique identifier of the division this employee belongs to.
	Gender string `json:"gender,omitempty"` // The gender represents the gender identity of a person.
	Employment_start_date string `json:"employment_start_date,omitempty"` // A Start Date is the date that the employee started working at the company
	Preferred_language string `json:"preferred_language,omitempty"` // language code according to ISO 639-1. For the United States - EN
}

// UpdateDepartmentResponse represents the UpdateDepartmentResponse schema from the OpenAPI specification
type UpdateDepartmentResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// Department represents the Department schema from the OpenAPI specification
type Department struct {
	Code string `json:"code,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Name string `json:"name,omitempty"` // Department name
	Parent_id string `json:"parent_id,omitempty"` // Parent ID
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
}

// CreateHrisCompanyResponse represents the CreateHrisCompanyResponse schema from the OpenAPI specification
type CreateHrisCompanyResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
}

// EmployeesFilter represents the EmployeesFilter schema from the OpenAPI specification
type EmployeesFilter struct {
	Last_name string `json:"last_name,omitempty"` // Last Name to filter on
	Employee_number string `json:"employee_number,omitempty"` // Employee number to filter on
	First_name string `json:"first_name,omitempty"` // First Name to filter on
	Company_id string `json:"company_id,omitempty"` // Company ID to filter on
	Employment_status string `json:"employment_status,omitempty"` // Employment status to filter on
	Manager_id string `json:"manager_id,omitempty"` // Manager id to filter on
	Title string `json:"title,omitempty"` // Job title to filter on
	Department_id string `json:"department_id,omitempty"` // ID of the department to filter on
	Email string `json:"email,omitempty"` // Email to filter on
}

// GetHrisCompaniesResponse represents the GetHrisCompaniesResponse schema from the OpenAPI specification
type GetHrisCompaniesResponse struct {
	Data []HrisCompany `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetDepartmentResponse represents the GetDepartmentResponse schema from the OpenAPI specification
type GetDepartmentResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Department `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// EmployeeList represents the EmployeeList schema from the OpenAPI specification
type EmployeeList struct {
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	First_name string `json:"first_name,omitempty"` // The first name of the person.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Last_name string `json:"last_name,omitempty"` // The last name of the person.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
}

// GetDepartmentsResponse represents the GetDepartmentsResponse schema from the OpenAPI specification
type GetDepartmentsResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Department `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Id string `json:"id"` // Unique identifier for the custom field.
	Name string `json:"name,omitempty"` // Name of the custom field.
	Value interface{} `json:"value,omitempty"`
	Description string `json:"description,omitempty"` // More information about the custom field
}

// UpdateHrisCompanyResponse represents the UpdateHrisCompanyResponse schema from the OpenAPI specification
type UpdateHrisCompanyResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// Person represents the Person schema from the OpenAPI specification
type Person struct {
	Initials string `json:"initials,omitempty"` // Initials of the person
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Last_name string `json:"last_name,omitempty"` // The last name of the person.
	First_name string `json:"first_name,omitempty"` // The first name of the person.
	Gender string `json:"gender,omitempty"` // The gender represents the gender identity of a person.
	Middle_name string `json:"middle_name,omitempty"` // Middle name of the person.
	Deceased_on string `json:"deceased_on,omitempty"` // Date of death
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Birthday string `json:"birthday,omitempty"` // Date of birth
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// PassThroughQuery represents the PassThroughQuery schema from the OpenAPI specification
type PassThroughQuery struct {
	Example_downstream_property string `json:"example_downstream_property,omitempty"` // All passthrough query parameters are passed along to the connector as is (?pass_through[search]=leads becomes ?search=leads)
}

// CreateDepartmentResponse represents the CreateDepartmentResponse schema from the OpenAPI specification
type CreateDepartmentResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// GetEmployeePayrollsResponse represents the GetEmployeePayrollsResponse schema from the OpenAPI specification
type GetEmployeePayrollsResponse struct {
	Data []EmployeePayroll `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetPayrollsResponse represents the GetPayrollsResponse schema from the OpenAPI specification
type GetPayrollsResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Payroll `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// Benefit represents the Benefit schema from the OpenAPI specification
type Benefit struct {
	Employee_deduction float64 `json:"employee_deduction,omitempty"` // The amount deducted for benefit.
	Employer_contribution float64 `json:"employer_contribution,omitempty"` // The amount of employer contribution.
	Name string `json:"name,omitempty"` // The name of the benefit.
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// GetEmployeeResponse represents the GetEmployeeResponse schema from the OpenAPI specification
type GetEmployeeResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Employee `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Id string `json:"id,omitempty"` // Unique identifier for the email address
	TypeField string `json:"type,omitempty"` // Email type
	Email string `json:"email"` // Email address
}

// Deduction represents the Deduction schema from the OpenAPI specification
type Deduction struct {
	Amount float64 `json:"amount,omitempty"` // The amount deducted.
	Name string `json:"name,omitempty"` // The name of the deduction.
}

// GetHrisJobsResponse represents the GetHrisJobsResponse schema from the OpenAPI specification
type GetHrisJobsResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data HrisJobs `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// UpdateEmployeeResponse represents the UpdateEmployeeResponse schema from the OpenAPI specification
type UpdateEmployeeResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// CreateEmployeeResponse represents the CreateEmployeeResponse schema from the OpenAPI specification
type CreateEmployeeResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// HrisJob represents the HrisJob schema from the OpenAPI specification
type HrisJob struct {
	Is_primary bool `json:"is_primary,omitempty"` // Indicates whether this the employee's primary job.
	Start_date string `json:"start_date,omitempty"`
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Location map[string]interface{} `json:"location,omitempty"`
	Title string `json:"title,omitempty"` // The job title of the person.
	Department string `json:"department,omitempty"` // Department name
	Employee_id string `json:"employee_id,omitempty"` // Id of the employee
	Employment_status string `json:"employment_status,omitempty"` // The employment status of the employee, indicating whether they are currently employed, inactive, terminated, or in another status.
	End_date string `json:"end_date,omitempty"`
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// HrisJobs represents the HrisJobs schema from the OpenAPI specification
type HrisJobs struct {
	Jobs []HrisJob `json:"jobs,omitempty"`
	Employee Employee `json:"employee,omitempty"`
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	Area_code string `json:"area_code,omitempty"` // The area code of the phone number, e.g. 323
	Country_code string `json:"country_code,omitempty"` // The country code of the phone number, e.g. +1
	Extension string `json:"extension,omitempty"` // The extension of the phone number
	Id string `json:"id,omitempty"` // Unique identifier of the phone number
	Number string `json:"number"` // The phone number
	TypeField string `json:"type,omitempty"` // The type of phone number
}

// UpdateTimeOffRequestResponse represents the UpdateTimeOffRequestResponse schema from the OpenAPI specification
type UpdateTimeOffRequestResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// SocialLink represents the SocialLink schema from the OpenAPI specification
type SocialLink struct {
	Id string `json:"id,omitempty"` // Unique identifier of the social link
	TypeField string `json:"type,omitempty"` // Type of the social link, e.g. twitter
	Url string `json:"url"` // URL of the social link, e.g. https://www.twitter.com/apideck
}

// GetEmployeesResponse represents the GetEmployeesResponse schema from the OpenAPI specification
type GetEmployeesResponse struct {
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Employee `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
}

// CreateTimeOffRequestResponse represents the CreateTimeOffRequestResponse schema from the OpenAPI specification
type CreateTimeOffRequestResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// EmployeeSchedules represents the EmployeeSchedules schema from the OpenAPI specification
type EmployeeSchedules struct {
	Employee Employee `json:"employee,omitempty"`
	Schedules []Schedule `json:"schedules,omitempty"`
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// BankAccount represents the BankAccount schema from the OpenAPI specification
type BankAccount struct {
	Bsb_number string `json:"bsb_number,omitempty"` // A BSB is a 6 digit numeric code used for identifying the branch of an Australian or New Zealand bank or financial institution.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Iban string `json:"iban,omitempty"` // The International Bank Account Number (IBAN).
	Bic string `json:"bic,omitempty"` // The Bank Identifier Code (BIC).
	Routing_number string `json:"routing_number,omitempty"` // A routing number is a nine-digit code used to identify a financial institution in the United States.
	Account_type string `json:"account_type,omitempty"` // The type of bank account.
	Bank_name string `json:"bank_name,omitempty"` // The name of the bank
	Account_name string `json:"account_name,omitempty"` // The name which you used in opening your bank account.
	Branch_identifier string `json:"branch_identifier,omitempty"` // A branch identifier is a unique identifier for a branch of a bank or financial institution.
	Account_number string `json:"account_number,omitempty"` // A bank account number is a number that is tied to your bank account. If you have several bank accounts, such as personal, joint, business (and so on), each account will have a different account number.
	Bank_code string `json:"bank_code,omitempty"` // A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Name string `json:"name,omitempty"` // The name of the team.
	Id string `json:"id,omitempty"` // The unique identifier of the team.
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// GetEmployeeSchedulesResponse represents the GetEmployeeSchedulesResponse schema from the OpenAPI specification
type GetEmployeeSchedulesResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data EmployeeSchedules `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}
