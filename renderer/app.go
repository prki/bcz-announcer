package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type AckMessage struct {
	AckAction  string `json:"ack_action"`  // Action which has been ACK'd
	AckMessage string `json:"ack_message"` // Additional message for ACK - e.g. card ID
}

// App struct
type App struct {
	ctx  context.Context
	conn net.Conn
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	conn, err := net.Dial("tcp", "localhost:42069")
	if err != nil {
		log.Fatal("Error connecting to server: ", err)
	}

	a.conn = conn

	go a.handleConnection()

	log.Println("[GO] [INFO] Connected to server.")
}

func (a *App) shutdown(ctx context.Context) {
	err := a.conn.Close()
	if err != nil {
		fmt.Println("[ERROR] [GO] Error closing conn: ", err)
	}
}

func (a *App) sendCalloutDeleteAck(encoder *json.Encoder, cardId string) error {
	ackMsg := AckMessage{
		AckAction:  "callout/delete",
		AckMessage: cardId,
	}

	log.Println("[INFO] Sending callout delete ACK for cardId:", cardId)
	err := encoder.Encode(ackMsg)
	if err != nil {
		log.Println("[ERROR] Attempted to send callout delete ack, but received err:", err)
		return errors.New("ack connection failed")
	}

	return nil
}

func (a *App) handleConnection() {
	decoder := json.NewDecoder(bufio.NewReader(a.conn))
	encoder := json.NewEncoder(a.conn)

	for {
		var msg Message
		err := decoder.Decode(&msg)
		if err != nil {
			fmt.Println("[ERROR] [GO] Error receiving message: ", err)
			continue
		}

		fmt.Println("[INFO] [GO] Received message: ", msg)
		if msg.Action == "control/change_view" {
			runtime.EventsEmit(a.ctx, "control/change_view", msg)
		} else if msg.Action == "callout/new" {
			runtime.EventsEmit(a.ctx, "callout/new", msg)
			// [TODO] Delete/write properly in case we use ACK for new callout in any way.
			/*ackMsg := AckMessage{
				AckAction:  "callout/update",
				AckMessage: "Test message",
			}

			err := encoder.Encode(ackMsg)
			if err != nil {
				log.Println("[ERROR] [GO] Error encoding ack msg:", err)
			}
			*/
		} else if msg.Action == "callout/update" {
			runtime.EventsEmit(a.ctx, "callout/update", msg)
		} else if msg.Action == "callout/delete" {
			cardId := msg.Message // Message: UUID of card as string
			err = a.sendCalloutDeleteAck(encoder, cardId)
			if err != nil {
				log.Println("[ERROR] [GO] Attempted to send callout delete ACK, but failed. Not removing.")
				continue
			}

			runtime.EventsEmit(a.ctx, "callout/delete", msg)
		} else if msg.Action == "footer/update" {
			runtime.EventsEmit(a.ctx, "footer/update", msg)
		} else if msg.Action == "header/update" {
			runtime.EventsEmit(a.ctx, "header/update", msg)
		} else if msg.Action == "textdisplay/update" {
			runtime.EventsEmit(a.ctx, "textdisplay/update", msg)
		} else {
			log.Println("[ERROR] Unhandled action:", msg.Action)
		}
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Function declaration existing to generate a model for Wails for trivial
// TS usage.
func (a *App) ReturnMessage(msg Message) Message {
	return msg
}
