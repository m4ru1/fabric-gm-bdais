/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"os"

	"github.com/m4ru1/fabric-gm-bdais/bccsp/factory"
	"github.com/m4ru1/fabric-gm-bdais/cmd/common"
	discovery "github.com/m4ru1/fabric-gm-bdais/discovery/cmd"
)

func main() {
	factory.InitFactories(nil)
	cli := common.NewCLI("discover", "Command line client for fabric discovery service")
	discovery.AddCommands(cli)
	cli.Run(os.Args[1:])
}
