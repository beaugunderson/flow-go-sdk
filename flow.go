package flow

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/onflow/flow-go-sdk/crypto"
)

// An Identifier is a 32-byte unique identifier for an entity.
type Identifier [32]byte

// ZeroID is the empty identifier.
var ZeroID = Identifier{}

// Bytes returns the bytes representation of this identifier.
func (i Identifier) Bytes() []byte {
	return i[:]
}

// Hex returns the hexadecimal string representation of this identifier.
func (i Identifier) Hex() string {
	return hex.EncodeToString(i[:])
}

// String returns the string representation of this identifier.
func (i Identifier) String() string {
	return i.Hex()
}

// BytesToID constructs an identifier from a byte slice.
func BytesToID(b []byte) Identifier {
	var id Identifier
	copy(id[:], b)
	return id
}

func HashToID(hash []byte) Identifier {
	return BytesToID(hash)
}

// DefaultHasher is the default hasher used by Flow.
var DefaultHasher crypto.Hasher

func init() {
	DefaultHasher = crypto.NewSHA3_256()
}

func rlpEncode(v interface{}) ([]byte, error) {
	return rlp.EncodeToBytes(v)
}

func rlpDecode(b []byte, v interface{}) error {
	return rlp.DecodeBytes(b, v)
}

func mustRLPEncode(v interface{}) []byte {
	b, err := rlpEncode(v)
	if err != nil {
		panic(err)
	}
	return b
}
