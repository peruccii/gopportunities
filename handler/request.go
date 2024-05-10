package handler

import "fmt"

func ErrParamisRequired(name string, typ string) error {
	return fmt.Errorf("params %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Salary <= 0 && r.Remote == nil && r.Location == "" && r.Link == "" {
		return fmt.Errorf("request body is empty")
	}
	if r.Role == "" {
		return ErrParamisRequired("role", "string")
	}
	if r.Company == "" {
		return ErrParamisRequired("company", "string")
	}
	if r.Location == "" {
		return ErrParamisRequired("location", "string")
	}
	if r.Link == "" {
		return ErrParamisRequired("link", "string")
	}
	if r.Remote == nil {
		return ErrParamisRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return ErrParamisRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r * UpdateOpeningRequest) Validate() error {
	if r.Role != "" && r.Company != "" && r.Salary > 0 && r.Remote != nil && r.Location != "" && r.Link != "" {
		return nil
	}
	// If none of the fields wew provided, retrun false
	return fmt.Errorf("at least one valid field must be provided")

}