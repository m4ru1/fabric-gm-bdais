/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package etcdraft_test

import (
	"testing"

	"github.com/m4ru1/fabric-gm-bdais/bccsp/sw"
	"github.com/m4ru1/fabric-gm-bdais/common/metrics/disabled"
	"github.com/m4ru1/fabric-gm-bdais/internal/pkg/comm"
	"github.com/m4ru1/fabric-gm-bdais/orderer/common/cluster"
	"github.com/m4ru1/fabric-gm-bdais/orderer/common/localconfig"
	"github.com/m4ru1/fabric-gm-bdais/orderer/consensus/etcdraft"
	"github.com/m4ru1/fabric-gm-bdais/orderer/consensus/etcdraft/mocks"
	"github.com/stretchr/testify/require"
)

func TestNewEtcdRaftConsenter(t *testing.T) {
	srv, err := comm.NewGRPCServer("127.0.0.1:0", comm.ServerConfig{})
	require.NoError(t, err)
	defer srv.Stop()
	dialer := &cluster.PredicateDialer{}
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	consenter := etcdraft.New(dialer,
		&localconfig.TopLevel{},
		comm.ServerConfig{
			SecOpts: comm.SecureOptions{
				Certificate: []byte{1, 2, 3},
			},
		},
		srv,
		&mocks.ChainManager{},
		&mocks.InactiveChainRegistry{},
		&disabled.Provider{},
		cryptoProvider,
	)

	// Assert that the certificate from the gRPC server was passed to the consenter
	require.Equal(t, []byte{1, 2, 3}, consenter.Cert)
	// Assert that all dependencies for the consenter were populated
	require.NotNil(t, consenter.Communication)
	require.NotNil(t, consenter.ChainManager)
	require.NotNil(t, consenter.ChainSelector)
	require.NotNil(t, consenter.Dispatcher)
	require.NotNil(t, consenter.Logger)
}

func TestNewEtcdRaftConsenterNoSystemChannel(t *testing.T) {
	srv, err := comm.NewGRPCServer("127.0.0.1:0", comm.ServerConfig{})
	require.NoError(t, err)
	defer srv.Stop()
	dialer := &cluster.PredicateDialer{}
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	consenter := etcdraft.New(
		dialer,
		&localconfig.TopLevel{},
		comm.ServerConfig{
			SecOpts: comm.SecureOptions{
				Certificate: []byte{1, 2, 3},
			},
		},
		srv,
		&mocks.ChainManager{},
		nil, // without a system channel we have InactiveChainRegistry == nil
		&disabled.Provider{},
		cryptoProvider,
	)

	// Assert that the certificate from the gRPC server was passed to the consenter
	require.Equal(t, []byte{1, 2, 3}, consenter.Cert)
	// Assert that all dependencies for the consenter were populated
	require.NotNil(t, consenter.Communication)
	require.NotNil(t, consenter.ChainManager)
	require.NotNil(t, consenter.ChainSelector)
	require.NotNil(t, consenter.Dispatcher)
	require.NotNil(t, consenter.Logger)
	require.Nil(t, consenter.InactiveChainRegistry)
}
