// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package skeleton

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

var (
	// DefaultPermissions permissions for all created files
	DefaultPermissions os.FileMode = 0777
)

var (
	// Cmd helper for cli usage
	Cmd = cli.Command{
		Name:      "skeleton",
		Aliases:   []string{"skelet"},
		Action:    cmdAction,
		Usage:     "guse skeleton [config [target dir]]",
		UsageText: "if filename not present, will be used .skeleton file",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "dir",
			},
		},
	}
)

func cmdAction(ctx *cli.Context) {
	var (
		fname     = ".skeleton"
		targetDir = "."
	)

	if ctx.NArg() > 0 {
		fname = ctx.Args().Get(0)
	}

	if ctx.NArg() > 1 {
		targetDir = ctx.Args().Get(1)
	}

	if dir := ctx.String("dir"); dir != "" {
		targetDir = dir
	}

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fpath := strings.TrimSpace(scanner.Text())
		if fpath == "" {
			continue
		}
		hasTrailingSlash := fpath[len(fpath)-1] == '/'
		targetPath := filepath.Join(targetDir, fpath)
		if hasTrailingSlash {
			targetPath += "/"
		}
		err = create(targetPath)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func create(fpath string) (err error) {
	var (
		isDir   = fpath[len(fpath)-1] == '/'
		dirName = fpath
	)

	if !isDir {
		dirName = filepath.Dir(fpath)
		err = os.MkdirAll(dirName, DefaultPermissions)
		if err != nil {
			return
		}

		_, err = os.Create(fpath)
		if err != nil {
			return
		}
		return
	}

	err = os.MkdirAll(dirName, DefaultPermissions)
	if err != nil {
		return
	}

	return
}
