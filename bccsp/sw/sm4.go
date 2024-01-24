package sw

import (
	"fmt"

	"github.com/m4ru1/fabric-gm-bdais/bccsp"
	"github.com/m4ru1/fabric-gm-bdais/pkg/ccs-gm/sm4"
)

type sm4Encryptor struct{}

func (e *sm4Encryptor) Encrypt(k bccsp.Key, plaintext []byte, opts bccsp.EncrypterOpts) ([]byte, error) {
	switch o := opts.(type) {
	case *bccsp.SM4ModeOpts:
		// GM directly encrypt easily, don't think about iv, prng, or padding
		return sm4.Sm4Cbc(k.(*sm4PrivateKey).privKey, plaintext, sm4.ENC)
	case bccsp.SM4ModeOpts:
		return e.Encrypt(k, plaintext, &o)
	default:
		return nil, fmt.Errorf("Mode not recognized [%s]", opts)
	}
}

type sm4Decryptor struct{}

func (*sm4Decryptor) Decrypt(k bccsp.Key, ciphertext []byte, opts bccsp.DecrypterOpts) ([]byte, error) {
	switch opts.(type) {
	case *bccsp.SM4ModeOpts, bccsp.SM4ModeOpts:
		return sm4.Sm4Cbc(k.(*sm4PrivateKey).privKey, ciphertext, sm4.DEC)
	default:
		return nil, fmt.Errorf("Mode not recognized [%s]", opts)
	}
}
