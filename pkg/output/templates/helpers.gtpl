{{- define "release.header" -}}
  {{- if .Chart.Deprecated -}}
    {{- printf "## %s (DEPRECATED)\n" .Chart.Version -}}
  {{- else -}}
    {{- printf "## %s\n" .Chart.Version -}}
  {{- end -}}
{{- end -}}

{{- define "release.releaseDate" -}}
  {{- with .ReleaseDate -}}
    {{- printf "**Release date:** %s\n" (date "2006-01-02" .) -}}
    {{/*
    {{- printf "**Release date:** %s\n" . -}}
     */}}
  {{- end -}}
{{- end -}}

{{- define "release.commits" -}}
  {{- range $commit := .Commits -}}
    {{- printf "* %s\n" $commit.Subject -}}
  {{- end -}}
{{- end -}}

{{- define "release.valuesDiffs" -}}
  {{- printf "### Default value changes\n\n" -}}
  {{- printf "```diff\n" -}}
  {{- if .ValueDiff -}}
    {{- printf "%s\n" .ValueDiff -}}
  {{- else -}}
    {{- printf "# No changes in this release\n" -}}
  {{- end -}}
  {{- printf "```\n" -}}
{{- end -}}

{{- define "release.badges" -}}
  {{- template "release.badges.appVersion" . }}
  {{- template "release.badges.kubeVersion" . }}
  {{- template "release.badges.helmVersion" . }}
{{- end -}}

{{- define "release.badges.appVersion" -}}
  {{- with .Chart.AppVersion -}}
    {{- printf "![AppVersion: %[1]s](https://img.shields.io/static/v1?label=AppVersion&message=%[1]s&color=success)\n" . -}}
  {{- end -}}
{{- end -}}

{{- define "release.badges.kubeVersion" -}}
  {{- with .Chart.KubeVersion -}}
    {{- printf "![Kubernetes: %[1]s](https://img.shields.io/static/v1?label=Kubernetes&message=%[1]s&color=informational&logo=kubernetes)\n" . -}}
  {{- end -}}
{{- end -}}

{{- define "release.badges.helmVersion" -}}
  {{- if eq .Chart.APIVersion "v1" -}}
    {{- printf "![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)\n" -}}
    {{- printf "![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)\n" -}}
  {{- else if eq .Chart.APIVersion "v2" -}}
    {{- printf "![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)\n" -}}
  {{- else -}}
    {{- printf "![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)\n" -}}
  {{- end -}}
{{- end -}}
