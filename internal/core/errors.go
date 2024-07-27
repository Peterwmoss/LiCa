package core

import "errors"

var (
  ErrValidation = errors.New("Validation error")
  ErrAccessDenied = errors.New("Access Denied")
  ErrNotFound = errors.New("Not Found")
)
