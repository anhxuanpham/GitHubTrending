package banana

import "errors"

var (
	UserConflict   = errors.New("User Already Exists")
	SignUpFail     = errors.New("Signup Failed")
	UserNotFound   = errors.New("User Does Not Exist")
	UserNotUpdated = errors.New("User Not Updated")
)
