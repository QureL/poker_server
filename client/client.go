package client

const (
	CHANNEL_SIZE = 1
)

type Client struct {
	ChannelIn  chan string
	ChannelOut chan string
}

func NewClient() *Client {
	return &Client{
		ChannelIn:  make(chan string, CHANNEL_SIZE),
		ChannelOut: make(chan string, CHANNEL_SIZE),
	}
}
