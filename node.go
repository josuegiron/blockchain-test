package main

import (
	osnet "net"
	"os"

	"go.mnc.gt/log"
)

// Node doc ...
type Node struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// Net ...
type Net struct {
	Nodes []Node `json:"nodes"`
}

func (net *Net) addMe() error {
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	log.Error(err)
	// 	return err
	// }

	addrs, err := osnet.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	me := Node{}

	for _, a := range addrs {
		if ipnet, ok := a.(*osnet.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				me = Node{
					URL:  ipnet.IP.String(),
					Name: "me",
				}
			}
		}
	}

	net.Nodes = append(net.Nodes, me)

	return nil
}

func newNet() Net {
	newNet := Net{}
	err := newNet.addMe()
	if err != nil {
		log.Error(err)
		return Net{}
	}
	return newNet
}

func (net *Net) addNodes(nodes []Node) {
	for _, node := range nodes {
		net.Nodes = append(net.Nodes, node)
	}
}
