// email2matrix is an Email (SMTP) server relaying received messages to a Matrix room.
// Copyright (C) 2019 Slavi Pantaleev
//
// https://devture.com/
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"devture-email2matrix/email2matrix/configuration"
	"devture-email2matrix/email2matrix/container"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"fmt"

	guerrilla "github.com/flashmob/go-guerrilla"
	log "github.com/sirupsen/logrus"
)

// Following variables will be statically linked at the time of compiling
// Source: https://oddcode.daveamit.com/2018/08/17/embed-versioning-information-in-golang-binary/

// GitCommit holds short commit hash of source tree
var GitCommit string

// GitBranch holds current branch name the code is built off
var GitBranch string

// GitState shows whether there are uncommitted changes
var GitState string

// GitSummary holds output of git describe --tags --dirty --always
var GitSummary string

// BuildDate holds RFC3339 formatted UTC date (build time)
var BuildDate string

// Version holds contents of ./VERSION file, if exists, or the value passed via the -version option
var Version string

func main() {
	fmt.Printf(`
                      _ _ ___                  _        _      
                     (_) |__ \                | |      (_)     
  ___ _ __ ___   __ _ _| |  ) |_ __ ___   __ _| |_ _ __ ___  __
 / _ \ '_ ' _ \ / _' | | | / /| '_ ' _ \ / _' | __| '__| \ \/ /
|  __/ | | | | | (_| | | |/ /_| | | | | | (_| | |_| |  | |>  < 
 \___|_| |_| |_|\__,_|_|_|____|_| |_| |_|\__,_|\__|_|  |_/_/\_\
----------------------------------------------[ Version: %s ]
GitCommit: %s
GitBranch: %s
GitState: %s
GitSummary: %s
BuildDate: %s

`, Version, GitCommit, GitBranch, GitState, GitSummary, BuildDate)
	configPath := flag.String("config", "config.json", "configuration file to use")
	flag.Parse()

	configuration, err := configuration.LoadConfiguration(*configPath)
	if err != nil {
		panic(err)
	}

	container, shutdownHandler := container.BuildContainer(*configuration)

	d := container.Get("smtp.server").(guerrilla.Daemon)
	err = d.Start()
	if err != nil {
		log.Panic("start error", err)
	}

	channelComplete := make(chan bool)
	setupSignalHandling(
		channelComplete,
		shutdownHandler,
	)

	<-channelComplete
}

func setupSignalHandling(
	channelComplete chan bool,
	shutdownHandler *container.ContainerShutdownHandler,
) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel

		shutdownHandler.Shutdown()

		channelComplete <- true
	}()
}
