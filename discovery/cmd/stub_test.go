/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package discovery

import (
	"fmt"
	"net"
	"path/filepath"
	"testing"

	"github.com/m4ru1/fabric-gm-bdais/cmd/common"
	"github.com/m4ru1/fabric-gm-bdais/cmd/common/comm"
	"github.com/m4ru1/fabric-gm-bdais/cmd/common/signer"
	discovery "github.com/m4ru1/fabric-gm-bdais/discovery/client"
	corecomm "github.com/m4ru1/fabric-gm-bdais/internal/pkg/comm"
	"github.com/stretchr/testify/require"
)

func TestClientStub(t *testing.T) {
	srv, err := corecomm.NewGRPCServer("127.0.0.1:", corecomm.ServerConfig{
		SecOpts: corecomm.SecureOptions{},
	})
	require.NoError(t, err)
	go srv.Start()
	defer srv.Stop()

	_, portStr, _ := net.SplitHostPort(srv.Address())
	endpoint := fmt.Sprintf("localhost:%s", portStr)
	stub := &ClientStub{}

	req := discovery.NewRequest()

	_, err = stub.Send(endpoint, common.Config{
		SignerConfig: signer.Config{
			MSPID:        "Org1MSP",
			KeyPath:      filepath.Join("testdata", "8150cb2d09628ccc89727611ebb736189f6482747eff9b8aaaa27e9a382d2e93_sk"),
			IdentityPath: filepath.Join("testdata", "cert.pem"),
		},
		TLSConfig: comm.Config{},
	}, req)
	require.Contains(t, err.Error(), "Unimplemented desc = unknown service discovery.Discovery")
}

func TestRawStub(t *testing.T) {
	srv, err := corecomm.NewGRPCServer("127.0.0.1:", corecomm.ServerConfig{
		SecOpts: corecomm.SecureOptions{},
	})
	require.NoError(t, err)
	go srv.Start()
	defer srv.Stop()

	_, portStr, _ := net.SplitHostPort(srv.Address())
	endpoint := fmt.Sprintf("localhost:%s", portStr)
	stub := &RawStub{}

	req := discovery.NewRequest()

	_, err = stub.Send(endpoint, common.Config{
		SignerConfig: signer.Config{
			MSPID:        "Org1MSP",
			KeyPath:      filepath.Join("testdata", "8150cb2d09628ccc89727611ebb736189f6482747eff9b8aaaa27e9a382d2e93_sk"),
			IdentityPath: filepath.Join("testdata", "cert.pem"),
		},
		TLSConfig: comm.Config{},
	}, req)
	require.Contains(t, err.Error(), "Unimplemented desc = unknown service discovery.Discovery")
}
