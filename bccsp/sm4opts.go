package bccsp

import "io"

// Notice that both IV and PRNG can be nil. In that case, the BCCSP implementation
// is supposed to sample the IV using a cryptographic secure PRNG.
// Notice also that either IV or PRNG can be different from nil.
type SM4ModeOpts struct {
	IV   []byte
	PRNG io.Reader
}
