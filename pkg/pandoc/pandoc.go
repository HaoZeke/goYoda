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

package pandoc

import (
	"fmt"
	"os/exec"
	fp "path/filepath"
	"strings"

	"github.com/rjeczalik/notify"
	log "github.com/sirupsen/logrus"
)

// handFileChange handles the file listener event, checks if its on a *.md file
// and is the final change (makes various different file changes when rewriting
// a file.
func handleFileChanges(event notify.EventInfo) {
	switch {
	case strings.Contains(event.Path(), ".md"):
		markdownChanges(event)
	}
}

func markdownChanges(event notify.EventInfo) {
	file, err := fp.Glob(event.Path())
	if err != nil {
		log.Fatal(err)
	}
	// strings.Join(file, " ") gives the filename
	CompileAndRefresh(strings.TrimSuffix(strings.Join(file, " "), fp.Ext(strings.Join(file, " "))))
}

// CompileAndRefresh recompiles the given *.md file into *.pdf, refocuses
// so that it picks up the changes on disk, then refocuses
// to continue editing.
func CompileAndRefresh(baseFilename string) {
	var err error
	err = compileMarkdownToPdf(baseFilename)
	if err != nil {
		log.Info("Could not compile markdown to pdf using base filename: " + baseFilename)
		log.Error(err)
	}
	// err = openPreview(baseFilename)
	// if err != nil {
	// 	log.Fatal("Could not open a pdf viewer")
	// }
	// err = openEditor(baseFilename)
	// if err != nil {
	// 	log.Fatal("Could not open an editor")
	// }
}

// FindFile is given a filename, then it attempts to find where that file is
// to return a full path. It first tries just the filename, then the current
// working directory plus the filename. If the file can't be found, return
// an error
func FindFile(baseFilename string) error {
	return nil
}

// compileMarkdownToPdf takes in the baseFilename, then compiles the *.md
// file to *.pdf
func compileMarkdownToPdf(baseFilename string) error {
	pandocPath, err := exec.LookPath("pandoc")
	if err != nil {
		log.Fatal("Could not find an installation of pandoc")
	}
	texPath, err := exec.LookPath("latexmk")
	if err != nil {
		log.Fatal("Could not find an installation of latexmk")
	}
	inpMd := fmt.Sprintf("%s.md", baseFilename)
	outTex := fmt.Sprintf("%s.tex", baseFilename)
	cmdTex := exec.Command(pandocPath, inpMd, "--standalone", "-o", outTex)
	errTex := cmdTex.Run()
	if errTex != nil {
		log.Fatal(err)
	}
	config, err := fp.Glob("d/.latexmkrc")
	if err != nil {
		log.Fatal(err)
	}
	configTex, err := fp.Abs(strings.Join(config, " "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(configTex)
	cmdPdf := exec.Command(texPath, "-silent", "-f", "-r", configTex, outTex)
	return cmdPdf.Run()

}

// openPreview uses open to refocus the viewwer
// (or open the file if its not open)
func openPreview(baseFilename string) error {
	openPath, err := exec.LookPath("okular")
	if err != nil {
		log.Fatal("Could not find an installation of xdg-open")
	}
	file := fmt.Sprintf("%s.pdf", baseFilename)
	cmd := exec.Command(openPath, file)
	return cmd.Run()
}

// openEditor uses open to refocus the editor
func openEditor(baseFilename string) error {
	openPath, err := exec.LookPath("subl")
	if err != nil {
		log.Fatal("Could not find an installation of open")
	}
	file := fmt.Sprintf("%s.md", baseFilename)
	cmd := exec.Command(openPath, file)
	return cmd.Run()
}
