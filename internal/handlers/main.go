package handlers

import "github.com/Peterwmoss/LiCa/internal/views"

var templates *views.Templates

func init() {
	templates = views.NewTemplates()
}
