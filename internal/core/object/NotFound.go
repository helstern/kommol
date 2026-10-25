package object

import "errors"

// ErrNotFound marks an object that does not exist, as opposed to one that could not be read.
var ErrNotFound = errors.New("object not found")
