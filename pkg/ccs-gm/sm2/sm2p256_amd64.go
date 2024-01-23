// SPDX-License-Identifier: Apache-2.0
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the Go wrapper for the constant-time, 64-bit assembly
// implementation of sm2 curve. The optimizations performed here are described in
// detail in:
// S.Gueron and V.Krasnov, "Fast prime field elliptic-curve cryptography with
//                          256-bit primes"
// http://link.springer.com/article/10.1007%2Fs13389-014-0090-x
// https://eprint.iacr.org/2013/816.pdf

//go:build amd64
// +build amd64

package sm2

import (
	"crypto/elliptic"
	"math/big"

	"github.com/dustinxie/ecc"
)

type (
	p256Curve struct {
		*ecc.CurveParams
		// *elliptic.CurveParams
	}
)

var (
	p256 p256Curve
)

func initP256() {
	// // See FIPS 186-3, section D.2.3
	p256.CurveParams = &ecc.CurveParams{}
	p256.Name = "bcy256"
	p256.P, _ = new(big.Int).SetString("F21D860022F6FCD43E1F53A2A2CFEFF7823BD5430E0000BFD7B22DFFE71B2F49", 16)
	p256.N, _ = new(big.Int).SetString("306C4E0006FE3290D939772086F6633187FF35E883E434618BDD7BF21A9F91B9", 16)
	p256.B, _ = new(big.Int).SetString("8445D72302DEF7C8827AEC9808111498AC6BBB9CAD948A68A5FF116A2C0285D1", 16)
	p256.Gx, _ = new(big.Int).SetString("EC45179388F6E8E92E688A368F5D09E26D3129DEDCAC5C88EB6531B8B3272BE5", 16)
	p256.Gy, _ = new(big.Int).SetString("2D611DE19E2CBCD3C5C27046056B9AEEBE2BAF5BD95E4871FCF1235BB3F0677E", 16)
	p256.BitSize = 256
	p256.A, _ = new(big.Int).SetString("A8B8A1E70A28B7770D396A55163701C389CBDC72D616295689664AE93E58F4CE", 16)
	// SM2PARAM_A = A.Bytes()
	// precomputeOnce.Do(initTable)

	// // See FIPS 186-3, section D.2.3
	// p256.CurveParams = &ecc.CurveParams{Name: "SM2-P-256"}
	// p256.P, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	// p256.N, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	// p256.B, _ = new(big.Int).SetString("28E9FA9E9D9F5E344D5A9E4BCF6509A7F39789F515AB8F92DDBCBD414D940E93", 16)
	// p256.Gx, _ = new(big.Int).SetString("32C4AE2C1F1981195F9904466A39C9948FE30BBFF2660BE1715A4589334C74C7", 16)
	// p256.Gy, _ = new(big.Int).SetString("BC3736A2F4F6779C59BDCEE36B692153D0A9877CC62A474002DF32E52139F0A0", 16)
	// p256.BitSize = 256
	// A, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFC", 16)
	// SM2PARAM_A = A.Bytes()
	// precomputeOnce.Do(initTable)
}

func (curve p256Curve) Params() *elliptic.CurveParams {
	p := curve.CurveParams.CurveParams

	return &p
}

func P256CurveA() *big.Int {
	return p256.A
}
