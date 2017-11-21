// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package boilerplate

import (
	"log"

	"github.com/urfave/cli"
	"github.com/zhuharev/guse/routes/parser"
)

var (
	// Cmd for command usage
	Cmd = cli.Command{
		Name:      "boilerplate",
		ShortName: "bp",
		Subcommands: cli.Commands{
			CmdRoutes,
		},
	}

	// CmdRoutes create routes
	CmdRoutes = cli.Command{
		Name:   "routes",
		Action: routesCmd,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "projectName",
			},
			cli.StringFlag{
				Name: "prefix",
			},
			cli.StringFlag{
				Name:  "routes",
				Value: "routes",
			},
		},
	}
)

func routesCmd(ctx *cli.Context) {
	var (
		projectName = ctx.String("projectName")
		prefix      = ctx.String("prefix")
		routesFile  = ctx.String("routes")
	)

	wr := NewRoutesWriter(Prefix(prefix), ProjectName(projectName))

	routes, err := parser.ParseFile(routesFile)
	if err != nil {
		log.Fatalln(err)
	}

	err = wr.WriteRoutes("routes", routes)
	if err != nil {
		log.Fatalln(err)
	}

	err = wr.WriteRegistreRoutes("./cmd/routes.go", routes)
	if err != nil {
		log.Fatalln(err)
	}
}
