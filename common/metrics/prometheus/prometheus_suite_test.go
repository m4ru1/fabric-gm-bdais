/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package prometheus_test

import (
	"testing"

<<<<<<< HEAD
	. "github.com/onsi/ginkgo"
=======
	. "github.com/onsi/ginkgo/v2"
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	. "github.com/onsi/gomega"
)

func TestPrometheus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prometheus Suite")
}
