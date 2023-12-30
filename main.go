package main

import (
	"fmt"
	"time"

	"ws/internal"
)

func main() {
	res, err := internal.Login()
	if err != nil {
		panic(err)
	}
	cops := internal.ConnOpts{
		InsecureSkipVerify: false,
		ReadBufferSize:     2048000,
		PingInterval:       12 * time.Second,
		Timeout:            5 * time.Second,
		SubScriptionDelay:  100 * time.Millisecond,
	}

	var ws internal.WSConnection
	go func() {
		for {
			time.Sleep(5 * time.Second)
			// simulating an incident
			// ws.ForceStop()
		}
	}()
	i := 1
	for {
		fmt.Println(i)
		i++
		ws = internal.NewWsConnection(cops)
		err = ws.Connect(res)
		var syms *internal.Symbols
		if err != nil {
			fmt.Println("cannot connect to server : %v", err)
			goto end
		}
		syms, err = internal.GetSymbols()
		if err != nil {
			fmt.Println("cannot connect to server : %v", err)
			goto end
		}
		fmt.Println(internal.SymbolsTopics(syms)[:10])
		err = ws.SubscribeCandle(internal.Candle_1Min, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_3Min, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_5Min, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_15Min, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_1Hour, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_2Hour, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_4Hour, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_1Day, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}
		err = ws.SubscribeCandle(internal.Candle_1Week, internal.SymbolsTopics(syms)[:10]...)
		if err != nil {
			panic(err)
		}

		ws.Wait()
	end:
		fmt.Println("connection dropped retrying in 10 seconds...")
		time.Sleep(10 * time.Second)
	}
}
