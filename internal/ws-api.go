package internal

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// Usable for Sub and Unsub
func (ws *WSConnection) Subscribe(unsubscribe bool, topics ...string) error {
	sub := SubscribeMessage
	if unsubscribe {
		sub = UnsubscribeMessage
	}
	for _, v := range topics {
		c := WSSubscribeMessage{
			WSMessage: &WSMessage{
				Id:   ws.id,
				Type: sub,
			},
			Topic:          fmt.Sprintf("/market/ticker:%s", v),
			PrivateChannel: false,
			Response:       true,
		}
		m := ToJsonString(c)
		if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			return err
		}
		select {
		case id := <-ws.acks:
			if id != c.Id {
				return errors.Errorf("Invalid ack id %s, expect %s", id, c.Id)
			}
		case err := <-ws.errors:
			return errors.Errorf("Subscribe failed, %s", err.Error())
		case <-time.After(ws.ConnOpts.Timeout):
			return errors.Errorf("Wait ack message timeout in %v", ws.ConnOpts.Timeout)
		}
		time.Sleep(ws.ConnOpts.SubScriptionDelay)
	}
	return nil
}

type timeScale string

const (
	Candle_1Min  timeScale = "1min"
	Candle_3Min  timeScale = "3min"
	Candle_5Min  timeScale = "5min"
	Candle_15Min timeScale = "15min"
	Candle_1Hour timeScale = "1hour"
	Candle_2Hour timeScale = "2hour"
	Candle_4Hour timeScale = "4hour"
	Candle_1Day  timeScale = "1day"
	Candle_1Week timeScale = "1week"
)

func (ws *WSConnection) SubscribeCandle(ts timeScale, topics ...string) error {
	timeScale := Candle_1Hour
	if ts != "" {
		timeScale = ts
	}
	for _, v := range topics {
		c := WSSubscribeMessage{
			WSMessage: &WSMessage{
				Id:   ws.id,
				Type: SubscribeMessage,
			},
			Topic:          fmt.Sprintf("/market/candles:%s_%s", v, timeScale),
			PrivateChannel: false,
			Response:       true,
		}
		m := ToJsonString(c)
		if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			return err
		}
		select {
		case id := <-ws.acks:
			if id != c.Id {
				return errors.Errorf("Invalid ack id %s, expect %s", id, c.Id)
			}
		case err := <-ws.errors:
			return errors.Errorf("Subscribe failed, %s", err.Error())
		case <-time.After(ws.ConnOpts.Timeout):
			return errors.Errorf("Wait ack message timeout in %v", ws.ConnOpts.Timeout)
		}
		time.Sleep(ws.ConnOpts.SubScriptionDelay)
	}
	return nil
}
