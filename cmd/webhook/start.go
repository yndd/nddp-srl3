/*
Copyright 2022 NDD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	_ "net/http/pprof"

	"github.com/spf13/cobra"

	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/nddp-srl3/internal/webhook/server"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	webhookAddr string
)

// startCmd represents the start command for the network device driver
var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "start the srl3 ndd provider webhook",
	Long:         "start the srl3 ndd provider webhook",
	Aliases:      []string{"start"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		zlog := zap.New(zap.UseDevMode(debug), zap.JSONEncoder())

		zlog.Info("create webhook")

		s := server.New(
			webhookAddr,
			server.WithLogger(logging.NewLogrLogger(zlog.WithName("webhook"))),
		)

		return s.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&webhookAddr, "webhook-bind-address", "m", ":9443", "The address the webhook endpoint binds to.")
}
