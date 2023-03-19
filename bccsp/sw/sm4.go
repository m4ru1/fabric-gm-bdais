package sw

import (
	"bytes"
	"github.com/Hyperledger-TWGC/ccs-gm/sm4"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"fabric-gm-bdais\bccsp"
)


type sm4Encryptor struct{}

func (e *sm4Encryptor) Encrypt(k bccsp.Key, plaintext []byte, opts bccsp.EncrypterOpts)([]byte, error){
	switch o := opts.(type){
	case SM4ModeOpts:
		return
	
	}
}

type sm4Decryptor struct{}

func (*sm4Decryptor) Decrypt(k bccsp.Key, ciphertext []byte, opts bccsp.DecrypterOpts)([]byte, error){
	
}


func SM4Encrypt(key, src []byte) ([]byte, error){

}

func SM4EncryptWithRand(key, src []byte) ([]byte, error){

}

func SM4EncryptWithIV(key, src []byte) ([]byte, error){

}

func SM4Decrypt(key, src []byte) ([]byte, error){

}

func sm4EncryptWithRand(key, src []byte) ([]byte, error){

}

func sm4EncryptWithIV(key, src []byte) ([]byte, error){

}

func sm4PkcsPadding(src []byte) []byte {

}

func sm4PkcsUnPadding(src []byte) []byte {

}