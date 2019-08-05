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

	guerrilla "github.com/flashmob/go-guerrilla"
	log "github.com/sirupsen/logrus"
)

func main() {
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
