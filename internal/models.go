package internal

import "encoding/json"

const (
	// commands
	WelcomeMessage     = "welcome"
	PingMessage        = "ping"
	PongMessage        = "pong"
	SubscribeMessage   = "subscribe"
	AckMessage         = "ack"
	UnsubscribeMessage = "unsubscribe"
	ErrorMessage       = "error"
	Message            = "message"
	Notice             = "notice"
	Command            = "command"

	// topics
	Symbol_ETHBTC = "ETH-BTC"
)

type WSMessage struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type WSIncomingMessage struct {
	*WSMessage
	Sn      string          `json:"sn"`
	Topic   string          `json:"topic"`
	Subject string          `json:"subject"`
	RawData json.RawMessage `json:"data"`
}

type WSSubscribeMessage struct {
	*WSMessage
	Topic          string `json:"topic"`
	PrivateChannel bool   `json:"privateChannel"`
	Response       bool   `json:"response"`
}
