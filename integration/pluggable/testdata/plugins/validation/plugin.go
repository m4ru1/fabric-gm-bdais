/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	validation "github.com/m4ru1/fabric-gm-bdais/core/handlers/validation/api"
	"github.com/m4ru1/fabric-gm-bdais/core/handlers/validation/builtin"
	"github.com/m4ru1/fabric-gm-bdais/integration/pluggable"
)

// go build -buildmode=plugin -o plugin.so

// NewPluginFactory is the function ran by the plugin infrastructure to create a validation plugin factory.
func NewPluginFactory() validation.PluginFactory {
	pluggable.PublishValidationPluginActivation()
	return &builtin.DefaultValidationFactory{}
}
