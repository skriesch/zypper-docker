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

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

// logFileName stores the name of the log file when logging is not done through
// the standard output.
const logFileName = ".zypper-docker.log"

// setupLogger picks the proper output file for this application.
func setupLogger(ctx *cli.Context) error {
	// If the debug flag is set, just print the log to stdout.
	if ctx.GlobalBool("debug") {
		log.SetOutput(os.Stdout)
		return nil
	}

	// Try to set the log inside of the HOME directory. If this is not
	// possible, just use stdout.
	path := filepath.Join(os.Getenv("HOME"), logFileName)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Printf("Could not open log file: %v\n", err)
	} else {
		log.SetOutput(file)
	}
	return nil
}
