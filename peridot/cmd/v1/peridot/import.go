// Copyright (c) All respective contributors to the Peridot Project. All rights reserved.
// Copyright (c) 2021-2022 Rocky Enterprise Software Foundation, Inc. All rights reserved.
// Copyright (c) 2021-2022 Ctrl IQ, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"github.com/spf13/cobra"
	"log"
	"openapi.peridot.resf.org/peridotopenapi"
	"time"
)

var impCmd = &cobra.Command{
	Use:  "import [name]",
	Args: cobra.ExactArgs(1),
	Run:  impCmdMn,
}

var (
	version string
	release string
)

func init() {
	impCmd.Flags().StringVar(&version, "version", "", "Version of the package to import (if specified, you also need to specify release)")
	impCmd.Flags().StringVar(&release, "release", "", "Release of the package to import (if specified, you also need to specify version)")
	impCmd.Flags().BoolVar(&setInactive, "set-inactive", false, "Set build as inactive")
}

func impCmdMn(_ *cobra.Command, args []string) {
	if (version == "" && release != "") || (version != "" && release == "") {
		log.Fatal("You must specify both version and release")
	}

	// Ensure project id exists
	projectId := mustGetProjectID()

	importCl := getClient(serviceImport).(peridotopenapi.ImportServiceApi)
	body := peridotopenapi.ImportServiceImportPackageBody{
		PackageName: &args[0],
		SetInactive: &setInactive,
	}
	if version != "" {
		body.Vre = &peridotopenapi.V1VersionRelease{
			Version: &version,
			Release: &release,
		}
	}
	req := importCl.ImportPackage(getContext(), projectId).Body(body)
	importRes, _, err := req.Execute()
	errFatal(err)

	// Wait for import to finish
	taskCl := getClient(serviceTask).(peridotopenapi.TaskServiceApi)
	log.Printf("Waiting for import %s to finish\n", importRes.GetTaskId())
	for {
		res, _, err := taskCl.GetTask(getContext(), projectId, importRes.GetTaskId()).Execute()
		if err != nil {
			log.Printf("Error getting task: %s", err.Error())
			time.Sleep(5 * time.Second)
		}
		task := res.GetTask()
		if task.GetDone() {
			if task.GetSubtasks()[0].GetStatus() == peridotopenapi.SUCCEEDED {
				log.Printf("Import %s finished successfully\n", importRes.GetTaskId())
				break
			} else {
				log.Fatalf("Import %s failed with status %s\n", importRes.GetTaskId(), task.GetSubtasks()[0].GetStatus())
			}
		}

		time.Sleep(5 * time.Second)
	}
}
