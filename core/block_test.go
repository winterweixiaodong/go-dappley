package core

import (
	"testing"

	"fmt"
	"time"

	"github.com/dappley/go-dappley/core/pb"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

var header = &BlockHeader{
	hash:      []byte{},
	prevHash:  []byte{},
	nonce:     0,
	timestamp: time.Now().Unix(),
}
var blk = &Block{
	header: header,
}

var expect = []byte{0x42, 0xff, 0x81, 0x3, 0x1, 0x1, 0xb, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1, 0xff, 0x82, 0x0, 0x1, 0x3, 0x1, 0x6, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1, 0xff, 0x84, 0x0, 0x1, 0xc, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1, 0xff, 0x90, 0x0, 0x1, 0x6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x1, 0x6, 0x0, 0x0, 0x0, 0x4d, 0xff, 0x83, 0x3, 0x1, 0x1, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1, 0xff, 0x84, 0x0, 0x1, 0x4, 0x1, 0x4, 0x48, 0x61, 0x73, 0x68, 0x1, 0xa, 0x0, 0x1, 0x8, 0x50, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x1, 0xa, 0x0, 0x1, 0x5, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x1, 0x4, 0x0, 0x1, 0x9, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1, 0x4, 0x0, 0x0, 0x0, 0x22, 0xff, 0x8f, 0x2, 0x1, 0x1, 0x13, 0x5b, 0x5d, 0x2a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1, 0xff, 0x90, 0x0, 0x1, 0xff, 0x86, 0x0, 0x0, 0x2e, 0xff, 0x85, 0x3, 0x1, 0x2, 0xff, 0x86, 0x0, 0x1, 0x4, 0x1, 0x2, 0x49, 0x44, 0x1, 0xa, 0x0, 0x1, 0x3, 0x56, 0x69, 0x6e, 0x1, 0xff, 0x8a, 0x0, 0x1, 0x4, 0x56, 0x6f, 0x75, 0x74, 0x1, 0xff, 0x8e, 0x0, 0x1, 0x3, 0x54, 0x69, 0x70, 0x1, 0x4, 0x0, 0x0, 0x0, 0x1d, 0xff, 0x89, 0x2, 0x1, 0x1, 0xe, 0x5b, 0x5d, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x58, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1, 0xff, 0x8a, 0x0, 0x1, 0xff, 0x88, 0x0, 0x0, 0x40, 0xff, 0x87, 0x3, 0x1, 0x1, 0x7, 0x54, 0x58, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1, 0xff, 0x88, 0x0, 0x1, 0x4, 0x1, 0x4, 0x54, 0x78, 0x69, 0x64, 0x1, 0xa, 0x0, 0x1, 0x4, 0x56, 0x6f, 0x75, 0x74, 0x1, 0x4, 0x0, 0x1, 0x9, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1, 0xa, 0x0, 0x1, 0x6, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x1, 0xa, 0x0, 0x0, 0x0, 0x1e, 0xff, 0x8d, 0x2, 0x1, 0x1, 0xf, 0x5b, 0x5d, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x1, 0xff, 0x8e, 0x0, 0x1, 0xff, 0x8c, 0x0, 0x0, 0x2f, 0xff, 0x8b, 0x3, 0x1, 0x1, 0x8, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x1, 0xff, 0x8c, 0x0, 0x1, 0x2, 0x1, 0x5, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1, 0x4, 0x0, 0x1, 0xa, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x1, 0xa, 0x0, 0x0, 0x0, 0x13, 0xff, 0x82, 0x1, 0x2, 0x1, 0x61, 0x2, 0xfc, 0xb6, 0xb2, 0x24, 0x6a, 0x0, 0x1, 0x1, 0x0, 0x1, 0x1, 0x0}

var header2 = &BlockHeader{
	hash:      []byte{'a'},
	prevHash:  []byte{'e', 'c'},
	nonce:     0,
	timestamp: time.Now().Unix(),
}
var blk2 = &Block{
	header: header2,
}

func TestHashTransactions(t *testing.T) {
	block := NewBlock([]*Transaction{&Transaction{}}, blk2)
	hash := block.HashTransactions()
	assert.NotNil(t, hash)
}

func TestNewBlock(t *testing.T) {
	block1 := NewBlock(nil, nil)
	assert.NotNil(t, block1.header.prevHash)
	assert.NotNil(t, block1.transactions)

	block2 := NewBlock(nil, blk)
	assert.NotNil(t, block2.header.prevHash)
	assert.Equal(t, Hash(Hash{}), block2.header.prevHash)
	assert.NotNil(t, block2.transactions)

	fmt.Println(blk2.GetHash())

	block3 := NewBlock(nil, blk2)
	fmt.Println(block3.header.prevHash)
	assert.NotNil(t, block3.header.prevHash)
	assert.Equal(t, Hash(Hash{'a'}), block3.header.prevHash)
	assert.Equal(t, []byte{'a'}[0], block3.header.prevHash[0])
	assert.Equal(t, uint64(1), block3.height)
	assert.NotNil(t, block3.transactions)

	block4 := NewBlock([]*Transaction{}, nil)
	assert.NotNil(t, block4.header.prevHash)
	assert.NotNil(t, block4.transactions)
	assert.Equal(t, []*Transaction{}, block4.transactions)
	assert.Equal(t, Hash(Hash{}), block4.header.prevHash)

	block5 := NewBlock([]*Transaction{&Transaction{}}, nil)
	assert.NotNil(t, block5.header.prevHash)
	assert.Equal(t, []*Transaction{&Transaction{}}, block5.transactions)
	assert.Equal(t, &Transaction{}, block5.transactions[0])
	assert.NotNil(t, block5.transactions)
}

func TestBlockHeader_Proto(t *testing.T) {
	bh1 := BlockHeader{
		[]byte("hash"),
		[]byte("hash"),
		1,
		2,
	}

	pb := bh1.ToProto()
	var i interface{} = pb
	_, correct := i.(proto.Message)
	assert.Equal(t, true, correct)
	mpb, err := proto.Marshal(pb)
	assert.Nil(t, err)

	newpb := &corepb.BlockHeader{}
	err = proto.Unmarshal(mpb, newpb)
	assert.Nil(t, err)

	bh2 := BlockHeader{}
	bh2.FromProto(newpb)

	assert.Equal(t, bh1, bh2)
}

func TestBlock_Proto(t *testing.T) {

	b1 := GenerateMockBlock()

	pb := b1.ToProto()
	var i interface{} = pb
	_, correct := i.(proto.Message)
	assert.Equal(t, true, correct)
	mpb, err := proto.Marshal(pb)
	assert.Nil(t, err)

	newpb := &corepb.Block{}
	err = proto.Unmarshal(mpb, newpb)
	assert.Nil(t, err)

	b2 := &Block{}
	b2.FromProto(newpb)

	assert.Equal(t, *b1, *b2)
}

func TestBlock_VerifyHash(t *testing.T) {
	b1 := GenerateMockBlock()
	fmt.Println(b1)

	//The mocked block does not have correct hash value
	assert.False(t, b1.VerifyHash())

	//calculate correct hash value
	hash := b1.CalculateHash()
	b1.SetHash(hash)

	//then this should be correct
	assert.True(t, b1.VerifyHash())
}


func TestBlock_Rollback(t *testing.T) {
	b := GenerateMockBlock()
	tx := MockTransaction()
	b.transactions = []*Transaction{tx}

	b.Rollback()
	txnPool := GetTxnPoolInstance()
	assert.ElementsMatch(t, tx.ID, txnPool.Pop().(Transaction).ID)
}