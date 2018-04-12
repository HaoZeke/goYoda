// Handle filesystem events
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

// Build Constraints [see https://stackoverflow.com/questions/19847594/how-to-reliably-detect-os-platform-in-go]

// +build windows

package watch

import (
	"github.com/HaoZeke/goYoda/pkg/pandoc"
	"github.com/rjeczalik/notify"
	log "github.com/sirupsen/logrus"
)

// dirWatcher takes a directory to listen to, then should print the
// different changes that occur within it
func DirWatcher(directory string) {
	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 1)

	// Set up a watchpoint listening for events within a directory tree rooted
	// at current working directory. Dispatch remove events to c.
	// Windows doesn't recognize InCloseWrite and InMoveTo
	if err := notify.Watch(directory, c, notify.Create, notify.Remove); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	// Block until an event is received.
	for {
		ei := <-c
		log.Info("Got event:", ei)
		pandoc.HandleFileChanges(ei)
	}
}
