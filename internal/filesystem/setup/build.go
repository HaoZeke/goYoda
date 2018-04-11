// Project builder
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

package setup

import (
	"bufio"
	"strings"

	"github.com/gobuffalo/packr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

var AppFs = afero.NewOsFs()
var afs = &afero.Afero{Fs: AppFs}
var box = packr.NewBox("../../starters")

func CreateProj(projName string) {
	log.Info("Setting project up")
	log.Info("Project Name: " + projName)
	createDirs(projName)
	createFiles(projName)
}

func createDirs(projName string) {
	_, err := afero.IsEmpty(AppFs, projName)
	if err != nil {
		AppFs.MkdirAll((projName + "/src/conf"), 0755)
		AppFs.MkdirAll((projName + "/src/filters"), 0755)
		AppFs.MkdirAll((projName + "/src/img"), 0755)
		AppFs.MkdirAll((projName + "/src/md"), 0755)
		AppFs.MkdirAll((projName + "/src/templates"), 0755)
		AppFs.MkdirAll((projName + "/src/tex"), 0755)
	} else {
		log.Info(projName + " is not empty or it exists")
		log.Fatal("Delete the " + projName + " folder and try again")
	}
}

func createFiles(projName string) {
	// Convert the []string to have newlines
	stringSlices := strings.Join(box.List(), "\n")

	// Add a file containing every filename in the box
	box.AddBytes("filego", []byte(stringSlices))

	// Readin the box contents
	file, err := box.Open("filego")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Buffered read of file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// For every line in the file create a corresponding file
		s := box.Bytes(scanner.Text())
		afs.WriteFile((projName + "/" + scanner.Text()), s, 0755)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
