package validation

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResult struct {
	Valid  bool              `json:"-"`
	Errors []ValidationError `json:"errors,omitempty"`
}

func NewValidationResult() *ValidationResult {
	return &ValidationResult{
		Valid:  true,
		Errors: make([]ValidationError, 0),
	}
}

func (v *ValidationResult) AddError(field, message string) {
	v.Valid = false
	v.Errors = append(v.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

func (v *ValidationResult) HasError() bool {
	return !v.Valid
}
