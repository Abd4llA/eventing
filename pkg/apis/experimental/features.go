/*
Copyright 2021 The Knative Authors

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

package experimental

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

// Flags is a map containing all the enabled/disabled flags for the experimental features.
// Missing entry in the map means feature is equal to feature not enabled.
type Flags map[string]bool

// IsEnabled returns true if the feature is enabled
func (e Flags) IsEnabled(featureName string) bool {
	return e != nil && e[featureName]
}

// NewFlagsConfigFromMap creates a Flags from the supplied Map
func NewFlagsConfigFromMap(data map[string]string) (Flags, error) {
	flags := Flags{}

	for k, v := range data {
		if strings.HasPrefix(k, "_") {
			// Ignore all the keys starting with _
			continue
		}
		sanitizedKey := strings.TrimSpace(k)
		if strings.EqualFold(v, "true") {
			flags[sanitizedKey] = true
		} else if strings.EqualFold(v, "false") {
			flags[sanitizedKey] = false
		} else {
			return Flags{}, fmt.Errorf("cannot parse the boolean flag '%s' = '%s'. Allowed values: [true, false]", k, v)
		}
	}

	return flags, nil
}

// NewFlagsConfigFromConfigMap creates a Flags from the supplied configMap
func NewFlagsConfigFromConfigMap(config *corev1.ConfigMap) (Flags, error) {
	return NewFlagsConfigFromMap(config.Data)
}
