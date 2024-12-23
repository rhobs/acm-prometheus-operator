// Copyright The prometheus-operator Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// SecretOrConfigMapApplyConfiguration represents a declarative configuration of the SecretOrConfigMap type for use
// with apply.
type SecretOrConfigMapApplyConfiguration struct {
	Secret    *v1.SecretKeySelector    `json:"secret,omitempty"`
	ConfigMap *v1.ConfigMapKeySelector `json:"configMap,omitempty"`
}

// SecretOrConfigMapApplyConfiguration constructs a declarative configuration of the SecretOrConfigMap type for use with
// apply.
func SecretOrConfigMap() *SecretOrConfigMapApplyConfiguration {
	return &SecretOrConfigMapApplyConfiguration{}
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *SecretOrConfigMapApplyConfiguration) WithSecret(value v1.SecretKeySelector) *SecretOrConfigMapApplyConfiguration {
	b.Secret = &value
	return b
}

// WithConfigMap sets the ConfigMap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMap field is set to the value of the last call.
func (b *SecretOrConfigMapApplyConfiguration) WithConfigMap(value v1.ConfigMapKeySelector) *SecretOrConfigMapApplyConfiguration {
	b.ConfigMap = &value
	return b
}
