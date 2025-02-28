// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config // import "go.opentelemetry.io/collector/config"
import (
	"go.opentelemetry.io/collector/component"
)

// ExporterSettings defines common settings for a component.Exporter configuration.
// Specific exporters can embed this struct and extend it with more fields if needed.
//
// It is highly recommended to "override" the Validate() function.
//
// When embedded in the exporter config, it must be with `mapstructure:",squash"` tag.
type ExporterSettings struct {
	id component.ID `mapstructure:"-"`
	component.ExporterConfig
}

// NewExporterSettings return a new ExporterSettings with the given ComponentID.
func NewExporterSettings(id component.ID) ExporterSettings {
	return ExporterSettings{id: id}
}

var _ component.ExporterConfig = (*ExporterSettings)(nil)

// ID returns the receiver component.ID.
func (es *ExporterSettings) ID() component.ID {
	return es.id
}

// SetIDName sets the receiver name.
func (es *ExporterSettings) SetIDName(idName string) {
	es.id = component.NewIDWithName(es.id.Type(), idName)
}

// Deprecated: [v0.65.0] use component.ValidateConfig.
func (es *ExporterSettings) Validate() error {
	return nil
}
