// Copyright 2022 OnMetal authors
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

package app

import (
	"context"
	"fmt"

	admissioncmd "github.com/onmetal/gardener-extension-provider-onmetal/pkg/admission/cmd"
	onmetalinstall "github.com/onmetal/gardener-extension-provider-onmetal/pkg/apis/onmetal/install"
	provideronmetal "github.com/onmetal/gardener-extension-provider-onmetal/pkg/onmetal"

	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	"github.com/gardener/gardener/extensions/pkg/util"
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"
	"github.com/gardener/gardener/pkg/apis/core/install"
	gardenerhealthz "github.com/gardener/gardener/pkg/healthz"
	"github.com/spf13/cobra"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/component-base/version/verflag"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var log = logf.Log.WithName("gardener-extension-admission-onmetal")

// NewAdmissionCommand creates a new command for running a onmetal gardener-extension-admission-onmetal webhook.
func NewAdmissionCommand(ctx context.Context) *cobra.Command {
	var (
		restOpts = &controllercmd.RESTOptions{}
		mgrOpts  = &controllercmd.ManagerOptions{
			WebhookServerPort:  443,
			MetricsBindAddress: ":8080",
			HealthBindAddress:  ":8081",
		}

		webhookSwitches = admissioncmd.GardenWebhookSwitchOptions()
		webhookOptions  = webhookcmd.NewAddToManagerSimpleOptions(webhookSwitches)

		aggOption = controllercmd.NewOptionAggregator(
			restOpts,
			mgrOpts,
			webhookOptions,
		)
	)

	cmd := &cobra.Command{
		Use: fmt.Sprintf("admission-%s", provideronmetal.Type),

		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.PrintAndExitIfRequested()

			if err := aggOption.Complete(); err != nil {
				return fmt.Errorf("error completing options: %v", err)
			}

			util.ApplyClientConnectionConfigurationToRESTConfig(&componentbaseconfig.ClientConnectionConfiguration{
				QPS:   100.0,
				Burst: 130,
			}, restOpts.Completed().Config)

			mgr, err := manager.New(restOpts.Completed().Config, mgrOpts.Completed().Options())
			if err != nil {
				return fmt.Errorf("could not instantiate manager: %v", err)
			}

			install.Install(mgr.GetScheme())

			if err := onmetalinstall.AddToScheme(mgr.GetScheme()); err != nil {
				return fmt.Errorf("could not update manager scheme: %v", err)
			}

			log.Info("Setting up webhook server")

			if err := webhookOptions.Completed().AddToManager(mgr); err != nil {
				return err
			}

			if err := mgr.AddReadyzCheck("informer-sync", gardenerhealthz.NewCacheSyncHealthz(mgr.GetCache())); err != nil {
				return fmt.Errorf("could not add readycheck for informers: %w", err)
			}

			if err := mgr.AddHealthzCheck("ping", healthz.Ping); err != nil {
				return fmt.Errorf("could not add healthcheck: %w", err)
			}

			if err := mgr.AddReadyzCheck("webhook-server", mgr.GetWebhookServer().StartedChecker()); err != nil {
				return fmt.Errorf("could not add readycheck of webhook to manager: %w", err)
			}

			return mgr.Start(ctx)
		},
	}

	verflag.AddFlags(cmd.Flags())
	aggOption.AddFlags(cmd.Flags())

	return cmd
}
