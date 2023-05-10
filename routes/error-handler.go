package routes

import "gorm.io/gorm"

type ErrorTypes struct {
	Error  error
	Status int
}

func ErrorValidator(er error) *ErrorTypes {
	if er == gorm.ErrRecordNotFound {
		return &ErrorTypes{
			Error:  er,
			Status: 404,
		}
	}

	return &ErrorTypes{
		Error:  er,
		Status: 404,
	}

}
