package main

import (
	"encoding/json"
	"net/http"
	"time"

	"log"
)

func pushMessageHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Pushing message...")

	message := Message{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if message.Content == "" {
		log.Println("No push mesage...")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message.Timestap = time.Now().Format(time.RFC3339)

	queue.Messages = append(queue.Messages, message)

	log.Printf("Add new message '%v' to queue.", message.Content)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Server", "Success message")
	return
}

func getChainHandle(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(blockChain)
	if err != nil {
		log.Println("Error al devolver el blockchain")
		log.Println(err)
		return
	}

	log.Println("Return chain...")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func mineBlockHandle(w http.ResponseWriter, r *http.Request) {

	blockChain.consensus()

	index := blockChain.mine(queue.Messages)
	if index == 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	queue.clean() // cliean messages

	w.Header().Set("Server", "Success message")
	w.WriteHeader(http.StatusOK)
	return
}

func getPendingHandle(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(queue)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func addNodeHandle(w http.ResponseWriter, r *http.Request) {
	newNet := Net{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newNet)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newNet.Nodes == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	net.addNodes(newNet.Nodes)

	log.Println(net.Nodes)

	w.WriteHeader(http.StatusOK)
	return
}

func addNewBlock(w http.ResponseWriter, r *http.Request) {

	var messages []Message
	newBlock := Block{
		Transactions: messages,
	}

	err := json.NewDecoder(r.Body).Decode(&newBlock)
	if err != nil {
		log.Println(err)
		log.Println("No se ha podido Decodificar el cuerpo del mensaje")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	index := len(blockChain.Chain)

	newBlock.Index = int64(index)

	if !blockChain.validateBlock(newBlock) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blockChain.Chain = append(blockChain.Chain, newBlock)

	blockChain.consensus()

	w.WriteHeader(http.StatusOK)
	return
}
