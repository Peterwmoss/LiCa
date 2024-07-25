package handlers

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/views"
)

var templates *views.Templates

func init() {
	templates = views.NewTemplates()
}

