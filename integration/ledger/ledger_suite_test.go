/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ledger

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/integration"
	"github.com/hyperledger/fabric/integration/nwo"
<<<<<<< HEAD
	. "github.com/onsi/ginkgo"
=======
	. "github.com/onsi/ginkgo/v2"
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	. "github.com/onsi/gomega"
)

func TestLedger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ledger Suite")
}

var (
	buildServer *nwo.BuildServer
	components  *nwo.Components
)

var _ = SynchronizedBeforeSuite(func() []byte {
	buildServer = nwo.NewBuildServer()
	buildServer.Serve()

	components = buildServer.Components()
	payload, err := json.Marshal(components)
	Expect(err).NotTo(HaveOccurred())

	return payload
}, func(payload []byte) {
	err := json.Unmarshal(payload, &components)
	Expect(err).NotTo(HaveOccurred())

	flogging.SetWriter(GinkgoWriter)
})

var _ = SynchronizedAfterSuite(func() {
}, func() {
	buildServer.Shutdown()
})

func StartPort() int {
	return integration.LedgerPort.StartPortForNode()
}
