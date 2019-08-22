package main

type room struct {

	// forward is channel that will receive message from client and send another client.

	forward chan []byte
}
