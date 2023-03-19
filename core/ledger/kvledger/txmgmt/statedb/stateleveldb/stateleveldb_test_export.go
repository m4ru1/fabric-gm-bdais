/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package stateleveldb

import (
<<<<<<< HEAD
	"io/ioutil"
	"os"
=======
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	"testing"

	"github.com/stretchr/testify/require"
)

// TestVDBEnv provides a level db backed versioned db for testing
type TestVDBEnv struct {
	t          testing.TB
	DBProvider *VersionedDBProvider
	dbPath     string
}

// NewTestVDBEnv instantiates and new level db backed TestVDB
func NewTestVDBEnv(t testing.TB) *TestVDBEnv {
	t.Logf("Creating new TestVDBEnv")
<<<<<<< HEAD
	dbPath, err := ioutil.TempDir("", "statelvldb")
	if err != nil {
		t.Fatalf("Failed to create leveldb directory: %s", err)
	}
=======
	dbPath := t.TempDir()
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	dbProvider, err := NewVersionedDBProvider(dbPath)
	require.NoError(t, err)
	return &TestVDBEnv{t, dbProvider, dbPath}
}

// Cleanup closes the db and removes the db folder
func (env *TestVDBEnv) Cleanup() {
	env.t.Logf("Cleaningup TestVDBEnv")
	env.DBProvider.Close()
<<<<<<< HEAD
	os.RemoveAll(env.dbPath)
=======
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
}
