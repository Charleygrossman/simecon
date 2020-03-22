package db

import "math/big"

// maxint64 is a pointer to the largest int64 value.
var maxint64 = big.NewInt(int64(^uint64(0) >> 1))