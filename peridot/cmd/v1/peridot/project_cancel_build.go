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
)

var projectCancelBuild = &cobra.Command{
	Use:  "cancel-build [build-id...]",
	Args: cobra.MinimumNArgs(1),
	Run:  projectCancelBuildMn,
}

func projectCancelBuildMn(_ *cobra.Command, args []string) {
	projectID := mustGetProjectID()

	taskCl := getClient(serviceTask).(peridotopenapi.TaskServiceApi)

	for _, buildId := range args {
	  res, _, err := taskCl.GetTask(getContext(), projectID, buildId).Execute()
	  errFatal(err)

	  task := res.GetTask()
	  // Only parent tasks can be cancelled, and must be Pending or Running.
	  // A parent task is task whose first subtask's parent is null
	  subTask := task.GetSubtasks()[0]
	  if subTask.GetParentTaskId() == "0" && (subTask.GetStatus() == peridotopenapi.RUNNING || subTask.GetStatus() == peridotopenapi.PENDING) {
	    // Cancel the task. The response here is just empty
	    _, _, err := taskCl.CancelTask(getContext(), projectID, buildId).Execute()
	    errFatal(err)

	    log.Printf("Build with Task ID %s Successfully Submitted for Cancellation", buildId)
	  }
	}
}
