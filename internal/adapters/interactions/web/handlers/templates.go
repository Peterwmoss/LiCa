package handlers

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/views"
)

var Templates *views.Templates

func init() {
	Templates = views.NewTemplates()
}

