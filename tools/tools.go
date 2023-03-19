// +build tools

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tools

import (
	_ "github.com/AlekSi/gocov-xml"
	_ "github.com/axw/gocov/gocov"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
<<<<<<< HEAD
	_ "github.com/onsi/ginkgo/ginkgo"
=======
	_ "github.com/onsi/ginkgo/v2/ginkgo"
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	_ "github.com/vektra/mockery/cmd/mockery"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "mvdan.cc/gofumpt"
)
