// Copyright 2023 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplane

import (
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	"github.com/gardener/gardener/extensions/pkg/webhook/controlplane"
	"github.com/gardener/gardener/extensions/pkg/webhook/controlplane/genericmutator"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/operation/botanist/component/extensions/operatingsystemconfig/original/components/kubelet"
	oscutils "github.com/gardener/gardener/pkg/operation/botanist/component/extensions/operatingsystemconfig/utils"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/onmetal/gardener-extension-provider-onmetal/pkg/onmetal"
)

var logger = log.Log.WithName("onmetal-controlplane-webhook")

// AddToManager creates a webhook and adds it to the manager.
func AddToManager(mgr manager.Manager) (*extensionswebhook.Webhook, error) {
	logger.Info("Adding webhook to manager")
	fciCodec := oscutils.NewFileContentInlineCodec()
	return controlplane.New(mgr, controlplane.Args{
		Kind:     controlplane.KindShoot,
		Provider: onmetal.Type,
		Types: []extensionswebhook.Type{
			{Obj: &appsv1.Deployment{}},
			{Obj: &extensionsv1alpha1.OperatingSystemConfig{}},
		},
		Mutator: genericmutator.NewMutator(NewEnsurer(logger), oscutils.NewUnitSerializer(),
			kubelet.NewConfigCodec(fciCodec), fciCodec, logger),
	})
}
