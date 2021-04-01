// Package dshelp provides utilities for parsing and creating
// datastore keys used by go-ipfs
package dshelp

import (
	"github.com/daotl/go-datastore/key"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-base32"
	mh "github.com/multiformats/go-multihash"
)

// NewStrKeyFromBinary creates a new StrKey from a byte slice using Base32 encoding.
func NewStrKeyFromBinary(rawKey []byte) key.StrKey {
	buf := make([]byte, 1+base32.RawStdEncoding.EncodedLen(len(rawKey)))
	buf[0] = '/'
	base32.RawStdEncoding.Encode(buf[1:], rawKey)
	return key.RawStrKey(string(buf))
}

// BinaryFromDsKey returns the byte slice corresponding to the given Key encoded with Base32.
func BinaryFromDsKey(dsKey key.Key) ([]byte, error) {
	return base32.RawStdEncoding.DecodeString(dsKey.String()[1:])
}

// MultihashToStrKey creates a Key from the given Multihash.
// If working with Cids, you can call cid.Hash() to obtain
// the multihash. Note that different CIDs might represent
// the same multihash.
func MultihashToStrKey(k mh.Multihash) key.StrKey {
	return NewStrKeyFromBinary(k)
}

// DsKeyToMultihash converts a Key to the corresponding Multihash.
func DsKeyToMultihash(dsKey key.Key) (mh.Multihash, error) {
	kb, err := BinaryFromDsKey(dsKey)
	if err != nil {
		return nil, err
	}
	return mh.Cast(kb)
}

// DsKeyToCidV1Raw converts the given Key (which should be a raw multihash
// key) to a Cid V1 of the given type (see
// https://pkg.go.dev/github.com/ipfs/go-cid#pkg-constants).
func DsKeyToCidV1(dsKey key.Key, codecType uint64) (cid.Cid, error) {
	hash, err := DsKeyToMultihash(dsKey)
	if err != nil {
		return cid.Cid{}, err
	}
	return cid.NewCidV1(codecType, hash), nil
}
