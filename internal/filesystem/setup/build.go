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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

var AppFs = afero.NewOsFs()

func CreateProj(projName string) {
	log.Info("Hi, so you're getting a new project")
	log.Info("You said you want the project to be called " + projName)
	createDirs(projName)
}

func createDirs(projName string) {
	_, err := afero.IsEmpty(AppFs, projName)
	if err != nil {
		AppFs.MkdirAll((projName + "/src/test"), 0755)
	} else {
		log.Info(projName + " is not empty or it exists")
		log.Warning("I die now with error")
		log.Fatal("Delete the " + projName + " folder and try again")
	}
}
