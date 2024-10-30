package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
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

func (a *App) handleConnection() {
	decoder := json.NewDecoder(bufio.NewReader(a.conn))

	for {
		var msg Message
		err := decoder.Decode(&msg)
		if err != nil {
			fmt.Println("[ERROR] [GO] Error receiving message: ", err)
			continue
		}

		fmt.Println("[INFO] [GO] Received message: ", msg)
		runtime.EventsEmit(a.ctx, "message/display", msg)
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
