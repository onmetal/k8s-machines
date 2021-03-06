/*
 * Copyright (c) 2020 by The metal-stack Authors.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/onmetal/k8s-machines/pkg/client/machines/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BaseBoardManagementControllerInfos returns a BaseBoardManagementControllerInfoInformer.
	BaseBoardManagementControllerInfos() BaseBoardManagementControllerInfoInformer
	// DHCPLeases returns a DHCPLeaseInformer.
	DHCPLeases() DHCPLeaseInformer
	// MachineInfos returns a MachineInfoInformer.
	MachineInfos() MachineInfoInformer
	// MachineTypes returns a MachineTypeInformer.
	MachineTypes() MachineTypeInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BaseBoardManagementControllerInfos returns a BaseBoardManagementControllerInfoInformer.
func (v *version) BaseBoardManagementControllerInfos() BaseBoardManagementControllerInfoInformer {
	return &baseBoardManagementControllerInfoInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DHCPLeases returns a DHCPLeaseInformer.
func (v *version) DHCPLeases() DHCPLeaseInformer {
	return &dHCPLeaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineInfos returns a MachineInfoInformer.
func (v *version) MachineInfos() MachineInfoInformer {
	return &machineInfoInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineTypes returns a MachineTypeInformer.
func (v *version) MachineTypes() MachineTypeInformer {
	return &machineTypeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
