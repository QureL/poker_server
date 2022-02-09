package statemachine

import (
	"poker_server/client"
	"sync"
)

var clients sync.Map

func getClients(room int) [2]*client.Client {
	obj, _ := clients.Load(room)
	pair, _ := obj.([2]*client.Client)
	return pair
}

func rmClients(room int) {
	clients.Delete(room)
}

var startChannel sync.Map

func storeStartChannel(room int) chan struct{} {
	channel := make(chan struct{})
	startChannel.Store(room, channel)
	return channel
}

func getStartChannel(room int) chan struct{} {
	obj, ok := startChannel.Load(room)
	if !ok {
		return nil
	}
	channel, _ := obj.(chan struct{})
	return channel
}

func rmStartChannel(room int) {
	startChannel.Delete(room)
}

var stopChannel sync.Map

func storeStopChannel(room int) chan struct{} {
	channel := make(chan struct{})
	stopChannel.Store(room, channel)
	return channel
}

func getStopChannel(room int) chan struct{} {
	obj, ok := stopChannel.Load(room)
	if !ok {
		return nil
	}
	channel, _ := obj.(chan struct{})
	return channel
}

func rmStopChannel(room int) {
	startChannel.Delete(room)
}
