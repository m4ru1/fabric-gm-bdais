/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package inactive_test

import (
	"testing"

	"github.com/m4ru1/fabric-gm-bdais/orderer/common/types"
	"github.com/m4ru1/fabric-gm-bdais/orderer/consensus/inactive"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestInactiveChain(t *testing.T) {
	err := errors.New("foo")
	chain := &inactive.Chain{Err: err}

	require.Equal(t, err, chain.Order(nil, 0))
	require.Equal(t, err, chain.Configure(nil, 0))
	require.Equal(t, err, chain.WaitReady())
	require.NotPanics(t, chain.Start)
	require.NotPanics(t, chain.Halt)
	_, open := <-chain.Errored()
	require.False(t, open)

	cRel, status := chain.StatusReport()
	require.Equal(t, types.ConsensusRelationConfigTracker, cRel)
	require.Equal(t, types.StatusInactive, status)
}
