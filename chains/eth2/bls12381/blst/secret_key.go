//go:build ((linux && amd64) || (linux && arm64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) && !blst_disabled
// +build linux,amd64 linux,arm64 darwin,amd64 darwin,arm64 windows,amd64
// +build !blst_disabled

package blst

import (
	"crypto/subtle"
	"fmt"
	common2 "github.com/mapprotocol/atlas/chains/eth2/bls12381/common"
	"github.com/mapprotocol/atlas/chains/eth2/rand"

	blst "github.com/supranational/blst/bindings/go"
)

const BLSSecretKeyLength = 32

// bls12SecretKey used in the BLS signature scheme.
type bls12SecretKey struct {
	p *blst.SecretKey
}

// RandKey creates a new private key using a random method provided as an io.Reader.
func RandKey() (common2.SecretKey, error) {
	// Generate 32 bytes of randomness
	var ikm [32]byte
	_, err := rand.NewGenerator().Read(ikm[:])
	if err != nil {
		return nil, err
	}
	// Defensive check, that we have not generated a secret key,
	secKey := &bls12SecretKey{blst.KeyGen(ikm[:])}
	if IsZero(secKey.Marshal()) {
		return nil, common2.ErrZeroKey
	}
	return secKey, nil
}

// SecretKeyFromBytes creates a BLS private key from a BigEndian byte slice.
func SecretKeyFromBytes(privKey []byte) (common2.SecretKey, error) {
	if len(privKey) != BLSSecretKeyLength {
		return nil, fmt.Errorf("secret key must be %d bytes", BLSSecretKeyLength)
	}
	secKey := new(blst.SecretKey).Deserialize(privKey)
	if secKey == nil {
		return nil, common2.ErrSecretUnmarshal
	}
	wrappedKey := &bls12SecretKey{p: secKey}
	if IsZero(privKey) {
		return nil, common2.ErrZeroKey
	}
	return wrappedKey, nil
}

// PublicKey obtains the public key corresponding to the BLS secret key.
func (s *bls12SecretKey) PublicKey() common2.PublicKey {
	return &PublicKey{p: new(blstPublicKey).From(s.p)}
}

// IsZero checks if the secret key is a zero key.
func IsZero(sKey []byte) bool {
	b := byte(0)
	for _, s := range sKey {
		b |= s
	}
	return subtle.ConstantTimeByteEq(b, 0) == 1
}

// Sign a message using a secret key - in a beacon/validator client.
//
// In IETF draft BLS specification:
// Sign(SK, message) -> signature: a signing algorithm that generates
//      a deterministic signature given a secret key SK and a message.
//
// In Ethereum proof of stake specification:
// def Sign(SK: int, message: Bytes) -> BLSSignature
func (s *bls12SecretKey) Sign(msg []byte) common2.Signature {
	signature := new(blstSignature).Sign(s.p, msg, dst)
	return &Signature{s: signature}
}

// Marshal a secret key into a LittleEndian byte slice.
func (s *bls12SecretKey) Marshal() []byte {
	keyBytes := s.p.Serialize()
	return keyBytes
}
