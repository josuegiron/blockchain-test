package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

var blockChain BlockChain
var queue Queue
var net Net
var port string

func main() {

	blockChain = newBlockChain()
	net = newNet()

	router := mux.NewRouter()

	router.HandleFunc("/v1.0/messages", pushMessageHandle).Methods(http.MethodPost)
	router.HandleFunc("/v1.0/chain", getChainHandle).Methods(http.MethodGet)
	router.HandleFunc("/v1.0/chain/block", addNewBlock).Methods(http.MethodPost)
	router.HandleFunc("/v1.0/mine", mineBlockHandle).Methods(http.MethodPost)
	router.HandleFunc("/v1.0/pending", getPendingHandle).Methods(http.MethodGet)

	router.HandleFunc("/v1.0/net/nodes", addNodeHandle).Methods(http.MethodPost)

	port = ":9090"

	log.Println("Server run in port %v ...", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println(err)
		return
	}

}
