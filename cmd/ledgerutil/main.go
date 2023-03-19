/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

<<<<<<< HEAD
	"github.com/hyperledger/fabric/internal/ledgerutil"
=======
	"github.com/hyperledger/fabric/internal/ledgerutil/compare"
	"github.com/hyperledger/fabric/internal/ledgerutil/identifytxs"
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	compareErrorMessage = "Ledger Compare Error: "
	outputDirDesc       = "Snapshot comparison json results output directory. Default is the current directory."
<<<<<<< HEAD
	firstDiffsDesc      = "Maximum number of differences to record in " + ledgerutil.FirstDiffsByHeight +
		". Requesting a report with many differences may result in a large amount " +
		"of memory usage. Defaults to 10. If set to 0, will not produce " +
		ledgerutil.FirstDiffsByHeight + "."
=======
	firstDiffsDesc      = "Maximum number of differences to record in " + compare.FirstDiffsByHeight +
		". Requesting a report with many differences may result in a large amount of memory usage. Defaults " +
		" to 10. If set to 0, will not produce " + compare.FirstDiffsByHeight + "."
	identifytxsErrorMessage = "Ledger Identify Transactions Error: "
	snapshotDiffsPathDesc   = "Path to json file containing list of target records to search for in transactions. This is typically output " +
		"from ledgerutil compare."
	blockStorePathDesc = "Path to file system of target peer, used to access block store. Defaults to '/var/hyperledger/production'. " +
		"IMPORTANT: If the configuration for target peer's file system path was changed, the new path MUST be provided."
	blockStorePathDefault = "/var/hyperledger/production"
	outputDirIdDesc       = "Location for identified transactions json results output directory. Default is the current directory."
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
)

var (
	app = kingpin.New("ledgerutil", "Ledger Utility Tool")

<<<<<<< HEAD
	compare       = app.Command("compare", "Compare channel snapshots from two different peers.")
	snapshotPath1 = compare.Arg("snapshotPath1", "First ledger snapshot directory.").Required().String()
	snapshotPath2 = compare.Arg("snapshotPath2", "Second ledger snapshot directory.").Required().String()
	outputDir     = compare.Flag("outputDir", outputDirDesc).Short('o').String()
	firstDiffs    = compare.Flag("firstDiffs", firstDiffsDesc).Short('f').Default("10").Int()
=======
	compareApp    = app.Command("compare", "Compare channel snapshots from two different peers.")
	snapshotPath1 = compareApp.Arg("snapshotPath1", "First ledger snapshot directory.").Required().String()
	snapshotPath2 = compareApp.Arg("snapshotPath2", "Second ledger snapshot directory.").Required().String()
	outputDir     = compareApp.Flag("outputDir", outputDirDesc).Short('o').String()
	firstDiffs    = compareApp.Flag("firstDiffs", firstDiffsDesc).Short('f').Default("10").Int()

	identifytxsApp    = app.Command("identifytxs", "Identify potentially divergent transactions.")
	snapshotDiffsPath = identifytxsApp.Arg("snapshotDiffsPath", snapshotDiffsPathDesc).Required().String()
	blockStorePath    = identifytxsApp.Arg("blockStorePath", blockStorePathDesc).Default(blockStorePathDefault).String()
	outputDirId       = identifytxsApp.Flag("outputDir", outputDirIdDesc).Short('o').String()
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19

	args = os.Args[1:]
)

func main() {
	kingpin.Version("0.0.1")

	command, err := app.Parse(args)
	if err != nil {
		kingpin.Fatalf("parsing arguments: %s. Try --help", err)
		return
	}

	// Command logic
	switch command {
<<<<<<< HEAD
	case compare.FullCommand():
=======
	case compareApp.FullCommand():
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19

		// Determine result json file location
		if *outputDir == "" {
			*outputDir, err = os.Getwd()
			if err != nil {
				fmt.Printf("%s%s\n", compareErrorMessage, err)
<<<<<<< HEAD
				return
			}
		}

		count, outputDirPath, err := ledgerutil.Compare(*snapshotPath1, *snapshotPath2, *outputDir, *firstDiffs)
		if err != nil {
			fmt.Printf("%s%s\n", compareErrorMessage, err)
			return
=======
				os.Exit(1)
			}
		}

		count, outputDirPath, err := compare.Compare(*snapshotPath1, *snapshotPath2, *outputDir, *firstDiffs)
		if err != nil {
			fmt.Printf("%s%s\n", compareErrorMessage, err)
			os.Exit(1)
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
		}

		fmt.Print("\nSuccessfully compared snapshots. ")
		if count == -1 {
			fmt.Println("Both snapshot public state and private state hashes were the same. No results were generated.")
		} else {
			fmt.Printf("Results saved to %s. Total differences found: %d\n", outputDirPath, count)
<<<<<<< HEAD
		}
=======
			os.Exit(2)
		}

	case identifytxsApp.FullCommand():

		// Determine result json file location
		if *outputDirId == "" {
			*outputDirId, err = os.Getwd()
			if err != nil {
				fmt.Printf("%s%s\n", identifytxsErrorMessage, err)
				os.Exit(1)
			}
		}

		firstBlock, lastBlock, err := identifytxs.IdentifyTxs(*snapshotDiffsPath, *blockStorePath, *outputDirId)
		if err != nil {
			fmt.Printf("%s%s\n", identifytxsErrorMessage, err)
			os.Exit(1)
		}
		if firstBlock > lastBlock {
			fmt.Printf("%sFirst available block was greater than highest record. No transactions were checked.", identifytxsErrorMessage)
			os.Exit(1)
		}
		fmt.Printf("\nSuccessfully ran identify transactions tool. Transactions were checked between blocks %d and %d.", firstBlock, lastBlock)
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	}
}
