package product

import (
	"github.com/Peterwmoss/LiCa/internal/core/domain/category"
	"github.com/Peterwmoss/LiCa/internal/core/domain/user"
	"github.com/google/uuid"
)

type Product struct {
	Id              uuid.UUID
	Name            Name
	DefaultCategory *category.Category
	IsCustom        bool
	User            *user.User
}
