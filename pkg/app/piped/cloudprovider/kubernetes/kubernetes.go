// Copyright 2022 The PipeCD Authors.
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

package kubernetes

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

const (
	LabelManagedBy            = "pipecd.dev/managed-by"             // Always be piped.
	LabelPiped                = "pipecd.dev/piped"                  // The id of piped handling this application.
	LabelApplication          = "pipecd.dev/application"            // The application this resource belongs to.
	LabelCommitHash           = "pipecd.dev/commit-hash"            // Hash value of the deployed commit.
	LabelResourceKey          = "pipecd.dev/resource-key"           // The resource key generated by apiVersion, namespace and name. e.g. apps/v1/Deployment/namespace/demo-app
	LabelOriginalAPIVersion   = "pipecd.dev/original-api-version"   // The api version defined in git configuration. e.g. apps/v1
	LabelIgnoreDriftDirection = "pipecd.dev/ignore-drift-detection" // Whether the drift detection should ignore this resource.
	LabelSyncReplace          = "pipecd.dev/sync-by-replace"        // Use replace instead of apply.
	AnnotationConfigHash      = "pipecd.dev/config-hash"            // The hash value of all mouting config resources.
	ManagedByPiped            = "piped"
	IgnoreDriftDetectionTrue  = "true"
	UseReplaceEnabled         = "enabled"

	kustomizationFileName = "kustomization.yaml"
)
