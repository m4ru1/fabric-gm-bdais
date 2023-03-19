/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmgmt

import (
<<<<<<< HEAD
	"io/ioutil"
=======
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestChainMgmt is a basic sanity check test to catch any errors that could be caused by changes in the ledgermgmt or kvledger packages
func TestChainMgmt(t *testing.T) {
<<<<<<< HEAD
	dataDir, err := ioutil.TempDir("", "ledgerbenchmark_sanitycheck")
	require.NoError(t, err)
	require.NoError(t, os.RemoveAll(dataDir))
	defer os.RemoveAll(dataDir)
=======
	dataDir := t.TempDir()
	require.NoError(t, os.RemoveAll(dataDir))
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19

	mgrConf := &ChainMgrConf{
		DataDir:   dataDir,
		NumChains: 1,
	}
	batchConf := &BatchConf{BatchSize: 1}
	env := InitTestEnv(mgrConf, batchConf, ChainInitOpCreate)
	require.NotNil(t, env.mgr)
	require.Len(t, env.Chains(), 1)
	bcInfo, err := env.Chains()[0].PeerLedger.GetBlockchainInfo()
	require.NoError(t, err)
	require.Equal(t, uint64(1), bcInfo.Height)
}
