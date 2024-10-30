package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

// App struct
type App struct {
	ctx      context.Context
	listener net.Listener
}

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	listener, err := net.Listen("tcp", "localhost:42069")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	log.Println("[INFO] [GO] Server listening on localhost:42069")
	a.listener = listener

	go a.handleServer()
}

func (a *App) shutdown(ctx context.Context) {
	err := a.listener.Close()
	if err != nil {
		log.Println("[ERROR] [Go] Error shutting down server: ", err)
	}
}

func (a *App) handleConnection(conn net.Conn) {
	defer conn.Close()

	encoder := json.NewEncoder(conn)

	for {
		msg := Message{
			Action:  "main/update",
			Message: fmt.Sprintf("Message at %v", time.Now()),
		}

		err := encoder.Encode(msg)
		if err != nil {
			log.Println("[ERROR] [GO] Error sending message: ", err)
			break
		}
		fmt.Println("[DEBUG] [GO] Message sent: ", msg)

		time.Sleep(5 * time.Second)
	}
}

func (a *App) handleServer() {
	log.Println("[INFO] [GO] Handling server, awaiting connection")

	for {
		conn, err := a.listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err)
			continue
		}

		log.Println("[INFO] [GO] Client connected.")
		go a.handleConnection(conn)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SendMessage() string {
	fmt.Println("[INFO] [GO] Sending message.")

	return "Message"
}
