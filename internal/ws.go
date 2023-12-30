package internal

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type WSConnection struct {
	id string
	wg *sync.WaitGroup
	// Stop subscribing channel
	done chan struct{}
	// Pong channel to check pong message
	pongs chan string
	// ACK channel to check pong message
	acks chan string
	// Error channel
	errors chan error
	// Downstream message channel
	messages chan *WSIncomingMessage
	conn     *websocket.Conn
	ConnOpts
	AuthInfo
}
type ConnOpts struct {
	InsecureSkipVerify bool
	ReadBufferSize     int
	EnableHeartbeat    bool
	SkipVerifyTls      bool
	Timeout            time.Duration
	PingInterval       time.Duration
	PingTimeout        time.Duration
	SubScriptionDelay  time.Duration
}

func NewWsConnection(c ConnOpts) WSConnection {
	return WSConnection{
		pongs:    make(chan string),
		messages: make(chan *WSIncomingMessage),
		done:     make(chan struct{}),
		acks:     make(chan string),
		wg:       &sync.WaitGroup{},
		ConnOpts: c,
	}
}

func (ws *WSConnection) Wait() {
	ws.wg.Wait()
}

func (ws *WSConnection) Connect(res *AuthInfo) error {
	s := res.Data.InstanceServers[0]
	ws.AuthInfo = *res

	q := url.Values{}
	ws.id = strconv.FormatInt(time.Now().UnixNano(), 10)
	q.Add("connectId", ws.id)
	q.Add("token", res.Data.Token)
	u := fmt.Sprintf("%s?%s", s.Endpoint, q.Encode())

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: ws.ConnOpts.InsecureSkipVerify}
	websocket.DefaultDialer.ReadBufferSize = ws.ConnOpts.ReadBufferSize
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return err
	}
	for {
		m := &WSIncomingMessage{}
		if err := conn.ReadJSON(m); err != nil {
			return err
		}
		if m.Type == ErrorMessage {
			return errors.Errorf("Error message: %v", m)
		}
		if m.Type == WelcomeMessage {
			fmt.Println("connected")
			break
		}
	}
	ws.conn = conn

	ws.wg.Add(2)
	go ws.startReader()
	go ws.startHeartbeat()

	return nil
}

func (ws *WSConnection) ForceStop() {
	close(ws.done)
}

func (ws *WSConnection) startReader() {
	defer func() {
		close(ws.pongs)
		close(ws.messages)
		ws.wg.Done()
	}()

	for {
		select {
		case <-ws.done:
			return
		default:
			m := &WSIncomingMessage{}
			if err := ws.conn.ReadJSON(m); err != nil {
				ws.errors <- err
				return
			}
			fmt.Println(m.Topic)
			switch m.Type {
			case WelcomeMessage:
			case PongMessage:
				if ws.ConnOpts.EnableHeartbeat {
					ws.pongs <- m.Id
				}
			case AckMessage:
				ws.acks <- m.Id
			case ErrorMessage:
				ws.errors <- errors.Errorf("Error message: %v", *m)
				return
			case Message, Notice, Command:
				fmt.Println(ToJsonString(m))
			default:
				ws.errors <- errors.Errorf("Unknown message type: %s", m.Type)
			}
			// time.Sleep(500 * time.Millisecond)
		}
	}
}

func (ws *WSConnection) startHeartbeat() {
	pt := time.NewTicker(ws.ConnOpts.PingInterval - time.Millisecond*200)
	defer func() {
		ws.wg.Done()
		pt.Stop()
	}()

	for {
		select {
		case <-ws.done:
			return
		case <-pt.C:
			p := NewPingMessage()
			m := ToJsonString(p)
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
				ws.errors <- err
				return
			}

			select {
			case pid := <-ws.pongs:
				if pid != p.Id {
					ws.errors <- errors.Errorf("Invalid pong id %s, expect %s", pid, p.Id)
					return
				}
			case <-time.After(ws.ConnOpts.PingTimeout):
				ws.errors <- errors.Errorf("pong message timeout in %d ms", ws.ConnOpts.PingTimeout.Milliseconds())
				return
			}
		}
	}
}

// NewPingMessage creates a ping message instance.
func NewPingMessage() *WSMessage {
	return &WSMessage{
		Id:   strconv.FormatInt(time.Now().UnixNano(), 10),
		Type: PingMessage,
	}
}

// ToJsonString converts any value to JSON string.
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
