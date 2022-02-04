package client

import "poker_server/cards"

type Client struct {
	Channel chan []cards.Card
}

func NewClient() *Client {
	return &Client{
		Channel: make(chan []cards.Card),
	}
}
