package main

import "Go-Minichain/network"

func main() {
	network := network.NewNetWork()
	network.Start()
}
