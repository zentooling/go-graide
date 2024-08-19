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
	BASE        = layout + ".gohtml"
	CONTACT     = "contact.gohtml"
	INDEX       = "index.gohtml"
	INSTITUTION = "institution.gohtml"
	NEW         = "new.gohtml"
	EDIT        = "edit.gohtml"
	LOGIN       = "login.gohtml"
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
	loadTemplate(templateMap, INSTITUTION)
	loadTemplate(templateMap, NEW)
	loadTemplate(templateMap, EDIT)
	loadTemplate(templateMap, LOGIN)
	return &View{
		templateMap: templateMap,
	}

}

func loadTemplate(templateMap map[string]*template.Template, templateName string) {
	// read template file root dir from config file
	rootDir := config.Instance().Template.RootDir
	base := filepath.Join(rootDir, BASE)
	tmplFile := filepath.Join(rootDir, templateName)

	inject := func(in string) string { return "Inject " + in + " Inject" }
	templFuncMap := template.FuncMap{
		"tt": inject,
	}

	tmpl, err := template.New("").Funcs(templFuncMap).ParseFiles(tmplFile, base)
	if err != nil {
		panic(err)
	}
	//tmpl = tmpl.Funcs(templFuncMap)
	templateMap[templateName] = template.Must(tmpl, err)
}
