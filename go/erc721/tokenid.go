package erc721

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// A TokenID is a uint256 Solidity tokenId.
type TokenID [32]byte

// Uint256 returns id as a uint256.Int.
func (id *TokenID) Uint256() *uint256.Int {
	// This is guaranteed to never overflow, so it's safe to ignore the flag.
	i, _ := uint256.FromBig(id.Big())
	return i
}

// Big returns id as a big.Int.
func (id *TokenID) Big() *big.Int {
	return new(big.Int).SetBytes(id[:])
}

// Cmp compares id and oID and returns:
//
//   -1 if o <  oID
//    0 if o == oID
//    1 if o >  oID
func (id *TokenID) Cmp(oID *TokenID) int {
	return id.Big().Cmp(oID.Big())
}

// String returns a decimal text representation of id.
func (id *TokenID) String() string {
	return id.Text(10)
}

// String returns a string representation of id in the given base.
func (id *TokenID) Text(base int) string {
	return id.Big().Text(base)
}

// TokenIDFromUint256 returns the 32-byte buffer underlying u, typed as a
// TokenID pointer.
func TokenIDFromUint256(u *uint256.Int) *TokenID {
	t := TokenID(u.Bytes32())
	return &t
}

// TokenIDFromBig returns TokenIDFromUint256(uint256.FromBig(b)), or an error if
// b overflows 256 bits.
func TokenIDFromBig(b *big.Int) (*TokenID, error) {
	u, overflow := uint256.FromBig(b)
	if overflow {
		return nil, fmt.Errorf("%T %v overflows uint256", b, b)
	}
	return TokenIDFromUint256(u), nil
}

// TokenIDFromHex returns TokenIDFromUint256(uint256.FromHex(b)), or any error
// returned by FromHex().
func TokenIDFromHex(h string) (*TokenID, error) {
	u, err := uint256.FromHex(h)
	if err != nil {
		return nil, err
	}
	return TokenIDFromUint256(u), nil
}

// TokenIDFromUint64 returns TokenIDFromUint256(uint256.NewInt(u)).
func TokenIDFromUint64(u uint64) *TokenID {
	return TokenIDFromUint256(uint256.NewInt(u))
}

// TokenIDFromInt returns TokenIDFromUint256(uint256.NewInt(uint64(i))).
func TokenIDFromInt(i int) *TokenID {
	return TokenIDFromUint256(uint256.NewInt(uint64(i)))
}

// Interface is the ERC721 standard interface, as generated by abigen.
type Interface interface {
	OwnerOf(_ *bind.CallOpts, tokenId *big.Int) (common.Address, error)
}
