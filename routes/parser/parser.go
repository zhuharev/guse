// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package parser

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/zhuharev/guse/routes"
)

// Parse reader and return routes
func Parse(r io.Reader) (rts routes.Routes, err error) {
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		if scan.Err() == io.EOF {
			break
		}
		if strings.TrimSpace(scan.Text()) == "" {
			continue
		}
		// ignore comments
		if scan.Bytes()[0] == '#' {
			continue
		}
		fields := strings.Fields(strings.TrimSpace(scan.Text()))
		if len(fields) < 3 {
			color.Red("len < 3: %s", scan.Text())
			continue
		}
		rts = append(rts, routes.Route{fields[0], fields[1], fields[2]})
	}
	return
}

// ParseFile open and parse file
func ParseFile(fname string) (rts routes.Routes, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	return Parse(f)
}
