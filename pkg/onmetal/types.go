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

package onmetal

import (
	"path/filepath"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

const (
	// ProviderName is the name of the onmetal provider.
	ProviderName = "provider-onmetal"

	// CloudControllerManagerImageName is the name of the cloud-controller-manager image.
	CloudControllerManagerImageName = "cloud-controller-manager"
	// CSIDriverImageName is the name of the csi-driver image.
	CSIDriverImageName = "csi-driver"
	// CSIProvisionerImageName is the name of the csi-provisioner image.
	CSIProvisionerImageName = "csi-provisioner"
	// CSIAttacherImageName is the name of the csi-attacher image.
	CSIAttacherImageName = "csi-attacher"
	// CSIResizerImageName is the name of the csi-resizer image.
	CSIResizerImageName = "csi-resizer"
	// CSINodeDriverRegistrarImageName is the name of the csi-node-driver-registrar image.
	CSINodeDriverRegistrarImageName = "csi-node-driver-registrar"
	// CSILivenessProbeImageName is the name of the csi-liveness-probe image.
	CSILivenessProbeImageName = "csi-liveness-probe"
	// MachineControllerManagerImageName is the name of the MachineControllerManager image.
	MachineControllerManagerImageName = "machine-controller-manager"
	// MachineControllerManagerProviderOnmetalImageName is the name of the MachineController onmetal image.
	MachineControllerManagerProviderOnmetalImageName = "machine-controller-manager-provider-onmetal"

	// UsernameFieldName is the field in a secret where the namespace is stored at.
	UsernameFieldName = "username"
	// NamespaceFieldName is the field in a secret where the namespace is stored at.
	NamespaceFieldName = "namespace"
	// KubeConfigFieldName is containing the effective kubeconfig to access an onmetal cluster.
	KubeConfigFieldName = "kubeconfig"
	// TokenFieldName is containing the token to access an onmetal cluster.
	TokenFieldName = "token"
	// NetworkFieldName is the name of network field
	NetworkFieldName = "networkName"
	// PrefixFieldName is the name of the prefix field
	PrefixFieldName = "prefixName"
	// LabelsFieldName is the name of the labels field
	LabelsFieldName = "labels"
	// UserDataFieldName is the name of the user data field
	UserDataFieldName = "userData"
	// ImageFieldName is the name of the image field
	ImageFieldName = "image"
	// RootDiskFieldName is the name of the root disk field
	RootDiskFieldName = "rootDisk"
	// SizeFieldName is the name of the size field
	SizeFieldName = "size"
	// VolumeClassFieldName is the name of the volume class field
	VolumeClassFieldName = "volumeClassName"
	// ClusterNameLabel is the name is the label key of the cluster name
	ClusterNameLabel = "extension.api.onmetal.de/cluster-name"

	// CloudProviderConfigName is the name of the secret containing the cloud provider config.
	CloudProviderConfigName = "cloud-provider-config"
	// CloudControllerManagerName is a constant for the name of the CloudController deployed by the worker controller.
	CloudControllerManagerName = "cloud-controller-manager"
	// CSIControllerName is a constant for the name of the CSI controller deployment in the seed.
	CSIControllerName = "csi-driver-controller"
	// CSIControllerObservabilityConfigName is the name of the ConfigMap containing monitoring and logging stack configurations for csi-driver.
	CSIControllerObservabilityConfigName = "csi-driver-controller-observability-config"
	// CSINodeName is a constant for the name of the CSI node deployment in the shoot.
	CSINodeName = "csi-driver-node"
	// CSIDriverName is a constant for the name of the csi-driver component.
	CSIDriverName = "csi-driver"
	// CSIProvisionerName is a constant for the name of the csi-provisioner component.
	CSIProvisionerName = "csi-provisioner"
	// CSIAttacherName is a constant for the name of the csi-attacher component.
	CSIAttacherName = "csi-attacher"
	// CSIResizerName is a constant for the name of the csi-resizer component.
	CSIResizerName = "csi-resizer"
	// CSINodeDriverRegistrarName is a constant for the name of the csi-node-driver-registrar component.
	CSINodeDriverRegistrarName = "csi-node-driver-registrar"
	// CSILivenessProbeName is a constant for the name of the csi-liveness-probe component.
	CSILivenessProbeName = "csi-liveness-probe"
	// CSIStorageProvisioner is a constant with the storage provisioner name which is used in storageclasses.
	CSIStorageProvisioner = "onmetal-csi-driver"
	// MachineControllerManagerName is a constant for the name of the machine-controller-manager.
	MachineControllerManagerName = "machine-controller-manager"
	// MachineControllerManagerVpaName is the name of the VerticalPodAutoscaler of the machine-controller-manager deployment.
	MachineControllerManagerVpaName = "machine-controller-manager-vpa"
	// MachineControllerManagerMonitoringConfigName is the name of the ConfigMap containing monitoring stack configurations for machine-controller-manager.
	MachineControllerManagerMonitoringConfigName = "machine-controller-manager-monitoring-config"
)

var (
	// ChartsPath is the path to the charts
	ChartsPath = filepath.Join("charts")
	// InternalChartsPath is the path to the internal charts
	InternalChartsPath = filepath.Join(ChartsPath, "internal")
	// UsernamePrefix is a constant for the username prefix of components deployed by onmetal.
	UsernamePrefix = extensionsv1alpha1.SchemeGroupVersion.Group + ":" + ProviderName + ":"
)
