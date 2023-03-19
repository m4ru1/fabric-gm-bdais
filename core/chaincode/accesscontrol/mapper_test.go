/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package accesscontrol

import (
	"testing"
	"time"

	"github.com/m4ru1/fabric-gm-bdais/bccsp"
	"github.com/m4ru1/fabric-gm-bdais/bccsp/sw"
	"github.com/m4ru1/fabric-gm-bdais/common/crypto/tlsgen"
	"github.com/stretchr/testify/require"
)

func TestPurge(t *testing.T) {
	ca, _ := tlsgen.NewCA()
	backupTTL := ttl
	defer func() {
		ttl = backupTTL
	}()
	ttl = time.Second
	m := newCertMapper(ca.NewClientCertKeyPair)
	k, err := m.genCert("A")
	require.NoError(t, err)

	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)

	hash, err := cryptoProvider.Hash(k.TLSCert.Raw, &bccsp.SHA256Opts{})
	require.NoError(t, err)
	require.Equal(t, "A", m.lookup(certHash(hash)))
	time.Sleep(time.Second * 3)
	require.Empty(t, m.lookup(certHash(hash)))
}
