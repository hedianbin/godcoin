package chain

import (
	"github.com/hedianbin/godcoin/util/hashx"
)

type TxPool interface{
	HaveTransaction(hash string) bool
	MaybeAcceptTransaction(tx *Transaction) ([]*hashx.Hash, error)
}
