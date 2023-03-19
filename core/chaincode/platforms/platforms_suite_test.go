/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package platforms_test

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/platforms"
<<<<<<< HEAD
	. "github.com/onsi/ginkgo"
=======
	. "github.com/onsi/ginkgo/v2"
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	. "github.com/onsi/gomega"
)

//go:generate counterfeiter -o mock/platform.go --fake-name Platform . platform
type platform interface {
	platforms.Platform
}

//go:generate counterfeiter -o mock/package_writer.go --fake-name PackageWriter . packageWriter
type packageWriter interface {
	platforms.PackageWriter
}

func TestPlatforms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Platforms Suite")
}
