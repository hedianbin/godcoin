package chain

import (
	"math/big"
	"github.com/hedianbin/godcoin/util/hashx"
	"bytes"
)

const(
	MaxBlockTransactions = 100000
	MaxBlockSerializedSize = 2000000
)

// ValidateBlock validate block data
func (bc *Blockchain)ValidateBlock(block *Block, powLimit *big.Int) error {
	//TODO: check ProofOfWork
	//TODO: check block time

	//check transaction's count
	//must have at least one
	numTx := len(block.Transactions)
	if numTx == 0 {
		return ErrBlockNoTransactions
	}

	// check max block payload is bigger than limit.
	if numTx > MaxBlockTransactions {
		return ErrBlockTooManyTransactions
	}

	//check max block's serialized size
	serializedSize := len(SerializeBlock(block))
	if serializedSize > MaxBlockSerializedSize{
		return ErrBlockSizeTooBig
	}

	// The first transaction in a block must be a coinbase.
	transactions := block.Transactions
	if !transactions[0].IsCoinBase() {
		return ErrFirstTxNotCoinbase
	}

	// check coinbase transaction count
	// count == 1
	for _, tx := range transactions[1:] {
		if tx.IsCoinBase() {
			return ErrMultipleCoinbases
		}
	}

	// validate each transaction
	for _, tx := range transactions {
		if !bc.VerifyTransaction(tx){
			return ErrNotVerifyTransaction
		}
	}

	// check merkleRoot
	merkleRoot := block.HashTransactions()
	if bytes.Compare(block.MerkleRoot, merkleRoot) != 0{
		return ErrBlockBadMerkleRoot
	}

	// Check for duplicate transactions.
	existingTxHashes := make(map[hashx.Hash]struct{})
	for _, tx := range transactions {
		hash := *tx.GetHash()
		if _, exists := existingTxHashes[hash]; exists {
			return ErrBlockDuplicateTx
		}
		existingTxHashes[hash] = struct{}{}
	}

	return nil
}

