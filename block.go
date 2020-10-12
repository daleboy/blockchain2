package blockchain2

import "time"

//Block 区块结构新版，增加了计数器nonce，主要目的是为了校验区块是否合法
//即挖出的区块是否满足工作量证明要求的条件
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Nonce         int
	Hash          []byte
}

//NewBlock 创建普通区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, 0, []byte{}}

	//挖矿实质上是算出符合要求的哈希
	pow := NewProofOfWork(block) //注意传递block指针作为参数
	nonce, hash := pow.Run()

	//设置block的计数器和哈希
	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}

//NewGenesisBlock 创建创始区块。创建创始区块也需要挖矿。
func NewGenesisBlock() *Block {
	return NewBlock("创始区块", []byte{})
}
