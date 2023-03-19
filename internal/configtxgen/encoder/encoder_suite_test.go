/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package encoder_test

import (
	"testing"

	"github.com/hyperledger/fabric/bccsp/factory"
<<<<<<< HEAD
	. "github.com/onsi/ginkgo"
=======
	. "github.com/onsi/ginkgo/v2"
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	. "github.com/onsi/gomega"
)

func TestEncoder(t *testing.T) {
	factory.InitFactories(nil)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Encoder Suite")
}
