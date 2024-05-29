package list

import (
	"github.com/Peterwmoss/LiCa/internal/core/domain/category"
	"github.com/Peterwmoss/LiCa/internal/core/domain/user"
	"github.com/google/uuid"
)

type List struct {
	Id               uuid.UUID
	Name             Name
	Items            []ListItem
	CategoryOrdering map[int]*category.Category
	User             *user.User
}
