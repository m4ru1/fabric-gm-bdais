/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package history

import (
	"crypto/sha256"
<<<<<<< HEAD
	"io/ioutil"
	"os"
=======
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	"testing"

	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/hyperledger/fabric/core/ledger/kvledger/bookkeeping"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/txmgr"
	"github.com/hyperledger/fabric/core/ledger/mock"
	"github.com/stretchr/testify/require"
)

var testHashFunc = func(data []byte) ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

type levelDBLockBasedHistoryEnv struct {
	t                     testing.TB
	testBlockStorageEnv   *testBlockStoreEnv
	testDBEnv             privacyenabledstate.TestEnv
	testBookkeepingEnv    *bookkeeping.TestEnv
	txmgr                 *txmgr.LockBasedTxMgr
	testHistoryDBProvider *DBProvider
	testHistoryDB         *DB
	testHistoryDBPath     string
}

func newTestHistoryEnv(t *testing.T) *levelDBLockBasedHistoryEnv {
	testLedgerID := "TestLedger"

	blockStorageTestEnv := newBlockStorageTestEnv(t)

	testDBEnv := &privacyenabledstate.LevelDBTestEnv{}
	testDBEnv.Init(t)
	testDB := testDBEnv.GetDBHandle(testLedgerID)
	testBookkeepingEnv := bookkeeping.NewTestEnv(t)

<<<<<<< HEAD
	testHistoryDBPath, err := ioutil.TempDir("", "historyldb")
	if err != nil {
		t.Fatalf("Failed to create history database directory: %s", err)
	}
=======
	testHistoryDBPath := t.TempDir()
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19

	txmgrInitializer := &txmgr.Initializer{
		LedgerID:            testLedgerID,
		DB:                  testDB,
		StateListeners:      nil,
		BtlPolicy:           nil,
		BookkeepingProvider: testBookkeepingEnv.TestProvider,
		CCInfoProvider:      &mock.DeployedChaincodeInfoProvider{},
		CustomTxProcessors:  nil,
		HashFunc:            testHashFunc,
	}
	txMgr, err := txmgr.NewLockBasedTxMgr(txmgrInitializer)

	require.NoError(t, err)
	testHistoryDBProvider, err := NewDBProvider(testHistoryDBPath)
	require.NoError(t, err)
	testHistoryDB := testHistoryDBProvider.GetDBHandle("TestHistoryDB")

	return &levelDBLockBasedHistoryEnv{
		t,
		blockStorageTestEnv,
		testDBEnv,
		testBookkeepingEnv,
		txMgr,
		testHistoryDBProvider,
		testHistoryDB,
		testHistoryDBPath,
	}
}

func (env *levelDBLockBasedHistoryEnv) cleanup() {
	env.txmgr.Shutdown()
	env.testDBEnv.Cleanup()
	env.testBlockStorageEnv.cleanup()
	env.testBookkeepingEnv.Cleanup()
	// clean up history
	env.testHistoryDBProvider.Close()
<<<<<<< HEAD
	os.RemoveAll(env.testHistoryDBPath)
=======
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
}

/////// testBlockStoreEnv//////

type testBlockStoreEnv struct {
	t               testing.TB
	provider        *blkstorage.BlockStoreProvider
	blockStorageDir string
}

func newBlockStorageTestEnv(t testing.TB) *testBlockStoreEnv {
<<<<<<< HEAD
	testPath, err := ioutil.TempDir("", "historyleveldb-")
	if err != nil {
		panic(err)
	}
=======
	testPath := t.TempDir()
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	conf := blkstorage.NewConf(testPath, 0)

	attrsToIndex := []blkstorage.IndexableAttr{
		blkstorage.IndexableAttrBlockHash,
		blkstorage.IndexableAttrBlockNum,
		blkstorage.IndexableAttrTxID,
		blkstorage.IndexableAttrBlockNumTranNum,
	}
	indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}

	p, err := blkstorage.NewProvider(conf, indexConfig, &disabled.Provider{})
	require.NoError(t, err)
	return &testBlockStoreEnv{t, p, testPath}
}

func (env *testBlockStoreEnv) cleanup() {
	env.provider.Close()
<<<<<<< HEAD
	env.removeFSPath()
}

func (env *testBlockStoreEnv) removeFSPath() {
	fsPath := env.blockStorageDir
	os.RemoveAll(fsPath)
=======
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
}
