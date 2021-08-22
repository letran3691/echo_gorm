package handle
import (
	"github.com/labstack/echo/v4"
	"io"
	"html/template"

)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}