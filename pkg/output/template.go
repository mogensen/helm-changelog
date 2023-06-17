package output

import (
	"path"
	"text/template"

	_ "embed"

	"github.com/Masterminds/sprig/v3"
)

//go:embed templates/release-template.gtpl
var defaultReleaseTemplate string

//go:embed templates/helpers.gtpl
var defaultHelperTemplates string

func getReleaseTemplate(templateName, releaseTemplatePath string) (tmpl *template.Template, err error) {
	if releaseTemplatePath != "" {
		// templates created by ParseFiles are names after the basename of the file
		templateName = path.Base(releaseTemplatePath)
	}
	tmpl = template.New(templateName)
	_ = tmpl.Funcs(sprig.TxtFuncMap())
	_, err = tmpl.Parse(defaultHelperTemplates)
	if err != nil {
		return
	}

	if releaseTemplatePath == "" {
		_, err = tmpl.Parse(defaultReleaseTemplate)
	} else {
		_, err = tmpl.ParseFiles(releaseTemplatePath)
	}
	if err != nil {
		return
	}

	return
}
