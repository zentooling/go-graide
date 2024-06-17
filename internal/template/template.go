package template

import (
	"html/template"
	"io"
)

const (
	// private (lowercase)
	root   = "../../web/templates/"
	ext    = ".gtpl"
	layout = "layout"
	// exported
	BASE    = layout + ".html" + ext
	CONTACT = "contact.html" + ext
	INDEX   = "index.html" + ext
	NEW   = "new.html" + ext
	EDIT   = "edit.html" + ext
)

type View struct {
	templateMap map[string]*template.Template
}

func (v View) getTemplate(name string) *template.Template {
	return v.templateMap[name]
}

func (v View) ExecuteTemplate(wr io.Writer, name string, data any) error {
	// the base template is named 'layout'. It is executed as it refers to the 'content' template which is defined
	// in the outer template file
	return v.getTemplate(name).ExecuteTemplate(wr, layout, data)

}

// to be called in 'main' once - TODO make proper singleton
func New() *View {

	templateMap := make(map[string]*template.Template)
	// templates need to be stitched together, inheritance is not supported
	templateMap[INDEX] = newTemplate(INDEX)
	templateMap[CONTACT] = newTemplate(CONTACT)
	templateMap[NEW] = newTemplate(NEW)
	templateMap[EDIT] = newTemplate(EDIT)
	return &View{
		templateMap: templateMap,
	}

}

func newTemplate(templateName string) *template.Template {
	return template.Must(template.ParseFiles(root + BASE, root + templateName))
}
