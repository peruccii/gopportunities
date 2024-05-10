package handler

import "fmt"

func errParamisRequired(name string, typ string) error {
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
		return errParamisRequired("role", "string")
	}
	if r.Company == "" {
		return errParamisRequired("company", "string")
	}
	if r.Location == "" {
		return errParamisRequired("location", "string")
	}
	if r.Link == "" {
		return errParamisRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamisRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamisRequired("salary", "int64")
	}
	return nil
}