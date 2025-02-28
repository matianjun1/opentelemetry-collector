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

package loggingexporter

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestUnmarshalDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NoError(t, component.UnmarshalExporterConfig(confmap.New(), cfg))
	assert.Equal(t, factory.CreateDefaultConfig(), cfg)
}

func TestUnmarshalConfig(t *testing.T) {
	tests := []struct {
		filename    string
		cfg         *Config
		expectedErr string
	}{
		{
			filename: "config_loglevel.yaml",
			cfg: &Config{
				ExporterSettings:   config.NewExporterSettings(component.NewID(typeStr)),
				LogLevel:           zapcore.DebugLevel,
				Verbosity:          configtelemetry.LevelDetailed,
				SamplingInitial:    10,
				SamplingThereafter: 50,
				warnLogLevel:       true,
			},
		},
		{
			filename: "config_verbosity.yaml",
			cfg: &Config{
				ExporterSettings:   config.NewExporterSettings(component.NewID(typeStr)),
				LogLevel:           zapcore.InfoLevel,
				Verbosity:          configtelemetry.LevelDetailed,
				SamplingInitial:    10,
				SamplingThereafter: 50,
			},
		},
		{
			filename: "loglevel_info.yaml",
			cfg: &Config{
				ExporterSettings:   config.NewExporterSettings(component.NewID(typeStr)),
				LogLevel:           zapcore.InfoLevel,
				Verbosity:          configtelemetry.LevelNormal,
				SamplingInitial:    2,
				SamplingThereafter: 500,
				warnLogLevel:       true,
			},
		},
		{
			filename:    "invalid_verbosity_loglevel.yaml",
			expectedErr: "'loglevel' and 'verbosity' are incompatible. Use only 'verbosity' instead",
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			cm, err := confmaptest.LoadConf(filepath.Join("testdata", tt.filename))
			require.NoError(t, err)
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()
			err = component.UnmarshalExporterConfig(cm, cfg)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.cfg, cfg)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		cfg         *Config
		expectedErr string
	}{
		{
			name: "verbosity none",
			cfg: &Config{
				Verbosity: configtelemetry.LevelNone,
			},
			expectedErr: "verbosity level \"None\" is not supported",
		},
		{
			name: "verbosity detailed",
			cfg: &Config{
				Verbosity: configtelemetry.LevelDetailed,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
