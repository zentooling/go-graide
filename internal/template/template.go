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

	contact := root + CONTACT
	index := root + INDEX
	base := root + BASE

	templateMap := make(map[string]*template.Template)
	// templates need to be stitched together, inheritance is not supported
	templateMap[INDEX] = template.Must(template.ParseFiles(base, index))
	templateMap[CONTACT] = template.Must(template.ParseFiles(base, contact))
	return &View{
		templateMap: templateMap,
	}

}
