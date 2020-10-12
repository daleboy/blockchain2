package blockchain2

//Blockchain 区块链结构，用数组（array）存储有序的区块
type Blockchain struct {
	Blocks []*Block
}

//AddBlock 挖出普通区块并将新区块加入到区块链中
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash) //挖出区块
	bc.Blocks = append(bc.Blocks, newBlock)
}

//NewBlockchain 创建初始区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
