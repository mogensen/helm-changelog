package output

import (
	"text/template"

	_ "embed"

	"github.com/Masterminds/sprig"
)

//go:embed templates/release-template.gtpl
var defaultReleaseTemplate string

//go:embed templates/helpers.gtpl
var defaultHelperTemplates string

func getReleaseTemplate(templateName string) (*template.Template, error) {
	tmpl := template.New(templateName)
	_ = tmpl.Funcs(sprig.FuncMap())
	_, err := tmpl.Parse(defaultHelperTemplates)
	if err != nil {
		return tmpl, err
	}

	_, err = tmpl.Parse(defaultReleaseTemplate)
	if err != nil {
		return tmpl, err
	}

	return tmpl, nil
}
