/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

<<<<<<< HEAD
	"io/ioutil"
	"os"
=======
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	"os/exec"
	"testing"
	"time"
)

func TestArugments(t *testing.T) {

	testCases := map[string]struct {
		exitCode int
		args     []string
	}{
		"toofew": {
			exitCode: 1,
			args:     []string{"metadatadir"},
		},
		"twomany": {
			exitCode: 1,
			args:     []string{"wibble", "wibble", "wibble"},
		},
		"validtype": {
			exitCode: 0,
			args:     []string{"na", "testdata/validtype"},
		},
		"wrongtype": {
			exitCode: 1,
			args:     []string{"na", "testdata/wrongtype"},
		},
	}

	// Build ledger binary
	gt := NewWithT(t)
	detectCmd, err := gexec.Build("github.com/hyperledger/fabric/ccaas_builder/cmd/detect")
	gt.Expect(err).NotTo(HaveOccurred())
	defer gexec.CleanupBuildArtifacts()

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			cmd := exec.Command(detectCmd, testCase.args...)
			session, err := gexec.Start(cmd, nil, nil)
			gt.Expect(err).NotTo(HaveOccurred())
			gt.Eventually(session, 5*time.Second).Should(gexec.Exit(testCase.exitCode))
		})
	}
}

func TestMissingFile(t *testing.T) {
	gt := NewWithT(t)

<<<<<<< HEAD
	testPath, err := ioutil.TempDir("", "empty-test-dir")
	gt.Expect(err).NotTo(HaveOccurred())
	defer os.RemoveAll(testPath)
=======
	testPath := t.TempDir()
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19

	detectCmd, err := gexec.Build("github.com/hyperledger/fabric/ccaas_builder/cmd/detect")
	gt.Expect(err).NotTo(HaveOccurred())
	defer gexec.CleanupBuildArtifacts()

	cmd := exec.Command(detectCmd, []string{"na", testPath}...)
	session, err := gexec.Start(cmd, nil, nil)
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Eventually(session, 5*time.Second).Should(gexec.Exit(1))
	gt.Expect(session.Err).To(gbytes.Say("metadata.json not found"))
}
