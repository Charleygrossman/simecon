package db

import "tradesim/common"

type Transaction interface {
	getTxnType() common.TxnType
	getHash() string
}
