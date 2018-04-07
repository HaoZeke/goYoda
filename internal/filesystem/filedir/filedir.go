// Pandoc handler
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

package filedir

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func modeTest() {
	name := "FileOrDir"
	fi, err := os.Stat(name)
	if err != nil {
		log.Error(err)
		return
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		log.Info("Passed a directory")
	case mode.IsRegular():
		// do file stuff
		log.Infor("Passed a single file")
	}
}
