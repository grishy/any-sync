package signingkey

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/util/keys"
	"io"

	"github.com/anytypeio/go-anytype-infrastructure-experiments/util/strkey"
)

// Ed25519PrivateKey is an ed25519 private key.
type Ed25519PrivateKey struct {
	k ed25519.PrivateKey
}

// Ed25519PublicKey is an ed25519 public key.
type Ed25519PublicKey struct {
	k ed25519.PublicKey
}

func NewSigningEd25519PubKeyFromBytes(bytes []byte) (PubKey, error) {
	return UnmarshalEd25519PublicKey(bytes)
}

func NewSigningEd25519PrivKeyFromBytes(bytes []byte) (PrivKey, error) {
	return UnmarshalEd25519PrivateKey(bytes)
}

func GenerateRandomEd25519KeyPair() (PrivKey, PubKey, error) {
	return GenerateEd25519Key(rand.Reader)
}

// GenerateEd25519Key generates a new ed25519 private and public key pair.
func GenerateEd25519Key(src io.Reader) (PrivKey, PubKey, error) {
	pub, priv, err := ed25519.GenerateKey(src)
	if err != nil {
		return nil, nil, err
	}

	return &Ed25519PrivateKey{
			k: priv,
		},
		&Ed25519PublicKey{
			k: pub,
		},
		nil
}

// Raw private key bytes.
func (k *Ed25519PrivateKey) Raw() ([]byte, error) {
	// The Ed25519 private key contains two 32-bytes curve points, the private
	// key and the public key.
	// It makes it more efficient to get the public key without re-computing an
	// elliptic curve multiplication.
	buf := make([]byte, len(k.k))
	copy(buf, k.k)

	return buf, nil
}

func (k *Ed25519PrivateKey) pubKeyBytes() []byte {
	return k.k[ed25519.PrivateKeySize-ed25519.PublicKeySize:]
}

// Equals compares two ed25519 private keys.
func (k *Ed25519PrivateKey) Equals(o keys.Key) bool {
	edk, ok := o.(*Ed25519PrivateKey)
	if !ok {
		return keys.KeyEquals(k, o)
	}

	return subtle.ConstantTimeCompare(k.k, edk.k) == 1
}

// GetPublic returns an ed25519 public key from a private key.
func (k *Ed25519PrivateKey) GetPublic() PubKey {
	return &Ed25519PublicKey{k: k.pubKeyBytes()}
}

// Sign returns a signature from an input message.
func (k *Ed25519PrivateKey) Sign(msg []byte) ([]byte, error) {
	return ed25519.Sign(k.k, msg), nil
}

// Raw public key bytes.
func (k *Ed25519PublicKey) Raw() ([]byte, error) {
	return k.k, nil
}

// Equals compares two ed25519 public keys.
func (k *Ed25519PublicKey) Equals(o keys.Key) bool {
	edk, ok := o.(*Ed25519PublicKey)
	if !ok {
		return keys.KeyEquals(k, o)
	}

	return bytes.Equal(k.k, edk.k)
}

// Verify checks a signature agains the input data.
func (k *Ed25519PublicKey) Verify(data []byte, sig []byte) (bool, error) {
	return ed25519.Verify(k.k, data, sig), nil
}

// UnmarshalEd25519PublicKey returns a public key from input bytes.
func UnmarshalEd25519PublicKey(data []byte) (PubKey, error) {
	if len(data) != 32 {
		return nil, errors.New("expect ed25519 public key data size to be 32")
	}

	return &Ed25519PublicKey{
		k: ed25519.PublicKey(data),
	}, nil
}

// UnmarshalEd25519PrivateKey returns a private key from input bytes.
func UnmarshalEd25519PrivateKey(data []byte) (PrivKey, error) {
	switch len(data) {
	case ed25519.PrivateKeySize + ed25519.PublicKeySize:
		// Remove the redundant public key. See issue #36.
		redundantPk := data[ed25519.PrivateKeySize:]
		pk := data[ed25519.PrivateKeySize-ed25519.PublicKeySize : ed25519.PrivateKeySize]
		if subtle.ConstantTimeCompare(pk, redundantPk) == 0 {
			return nil, errors.New("expected redundant ed25519 public key to be redundant")
		}

		// No point in storing the extra data.
		newKey := make([]byte, ed25519.PrivateKeySize)
		copy(newKey, data[:ed25519.PrivateKeySize])
		data = newKey
	case ed25519.PrivateKeySize:
	default:
		return nil, fmt.Errorf(
			"expected ed25519 data size to be %d or %d, got %d",
			ed25519.PrivateKeySize,
			ed25519.PrivateKeySize+ed25519.PublicKeySize,
			len(data),
		)
	}

	return &Ed25519PrivateKey{
		k: ed25519.PrivateKey(data),
	}, nil
}

// TODO: remove this one in favor of new one
type Ed25519SigningPubKeyDecoder struct{}

func NewEd25519PubKeyDecoder() PubKeyDecoder {
	return &Ed25519SigningPubKeyDecoder{}
}

func (e *Ed25519SigningPubKeyDecoder) DecodeFromBytes(bytes []byte) (PubKey, error) {
	return NewSigningEd25519PubKeyFromBytes(bytes)
}

func (e *Ed25519SigningPubKeyDecoder) DecodeFromString(identity string) (PubKey, error) {
	pubKeyRaw, err := strkey.Decode(0x5b, identity)
	if err != nil {
		return nil, err
	}

	return e.DecodeFromBytes(pubKeyRaw)
}

func (e *Ed25519SigningPubKeyDecoder) DecodeFromStringIntoBytes(identity string) ([]byte, error) {
	return strkey.Decode(0x5b, identity)
}

func (e *Ed25519SigningPubKeyDecoder) EncodeToString(pubkey PubKey) (string, error) {
	raw, err := pubkey.Raw()
	if err != nil {
		return "", err
	}
	return strkey.Encode(0x5b, raw)
}

func NewEDPrivKeyDecoder() keys.Decoder {
	return keys.NewKeyDecoder(NewSigningEd25519PrivKeyFromBytes)
}

func NewEDPubKeyDecoder() keys.Decoder {
	return keys.NewKeyDecoder(NewSigningEd25519PubKeyFromBytes)
}
