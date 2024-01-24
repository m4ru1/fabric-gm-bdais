/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	"crypto/ecdsa"
	"crypto/rsa"

	//"github.com/m4ru1/fabric-gm-bdais/pkg/ccs-gm/x509"
	"errors"
	"fmt"
	"reflect"

	"github.com/m4ru1/fabric-gm-bdais/bccsp"
	"github.com/m4ru1/fabric-gm-bdais/pkg/ccs-gm/sm2"
	"github.com/m4ru1/fabric-gm-bdais/pkg/ccs-gm/utils"
	"github.com/m4ru1/fabric-gm-bdais/pkg/ccs-gm/x509"
)

type aes256ImportKeyOptsKeyImporter struct{}

func (*aes256ImportKeyOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	aesRaw, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}

	if aesRaw == nil {
		return nil, errors.New("Invalid raw material. It must not be nil.")
	}

	if len(aesRaw) != 32 {
		return nil, fmt.Errorf("Invalid Key Length [%d]. Must be 32 bytes", len(aesRaw))
	}

	return &aesPrivateKey{aesRaw, false}, nil
}

type hmacImportKeyOptsKeyImporter struct{}

func (*hmacImportKeyOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	aesRaw, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}

	if len(aesRaw) == 0 {
		return nil, errors.New("Invalid raw material. It must not be nil.")
	}

	return &aesPrivateKey{aesRaw, false}, nil
}

type ecdsaPKIXPublicKeyImportOptsKeyImporter struct{}

func (*ecdsaPKIXPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	der, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}

	if len(der) == 0 {
		return nil, errors.New("Invalid raw. It must not be nil.")
	}

	lowLevelKey, err := derToPublicKey(der)
	if err != nil {
		return nil, fmt.Errorf("Failed converting PKIX to ECDSA public key [%s]", err)
	}

	ecdsaPK, ok := lowLevelKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Failed casting to ECDSA public key. Invalid raw material.")
	}

	return &ecdsaPublicKey{ecdsaPK}, nil
}

type ecdsaPrivateKeyImportOptsKeyImporter struct{}

func (*ecdsaPrivateKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	der, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("[ECDSADERPrivateKeyImportOpts] Invalid raw material. Expected byte array.")
	}

	if len(der) == 0 {
		return nil, errors.New("[ECDSADERPrivateKeyImportOpts] Invalid raw. It must not be nil.")
	}

	lowLevelKey, err := derToPrivateKey(der)
	if err != nil {
		return nil, fmt.Errorf("Failed converting PKIX to ECDSA public key [%s]", err)
	}

	ecdsaSK, ok := lowLevelKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("Failed casting to ECDSA private key. Invalid raw material.")
	}

	return &ecdsaPrivateKey{ecdsaSK}, nil
}

type ecdsaGoPublicKeyImportOptsKeyImporter struct{}

func (*ecdsaGoPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	lowLevelKey, ok := raw.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *ecdsa.PublicKey.")
	}

	return &ecdsaPublicKey{lowLevelKey}, nil
}

type x509PublicKeyImportOptsKeyImporter struct {
	bccsp *CSP
}

func (ki *x509PublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	x509Cert, ok := raw.(*x509.Certificate)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *x509.Certificate.")
	}

	pk := x509Cert.PublicKey

	switch pk := pk.(type) {
	case *ecdsa.PublicKey:
		return ki.bccsp.KeyImporters[reflect.TypeOf(&bccsp.ECDSAGoPublicKeyImportOpts{})].KeyImport(
			pk,
			&bccsp.ECDSAGoPublicKeyImportOpts{Temporary: opts.Ephemeral()})
	case *rsa.PublicKey:
		// This path only exists to support environments that use RSA certificate
		// authorities to issue ECDSA certificates.
		return &rsaPublicKey{pubKey: pk}, nil
	case *sm2.PublicKey:
		// 思路就是通过反射找到对应的导入器，对该密钥内容进行导入
		return ki.bccsp.KeyImporters[reflect.TypeOf(&bccsp.SM2PublicKeyImportOpts{})].KeyImport(pk, &bccsp.SM2PublicKeyImportOpts{Temporary: opts.Ephemeral()})
	case *sm2.PrivateKey:
		// 思路就是通过反射找到对应的导入器，对该密钥内容进行导入
		return ki.bccsp.KeyImporters[reflect.TypeOf(&bccsp.SM2PrivateKeyImportOpts{})].KeyImport(pk, &bccsp.SM2PrivateKeyImportOpts{Temporary: opts.Ephemeral()})
	default:
		return nil, errors.New("Certificate's public key type not recognized. Supported keys: [ECDSA, RSA]")
	}
}

type sm4ImportKeyOptsKeyImporter struct{}

func (*sm4ImportKeyOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	sm4KeyRaw, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}
	if sm4KeyRaw == nil {
		return nil, errors.New("Invalid raw material. It must not be nil.")
	}
	if len(sm4KeyRaw) != 16 {
		return nil, fmt.Errorf("Invalid Key Length [%d]. Must be 16 bytes", len(sm4KeyRaw))
	}
	return &sm4PrivateKey{sm4KeyRaw, false}, nil
}

type sm2PrivateKeyOptsKeyImporter struct{}

func (*sm2PrivateKeyOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (k bccsp.Key, err error) {
	der, ok := raw.([]byte)
	// 类型转换判断
	if !ok {
		return nil, errors.New("Invalid raw material, Expected byte array")
	}
	// 密钥长度判断
	if len(der) != 0 {
		return nil, errors.New("Invalid raw material, It must botbe nil")
	}
	key, err := utils.PEMtoPrivateKey(der, nil)
	// 密钥导入结果判断
	if err != nil {
		return nil, errors.New("Failed converting to GMSM2 private key")
	}
	return &sm2PrivateKey{key}, nil
}

type sm2PublicKeyOptsKeyImporter struct{}

func (*sm2PublicKeyOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (k bccsp.Key, err error) {
	der, ok := raw.([]byte)
	// 类型转换判断
	if !ok {
		return nil, errors.New("Invalid raw material, Expected byte array")
	}
	// 密钥长度判断
	if len(der) != 0 {
		return nil, errors.New("Invalid raw material, It must botbe nil")
	}
	key, err := utils.PEMtoPublicKey(der, nil)
	// 密钥导入结果判断
	if err != nil {
		return nil, errors.New("Failed converting to GMSM2 public key")
	}
	return &sm2PublicKey{key}, nil
}
