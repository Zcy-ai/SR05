package handler

import (
	"bufio"
	"fmt"
	"os"
)

// To handle the send of messages and the reception of messages
type Handler struct {
	messageChan chan Message
}

type Message struct {
	key   string // type of message: {"send", "print"}
	value string // Message
}

func NewHandler() Handler {
	return Handler{
		messageChan: make(chan Message),
	}
}

// to handle the message of different types
func (h *Handler) handle() {
	for act := range h.messageChan {
		switch act.key {
		case "send":
			fmt.Fprintln(os.Stdout, act.value)
		case "print":
			fmt.Fprintln(os.Stderr, act.value)
			os.Stdout.Sync() // flush la sortie standard pour être sûr que le message est bien affiché
		default:
			fmt.Fprintln(os.Stderr, "Unknown Message:", act)
		}
	}
}

// receive the message from stdin and inform the handler to print the message
func (h *Handler) ReadMessage() {
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		h.messageChan <- Message{
			key:   "print",
			value: string(reader.Bytes()),
		}
	}
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error while scanning stdin:", err)
		os.Exit(1)
	}
}

// inform the handler to send a message
func (h *Handler) SendMessage(message string) {
	h.messageChan <- Message{
		key:   "send",
		value: message,
	}
}

func (h *Handler) Run() {
	go h.ReadMessage()
	go h.handle()
}

func (h *Handler) Close() {
	close(h.messageChan)
}
