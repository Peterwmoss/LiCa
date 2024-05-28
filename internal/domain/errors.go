package domain

import "errors"

var (
  UniqueViolationError error = errors.New("uniqe violation")
  EmptyNameError error = errors.New("no name provided")
)
