// goYoda pandoc driven documentation system.
// Copyright (C) 2017-Present  Rohit Goswami <rohit.goswami@aol.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"

	"github.com/HaoZeke/goYoda/internal/filesystem/setup"
	"github.com/HaoZeke/goYoda/pkg/pandoc"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "goYoda"
	app.Usage = "Pandoc for turtles with go."

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get the current working directory")
	}

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "tex, t",
			Usage: "Uses latexmk and generates an intermediate .tex file",
		},
		cli.BoolFlag{
			Name:  "edit, e",
			Usage: "Opens an HTML rendered viewer [single file ony]",
		},
	}

	// app.Action = func(c *cli.Context) error {
	// 	switch {
	// 	case c.Bool("new"):
	// 		setup.CreateProj(c.Args().First())
	// 	case c.Bool("b"):
	// 		log.Info("Running pandoc listener...")
	// 		pandoc.DirWatcher(wd + "/" + c.Args().First())
	// 	}
	// 	return nil
	// }

	app.Commands = []cli.Command{
		{
			Name:      "background",
			ShortName: "b",
			Usage:     "Watch and compile a file or directory",
			Action: func(c *cli.Context) error {
				log.Info("Running pandoc listener...")
				pandoc.DirWatcher(wd + "/" + c.Args().First())
				return nil
			},
		},
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "Create a new goYoda project",
			Action: func(c *cli.Context) error {
				setup.CreateProj(c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
