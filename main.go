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
	"fmt"
	"log"
	"os"

	"github.com/HaoZeke/goYoda/pkg/pandoc"
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
	app.Commands = []cli.Command{
		{
			Name:      "background",
			ShortName: "b",
			Usage:     "Start service which listens on input folder or current working directory, detecting changes to *.md and recompiling them to *.pdf",
			Action: func(c *cli.Context) {
				fmt.Println("Running pandoc listener...")
				pandoc.RunPandocListener(wd + "/" + c.Args().First())
			},
		},
		{
			Name:      "compile",
			ShortName: "c",
			Usage:     "Compile given file (without .md extension) to pdf",
			Action: func(c *cli.Context) {
				pandoc.CompileAndRefresh(wd + "/" + c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}
