/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package dockercontroller_test

import (
	"github.com/m4ru1/fabric-gm-bdais/core/container/dockercontroller"
)

//go:generate counterfeiter -o mock/platform_builder.go --fake-name PlatformBuilder . platformBuilder
type platformBuilder interface {
	dockercontroller.PlatformBuilder
}
