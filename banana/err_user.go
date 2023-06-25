package banana

import "errors"

var (
	UserConflict = errors.New("User is exist")
	SignUpFail   = errors.New("Signup Failed")
)
