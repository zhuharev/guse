// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/zhuharev/guse/boilerplate"
	"github.com/zhuharev/guse/skeleton"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			skeleton.Cmd,
			boilerplate.Cmd,
		},
	}
	app.Run(os.Args)
}
