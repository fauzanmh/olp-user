package constant

import "fmt"

type ErrorMessage error

var (
	ErrorMessageCourseCategoryNotFound ErrorMessage = fmt.Errorf("course category not found")
)
