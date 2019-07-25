package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mnc.gt/log"
)

var blockChain BlockChain
var queue Queue
var net Net

func main() {

	blockChain = newBlockChain()
	net = newNet()

	router := mux.NewRouter()

	router.HandleFunc("/v1.0/messages", pushMessageHandle).Methods(http.MethodPost)
	router.HandleFunc("/v1.0/chain", getChainHandle).Methods(http.MethodGet)
	router.HandleFunc("/v1.0/mine", mineBlockHandle).Methods(http.MethodPost)
	router.HandleFunc("/v1.0/pending", getPendingHandle).Methods(http.MethodGet)

	router.HandleFunc("/v1.0/net/nodes", addNodeHandle).Methods(http.MethodPost)

	port := ":9090"

	log.Infof("Server run in port %v ...", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Error(err)
		return
	}

}
