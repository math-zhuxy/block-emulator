package core

import (
	"blockEmulator/params"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// var (
// 	EmptyRootHash = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
// )

// Header represents a block header in the Ethereum blockchain.
type BlockHeader struct {
	ParentHash []byte `json:"parentHash"       gencodec:"required"`
	StateRoot  []byte `json:"stateRoot"        gencodec:"required"`
	TxHash     []byte `json:"transactionsRoot" gencodec:"required"`
	MigHash    []byte `json:"migrationRoot" gencodec:"required"`
	Number     int    `json:"number"           gencodec:"required"`
	Time       uint64 `json:"timestamp"        gencodec:"required"`
}

func (bh *BlockHeader) Encode() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(bh)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func (bh *BlockHeader) Hash() []byte {
	hash := sha256.Sum256(bh.Encode())
	return hash[:]
}

func DecodeBlockHeader(to_decode []byte) *BlockHeader {
	var blockHeader BlockHeader

	decoder := gob.NewDecoder(bytes.NewReader(to_decode))
	err := decoder.Decode(&blockHeader)
	if err != nil {
		log.Panic(err)
	}

	return &blockHeader
}

func (bh *BlockHeader) PrintBlockHeader() {
	vals := []interface{}{
		hex.EncodeToString(bh.ParentHash),
		hex.EncodeToString(bh.StateRoot),
		hex.EncodeToString(bh.TxHash),
		bh.Number,
		bh.Time,
	}
	fmt.Printf("%v\n", vals)
}

// 区块结构
type Block struct {
	Header       *BlockHeader
	Transactions []*Transaction
	TXmig1s      []*TXmig1
	TXmig2s      []*TXmig2
	Anns         []*TXann
	NSs          []*TXns
	Hash         []byte
	Fee          float64
}

// core/types/block.go
func NewBlock(blockHeader *BlockHeader, txs []*Transaction, txmig1s []*TXmig1, txmig2s []*TXmig2, anns []*TXann, nss []*TXns) *Block {
	b := &Block{
		Header:       blockHeader,
		Transactions: txs,
		TXmig1s:      txmig1s,
		TXmig2s:      txmig2s,
		Anns:         anns,
		NSs:          nss,
	}

	return b
}

func (b *Block) PrintBlock() {
	fmt.Printf("blockHeader: \n")
	b.Header.PrintBlockHeader()
	// fmt.Printf("transactions: \n")
	// for _, tx := range b.Transactions {
	// 	tx.PrintTx()
	// }
	fmt.Printf("# of transactions: %v\n", len(b.Transactions))
	if !params.Config.Stop_When_Migrating {
		fmt.Printf("# of TXmig1s: %v\n", len(b.TXmig1s))
		fmt.Printf("# of In1s: %v\n", len(b.TXmig2s))
		fmt.Printf("# of Anns: %v\n", len(b.Anns))
		fmt.Printf("# of NSs: %v\n", len(b.NSs))
	}
	fmt.Printf("blockHash: \n")
	fmt.Printf("%v\n", hex.EncodeToString(b.Hash))
}

// special
func (b *Block) GetHash() []byte {
	return b.Header.Hash()
}

func (b *Block) Encode() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func DecodeBlock(to_decode []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(to_decode))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
