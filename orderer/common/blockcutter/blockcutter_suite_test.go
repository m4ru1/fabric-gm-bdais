/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockcutter_test

import (
	"testing"

	"github.com/m4ru1/fabric-gm-bdais/common/channelconfig"
	"github.com/m4ru1/fabric-gm-bdais/common/metrics"
	"github.com/m4ru1/fabric-gm-bdais/orderer/common/blockcutter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate counterfeiter -o mock/metrics_histogram.go --fake-name MetricsHistogram . metricsHistogram
type metricsHistogram interface {
	metrics.Histogram
}

//go:generate counterfeiter -o mock/metrics_provider.go --fake-name MetricsProvider . metricsProvider
type metricsProvider interface {
	metrics.Provider
}

//go:generate counterfeiter -o mock/config_fetcher.go --fake-name OrdererConfigFetcher . ordererConfigFetcher
type ordererConfigFetcher interface {
	blockcutter.OrdererConfigFetcher
}

//go:generate counterfeiter -o mock/orderer_config.go --fake-name OrdererConfig . ordererConfig
type ordererConfig interface {
	channelconfig.Orderer
}

func TestBlockcutter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blockcutter Suite")
}
