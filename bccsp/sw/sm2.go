package sw

import (
	"crypto/rand"

	"github.com/Hyperledger-TWGC/ccs-gm/sm2"
	"github.com/m4ru1/fabric-gm-bdais/bccsp"
)

type sm2Signer struct{}

func (s *sm2Signer) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return signSm2(k.(*sm2PrivateKey).privKey, digest, opts)
}

type sm2PrivateKeyVerifier struct{}

func (v *sm2PrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return verifySm2(&k.(*sm2PrivateKey).privKey.PublicKey, signature, digest, opts)
}

type sm2PublicKeyVerifier struct{}

func (v *sm2PublicKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return verifySm2(k.(*sm2PublicKey).pubKey, signature, digest, opts)
}

// 签名方法，参考ccs-gm api
func signSm2(k *sm2.PrivateKey, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return k.Sign(rand.Reader, digest, opts)
}

// 验签方法，参考ccs-gm api
func verifySm2(k *sm2.PublicKey, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return k.Verify(digest, signature), nil
}
