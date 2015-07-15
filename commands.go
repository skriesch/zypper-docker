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
	"fmt"
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/samalba/dockerclient"
)

func stringInSlice(s string, strs []string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

// TODO: right now this test is a joke. We should be checking the digest ID.
func isSUSE(img *dockerclient.Image) bool {
	allowed := []string{"13.1", "13.2", "harlequin",
		"latest", "tumbleweed", "bottle"}

	for _, repo := range img.RepoTags {
		repoTag := strings.SplitN(repo, ":", 2)
		if len(repoTag) == 2 && repoTag[0] == "opensuse" &&
			stringInSlice(repoTag[1], allowed) {
			return true
		}
	}
	return false
}

func imagesCmd(ctx *cli.Context) {
	client := getDockerClient()

	imgs, err := client.ListImages(false)
	if err != nil {
		log.Printf("%v\n", err)
	} else {
		for _, img := range imgs {
			if isSUSE(img) {
				fmt.Printf("Img: %v\n", img)
			}
		}
	}
}
func listUpdatesCmd(ctx *cli.Context) {
}
func listPatchescmd(ctx *cli.Context) {
}
func patchCmd(ctx *cli.Context) {
}
func patchCheckCmd(ctx *cli.Context) {
}
func psCmd(ctx *cli.Context) {
}
