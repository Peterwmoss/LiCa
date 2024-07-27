package views

import (
	"bytes"
	"html/template"
	"io"
	"log/slog"
)

type Templates struct {
	Template *template.Template
}

func (t *Templates) Render(writer io.Writer, name string, data any) error {
	var buf bytes.Buffer

	err := t.Template.ExecuteTemplate(&buf, name, data)
	if err != nil {
		slog.Error("Broken", "error", err)
		return err
	}

	_, err = buf.WriteTo(writer)
	return err
}

func NewTemplates() *Templates {
	return &Templates{
		Template: template.Must(template.ParseGlob("internal/adapters/interactions/web/views/*.html")),
	}
}
