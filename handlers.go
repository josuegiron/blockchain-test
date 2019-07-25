package main

import (
	"encoding/json"
	"net/http"
	"time"

	"go.mnc.gt/log"
)

func pushMessageHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Pushing message...")

	message := Message{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if message.Content == "" {
		log.Error("No push mesage...")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message.Timestap = time.Now().Format(time.RFC3339)

	queue.Messages = append(queue.Messages, message)

	log.Infof("Add new message '%v' to queue.", message.Content)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Server", "Success message")
	return
}

func getChainHandle(w http.ResponseWriter, r *http.Request) {
	log.Info("Return chain...")
	response, err := json.Marshal(blockChain)
	if err != nil {
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func mineBlockHandle(w http.ResponseWriter, r *http.Request) {

	index := blockChain.mine(queue)
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
		log.Error(err)
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
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newNet.Nodes == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	net.addNodes(newNet.Nodes)

	log.Info(net.Nodes)

	w.WriteHeader(http.StatusOK)
	return
}
