// Copyright 2022 The Kube-burner Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workloads

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewClusterDensity holds cluster-density workload
func NewClusterDensity(wh *WorkloadHelper) *cobra.Command {
	var iterations int
	cmd := &cobra.Command{
		Use:   "cluster-density <flags>",
		Short: "Runs cluster-density workload",
		PreRun: func(cmd *cobra.Command, args []string) {
			wh.Metadata.Benchmark = cmd.Name()
			os.Setenv("JOB_ITERATIONS", fmt.Sprint(iterations))
		},
		Run: func(cmd *cobra.Command, args []string) {
			wh.run("cluster-density.yml")
		},
	}
	cmd.Flags().IntVar(&iterations, "iterations", 0, "Cluster-density iterations")
	cmd.MarkFlagRequired("iterations")
	return cmd
}
