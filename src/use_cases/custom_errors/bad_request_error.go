package error

import (
	"fmt"
)

type BadRequestError struct {
	//StatusCode int
	Err error
}

func (r *BadRequestError) Error() string {
	return fmt.Sprintf("%v", r.Err)
}
