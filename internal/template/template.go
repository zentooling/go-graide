package template

import (
	"github.com/zentooling/graide/internal/config"
	"html/template"
	"io"
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
	return v.templateMap["root"].ExecuteTemplate(wr, name, data)

}

func New() *View {

	templateMap := make(map[string]*template.Template)
	rootDir := config.Instance().Template.RootDir + "/*.gohtml"
	inject := func(in string) string { return "Inject " + in + " Inject" }
	// create a test template function as proof of concept
	templateFuncMap := template.FuncMap{
		"tt": inject,
	}

	tmpl, err := template.New("").Funcs(templateFuncMap).ParseGlob(rootDir)
	if err != nil {
		panic(err)
	}
	//tmpl = tmpl.Funcs(templateFuncMap)
	templateMap["root"] = template.Must(tmpl, err)
	return &View{
		templateMap: templateMap,
	}

}
