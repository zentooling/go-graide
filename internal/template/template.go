package template

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/zentooling/graide/internal/config"
)

const (
	// private (lowercase)
	layout = "layout"
	// exported
	BASE    = layout + ".html.gtpl"
	CONTACT = "contact.html.gtpl"
	INDEX   = "index.html.gtpl"
	NEW     = "new.html.gtpl"
	EDIT    = "edit.html.gtpl"
)

type View struct {
	templateMap map[string]*template.Template
}

func (v View) ExecuteTemplate(wr io.Writer, name string, data any) error {
	// the base template is named 'layout'. It is executed as it refers to the 'content' template which is defined
	// in the outer template file
	return v.templateMap[name].ExecuteTemplate(wr, layout, data)

}

// to be called in 'main' once - TODO make proper singleton
func New() *View {

	templateMap := make(map[string]*template.Template)
	// templates need to be stitched together, inheritance is not supported
	loadTemplate(templateMap, INDEX)
	loadTemplate(templateMap, CONTACT)
	loadTemplate(templateMap, NEW)
	loadTemplate(templateMap, EDIT)
	return &View{
		templateMap: templateMap,
	}

}

func loadTemplate(templateMap map[string]*template.Template, templateName string) {
	// read template file root dir from config file
	rootDir := config.Instance().Template.RootDir
	base := filepath.Join(rootDir, BASE)
	tmplFile := filepath.Join(rootDir, templateName)
	templateMap[templateName] = template.Must(template.ParseFiles(tmplFile, base))
}
