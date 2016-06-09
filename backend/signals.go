// Copyright (c) 2015 SUSE LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/SUSE/zypper-docker/logger"
)

// KillChannel is a channel which will receive a boolean whenever a signal has
// been received and we want to shut down gracefully.
var KillChannel chan bool

// listenSignals executes a background goroutine that listens to all signals
// and propagates SIGINT, SIGTSTP and SIGTERM to the backend.KillChannel
// channel.
func listenSignals() {
	KillChannel = make(chan bool)
	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTSTP,
				syscall.SIGTERM:
				logger.Printf("signal '%v' received: shutting down gracefully", sig)
				KillChannel <- true
			default:
				logger.Printf("signal '%v' not handled: doing nothing", sig)
			}
		}
	}()
}