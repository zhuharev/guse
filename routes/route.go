package routes

import (
	"path/filepath"
	"strings"
)

// Route main struct
type Route struct {
	Method  string
	Name    string
	Handler string
}

type Routes []Route

func (routes Routes) Packages(prefixes ...string) (res []string) {
	var dup = map[string]struct{}{}
	for _, v := range routes {
		pkgName := strings.Split(v.Handler, ".")[0]
		dup[pkgName] = struct{}{}
	}

	prefix := ""
	if len(prefixes) > 0 {
		prefix = prefixes[0]
	}

	for pkg := range dup {
		res = append(res, filepath.Join(prefix, pkg))
	}

	return
}
