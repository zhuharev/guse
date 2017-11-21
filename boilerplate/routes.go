package boilerplate

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhuharev/guse/routes"
)

var (
	// fow web
	//  ctx.HTML(200, "{{ $.packageName }}/{{ .Lower }}")
	routeTmpl = `{{ .license }}

package {{ .packageName }}

import (
  {{range .imports}}"{{ . }}"
{{ end }}
)

{{ range .funcs }}
// {{ .Name }} is {{ $.packageName }}.{{ .Lower }} controller
func {{ .Name }}(ctx *context.Context) {

	ctx.HTML(200, "{{ $.packageName }}/{{ .Lower }}")
}
{{ end }}`
)

type RoutesWriter struct {
	projectName string
	// prefix after goroot
	prefix string
}

func NewRoutesWriter(opts ...func(*RoutesWriter)) *RoutesWriter {
	rw := &RoutesWriter{}
	for _, v := range opts {
		v(rw)
	}
	return rw
}

func Prefix(prefix string) func(rw *RoutesWriter) {
	return func(rw *RoutesWriter) {
		rw.prefix = prefix
	}
}

func ProjectName(projectName string) func(rw *RoutesWriter) {
	return func(rw *RoutesWriter) {
		rw.projectName = projectName
	}
}

func (rw *RoutesWriter) WriteRoutes(routesDir string, routes []routes.Route) (err error) {
	for _, route := range routes {
		err = rw.WriteRoute(routesDir, route)
		if err != nil {
			return err
		}
	}
	return
}

type Func struct {
	Name string
}

func (f Func) Lower() string {
	return strings.ToLower(f.Name)
}

func NewFunc(name string) Func {
	return Func{Name: name}
}

func (rw *RoutesWriter) WriteRoute(routesDir string, route routes.Route) (err error) {

	splitted := strings.Split(route.Handler, ".")
	if len(splitted) != 2 {
		return fmt.Errorf("route handler must be subpackage func")
	}

	var (
		pkgName        = splitted[0]
		fileName       = strings.ToLower(splitted[1]) + ".go"
		targetDir      = filepath.Join(routesDir, pkgName)
		targetFileName = filepath.Join(targetDir, fileName)
	)

	err = os.MkdirAll(targetDir, 0777)
	if err != nil {
		return
	}

	f, err := os.OpenFile(targetFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return
	}

	tpl, err := template.New("lol").Parse(routeTmpl)
	if err != nil {
		return
	}

	var imports = []string{
		filepath.Join(rw.prefix, rw.projectName, "web", "context"),
		filepath.Join(rw.prefix, rw.projectName, "models"),
	}

	var funcs = []Func{
		NewFunc(splitted[1]),
	}

	var data = map[string]interface{}{
		"license": `// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.`,
		"packageName": pkgName,
		"imports":     imports,
		"funcs":       funcs,
	}

	return tpl.Execute(f, data)
}

var (
	routesTmpl = `{{ .license }}

package cmd

import(
  "{{ .frameworkPkg }}"
  {{ range .routes.Packages .prefix}}"{{ . }}"
	{{end}}
)

func registreRoutes({{ .frameworkShortName }} *{{ .frameworkPkgName }}.{{ .frameworkStruct }}) {
    {{range .routes}}{{ $.frameworkShortName }}.{{.Method}}("{{.Name}}", {{ .Handler }})
		{{end}}
}`
)

// WriteRegistreRoutes write routes register file
func (rw *RoutesWriter) WriteRegistreRoutes(targetPath string, rts routes.Routes) (err error) {
	f, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return
	}

	var data = map[string]interface{}{
		"license": `// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.`,
		"routes":             rts,
		"prefix":             filepath.Join(rw.prefix, rw.projectName, "routes"),
		"frameworkPkg":       "gopkg.in/macaron.v1", //"github.com/zhuharev/tamework"
		"frameworkPkgName":   "macaron",
		"frameworkShortName": "m", // tw
		"frameworkStruct":    "Macaron",
	}

	tpl, err := template.New("lol").Parse(routesTmpl)
	if err != nil {
		return
	}

	return tpl.Execute(f, data)
}
