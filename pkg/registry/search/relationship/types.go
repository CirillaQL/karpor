// Copyright The Karbour Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package relationship

type RelationshipGraph struct {
	RelationshipNodes []*RelationshipGraphNode `json:"relationship,omitempty" yaml:"relationship,omitempty"`
}

type RelationshipGraphNode struct {
	Group         string          `json:"group,omitempty" yaml:"group,omitempty"`
	Version       string          `json:"version,omitempty" yaml:"version,omitempty"`
	Kind          string          `json:"kind,omitempty" yaml:"kind,omitempty"`
	Parent        []*Relationship `json:"parent,omitempty" yaml:"parent,omitempty"`
	Children      []*Relationship `json:"children,omitempty" yaml:"children,omitempty"`
	ResourceCount int             `json:"resourceCount,omitempty" yaml:"resource_count,omitempty"`
}

type Relationship struct {
	ParentNode    *RelationshipGraphNode
	ChildNode     *RelationshipGraphNode
	AutoGenerated bool
	Group         string              `json:"group,omitempty" yaml:"group,omitempty"`
	Version       string              `json:"version,omitempty" yaml:"version,omitempty"`
	Kind          string              `json:"kind,omitempty" yaml:"kind,omitempty"`
	ClusterScoped bool                `json:"clusterScoped,omitempty" yaml:"cluster_scoped,omitempty"`
	Type          string              `json:"type,omitempty" yaml:"type,omitempty"`
	SelectorPath  string              `json:"selectorPath,omitempty" yaml:"selector_path,omitempty"`
	JSONPath      []map[string]string `json:"jsonPath,omitempty" yaml:"json_path,omitempty"`
}

type ResourceGraphNode struct {
	Name      string
	Namespace string
	Group     string
	Version   string
	Kind      string
	Parent    []*ResourceGraphNode
	Children  []*ResourceGraphNode
}