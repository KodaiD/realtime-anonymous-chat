package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client represents a user
type Client struct {
	ID   string
	Conn *websocket.Conn
	Room *Room
}

// Message represents a message
type Message struct {
	Type   int    `json:"type"`
	Body   string `json:"body"`
	Author string `json:"user"`
}

// Read messages sent from websocket
func (c *Client) Read() {
	defer func() {
		c.Room.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p), Author: c.ID}
		c.Room.Broadcast <- message
		fmt.Println("Message Received:", message.Body)
	}
}
