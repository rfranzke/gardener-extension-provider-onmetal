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

package cmd

import (
	"github.com/onmetal/gardener-extension-provider-onmetal/pkg/admission/validator"

	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"
)

// GardenWebhookSwitchOptions are the webhookcmd.SwitchOptions for the admission webhooks.
func GardenWebhookSwitchOptions() *webhookcmd.SwitchOptions {
	return webhookcmd.NewSwitchOptions(
		webhookcmd.Switch(validator.Name, validator.New),
		webhookcmd.Switch(validator.SecretsValidatorName, validator.NewSecretsWebhook),
	)
}
