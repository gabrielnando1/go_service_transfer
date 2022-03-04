package error

import (
	"fmt"
)

type UnauthorizedRequestError struct {
	//StatusCode int
	Err error
}

func (r *UnauthorizedRequestError) Error() string {
	return fmt.Sprintf("%v", r.Err)
}
