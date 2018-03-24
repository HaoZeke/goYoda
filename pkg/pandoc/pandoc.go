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
	"strings"

	"github.com/rjeczalik/notify"
	log "github.com/sirupsen/logrus"
)

// RunPandocListener takes a directory to listen to, then should print the
// different changes that occur within it
func RunPandocListener(directory string) {
	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 1)

	// Set up a watchpoint listening for events within a directory tree rooted
	// at current working directory. Dispatch remove events to c.
	if err := notify.Watch(directory, c, notify.InCloseWrite, notify.InMovedTo); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	// Block until an event is received.
	for {
		ei := <-c
		log.Info("Got event:", ei)
		handleFileChange(ei)
		if strings.Contains(ei.Path(), "spooky-action") {
			CompileAndRefresh("spooky-action")
		}
	}
}

// handFileChange handles the file listener event, checks if its on a *.md file
// and is the final change (mac makes various different file changes when rewriting
// a file using MacVim.
func handleFileChange(event notify.EventInfo) {
	fmt.Println(event)
}

// CompileAndRefresh recompiles the given *.md file into *.pdf, refocuses
// Preview.app so that it picks up the changes on disk, then refocuses
// MacVim.app to continue editing.
func CompileAndRefresh(baseFilename string) {
	var err error
	err = compileMarkdownToPdf(baseFilename)
	if err != nil {
		log.Fatal("Could not compile markdown to pdf using base filename: " + baseFilename)
	}
	err = openPreview(baseFilename)
	if err != nil {
		log.Fatal("Could not open a pdf viewer")
	}
	err = openMacVim()
	if err != nil {
		log.Fatal("Could not open an editor")
	}
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
	input := fmt.Sprintf("%s.md", baseFilename)
	output := fmt.Sprintf("%s.pdf", baseFilename)
	cmd := exec.Command(pandocPath, input, "-o", output)
	return cmd.Run()

}

// openPreview uses mac's command open to refocus Preview.app
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

// openMacVim uses mac's command open to refocus MacVim.app
func openMacVim() error {
	openPath, err := exec.LookPath("subl")
	if err != nil {
		log.Fatal("Could not find an installation of open")
	}
	cmd := exec.Command(openPath, "subl")
	return cmd.Run()
}
