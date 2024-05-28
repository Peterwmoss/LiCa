package domain

import "errors"

var (
  UniqueViolationError error = errors.New("uniqe violation")
)
