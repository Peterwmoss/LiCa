package templates

import (
	"html/template"
	"io"
)

type Templates struct {
  Template *template.Template
}

func (t *Templates) Render(writer io.Writer, name string, data any) error {
  return t.Template.ExecuteTemplate(writer, name, data)

}

func NewTemplates() *Templates {
  return &Templates{
    Template: template.Must(template.ParseGlob("internal/views/*.html")),
  }
}
