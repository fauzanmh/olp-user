package constant

import "fmt"

type ErrorMessage error

var (
	ErrorMessageCourseCategoryNotFound ErrorMessage = fmt.Errorf("course category not found")
	ErrorMessageUniqueEmail            ErrorMessage = fmt.Errorf("email has been taken")
	ErrorMessageUserHasBeenDeleted     ErrorMessage = fmt.Errorf("user has been deleted")
)
