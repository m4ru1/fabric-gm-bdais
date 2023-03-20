/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package main is the entrypoint for the orderer binary
// and calls only into the server.Main() function.  No other
// function should be included in this package.
package main

import "github.com/m4ru1/fabric-gm-bdais/orderer/common/server"

func main() {
	server.Main()
}
