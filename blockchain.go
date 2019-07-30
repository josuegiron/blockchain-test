package main 

import (
	"strings"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
)

// BlockChain doc ...
type BlockChain struct {
	Chain []Block `json:"blocks"`
}

const (
	difficulty = 2
) 

func newBlockChain() BlockChain {
	newBlockChain := BlockChain{}
	newBlockChain.createGenesisBlock()
	return newBlockChain
}

func (blockChain *BlockChain)createGenesisBlock(){
	
	genessisBlock := newBlock(0, "genesis", nil, 0)

	blockChain.Chain = append(blockChain.Chain, genessisBlock) 
	
}

func (blockChain *BlockChain)getDifficulty() (string) {
	str := ""
	for i := 0; i < difficulty; i++ {
		str += "0"
	}

	return str
}

func (blockChain *BlockChain)proofOfWorck(block *Block) string {
	block.Nonce = 0
	computedHash := block.generateHash()
	for !strings.HasPrefix(computedHash, blockChain.getDifficulty()) {
		block.Nonce++
		computedHash = block.generateHash()
	}

	return computedHash
}

func (blockChain *BlockChain)addBlock(block *Block) bool {
	previus :=  blockChain.Chain[block.Index-1].generateHash()

	if block.Previus != previus {
		return false
	}
	
	log.Println(blockChain.proofOfWorck(block))

	return true
}

func (blockChain *BlockChain)mine(transactions []Message) int64 {
	
	index:= int64(len(blockChain.Chain))
	previus := blockChain.Chain[index - 1].generateHash()
	newBlock := Block{
		Index: index,
		Transactions: transactions,
		Previus: previus,
		Nonce: 0,
	}

	if !blockChain.addBlock(&newBlock){
		return 0 
	}

	if !blockChain.announceNewBlock(newBlock) {
		return 0
	}
	
	blockChain.Chain = append(blockChain.Chain, newBlock)

	return newBlock.Index
}

func (blockChain *BlockChain)consensus(){
	 longestChain := blockChain.Chain
	 
	 for _, node := range net.Nodes{
		 
		currentLen := len(longestChain)

		 url := fmt.Sprintf("http://%v%v/v1.0/chain", node.URL, port)
		 response, err := http.Get(url)
		if err != nil {
			log.Println("Error al realizar la petición /chanin[GET]")
			log.Println(err)
			continue
		}
		var newChain *BlockChain
		decode := json.NewDecoder(response.Body)
		err = decode.Decode(&newChain)
		if err != nil {
			log.Println("Error al obtener la información del Body en /chain[GET]")
			log.Println(err)
			continue
		}

		log.Println(newChain)
		
		length := len(newChain.Chain)
		 if validateNewChain(*newChain) && currentLen < length{
			 log.Println("Replace chain")
			 longestChain = newChain.Chain
			 log.Println(longestChain)
		 }
	}

	log.Println(longestChain)
	blockChain.Chain = longestChain
	log.Println(blockChain)

}

func validateNewChain(blockChain BlockChain) bool {
	for i, block := range blockChain.Chain{
		if i == 0 {
			continue
		}
		if block.Previus != blockChain.Chain[i-1].generateHash(){
			log.Println("Cadena invádlida, cadena viene del consenso")
			return false
		}
	}
	return true
}

func (blockChain *BlockChain)validateBlock(block Block) bool{
	log.Println(block)
	if block.Previus != blockChain.Chain[block.Index-1].generateHash(){
		log.Println("El bloque anterior no coincide con el bloque anterior")
		return false
	}

	if !strings.HasPrefix(block.generateHash(), blockChain.getDifficulty()){
		log.Println("El bloque no genera un hash válido...")
		jsonBlock, _ := json.Marshal(block)
		log.Println(string(jsonBlock))
		return false
	}
	index := len(blockChain.Chain)
	messages := queue.Messages
	newTestBlock := Block{
		Transactions: messages,
	}
		
	newTestBlock.Index = int64(index)
	newTestBlock.Transactions = messages
	newTestBlock.Previus = blockChain.Chain[index-1].generateHash()
	newTestBlock.Nonce = block.Nonce

	if newTestBlock.generateHash() != block.generateHash(){
		log.Println("No coninciden los bloques")
		log.Println(newTestBlock)
		log.Println(block)
		return false
	}
	return true 
}


func (blockChain *BlockChain)announceNewBlock(block Block) bool {
	requestBody, err := json.Marshal(block)	
	if err != nil {
		log.Println(err)
		return false
	}
	for _, node := range net.Nodes {

		url := fmt.Sprintf("http://%v%v/v1.0/chain/block", node.URL, port)

		log.Println(url)

		resp, err := http.Post(url,"application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		stCode := resp.StatusCode

		if stCode != 200 {
			log.Println("No se ha podido publicar")
			log.Println(stCode)
			return false
		}
	}	
	return true
}