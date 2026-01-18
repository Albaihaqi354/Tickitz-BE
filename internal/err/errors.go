package err

import "errors"

var (
	ErrNoRowsUpdated = errors.New("no rows updated")
	ErrInvalidExt    = errors.New("invalid file extension")
)
