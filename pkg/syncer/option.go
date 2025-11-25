package syncer

import (
	"context"

	"github.com/kgateway-dev/kgateway/v2/pkg/agentgateway/translator"
	"github.com/kgateway-dev/kgateway/v2/pkg/reports"
)

type StatusSyncerConfig struct {
	CustomStatusSync func(ctx context.Context, rm reports.ReportMap)
}

type StatusSyncerOption func(*StatusSyncerConfig)

func ProcessStatusSyncerOptions(opts ...StatusSyncerOption) *StatusSyncerConfig {
	cfg := &StatusSyncerConfig{}
	for _, fn := range opts {
		fn(cfg)
	}
	return cfg
}

func WithCustomStatusSync(customSync func(ctx context.Context, rm reports.ReportMap)) StatusSyncerOption {
	return func(cfg *StatusSyncerConfig) {
		cfg.CustomStatusSync = customSync
	}
}

type AgentGatewaySyncerConfig struct {
	GatewayTransformationFunc translator.GatewaysTransformationFunction
}

type AgentGatewaySyncerOption func(*AgentGatewaySyncerConfig)

func ProcessAgentGatewaySyncerOptions(opts ...AgentGatewaySyncerOption) *AgentGatewaySyncerConfig {
	cfg := &AgentGatewaySyncerConfig{}
	for _, fn := range opts {
		fn(cfg)
	}
	return cfg
}

func WithGatewayForDeployerTransformationFunc(f translator.GatewaysTransformationFunction) AgentGatewaySyncerOption {
	return func(o *AgentGatewaySyncerConfig) {
		o.GatewayTransformationFunc = f
	}
}
