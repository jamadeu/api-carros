package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create Car
type CreateCarRequest struct {
	CarModel     string  `json:"carModel"`
	Manufacturer string  `json:"manufacturer"`
	Color        string  `json:"color"`
	Value        float64 `json:"value"`
}

func (r *CreateCarRequest) Validate() error {
	if r.CarModel == "" && r.Manufacturer == "" && r.Color == "" && r.Value <= 0 {
		return fmt.Errorf("reqest body is empty or malformed")
	}
	if r.CarModel == "" {
		return errParamIsRequired("carModel", "string")
	}
	if r.Manufacturer == "" {
		return errParamIsRequired("manufacturer", "string")
	}
	if r.Color == "" {
		return errParamIsRequired("color", "string")
	}
	if r.Value <= 0 {
		return errParamIsRequired("value", "float64")
	}
	return nil
}
