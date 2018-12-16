// Copyright © 2018 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

func initLogging() {
	// Setup Log Formatter
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Setup and initialise application logging
	initLogging()

	log.Infoln("SkyEye: Skycoin Blockchain Analysis Tool - Starting")
	defer log.Infoln("SkyEye: Stopping")

	// Setup OS Notification for Interrupt or Kill signal - to cleanly terminate the app
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt)
	log.Infoln("Press Ctrl+C to Exit")

	// Wait for the app to be signaled to terminate
	signal := <-osSignal
	log.Debugln("OS Interupt Signal Received. Exiting", signal)
}
