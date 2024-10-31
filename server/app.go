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
	conn     net.Conn
}

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type SetupMatch struct {
	P1Name string `json:"p1_name"`
	P2Name string `json:"p2_name"`
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
	a.conn = nil

	go a.handleServer()
}

func (a *App) shutdown(ctx context.Context) {
	err := a.conn.Close()
	if err != nil {
		log.Println("[ERROR] [Go] Error shutting down conn: ", err)
	}

	err = a.listener.Close()
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
		log.Println("[INFO] [GO] Listening for connections")
		conn, err := a.listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err)
			continue
		}

		log.Println("[INFO] [GO] Client connected.")
		a.conn = conn
		//go a.handleConnection(conn)
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

func (a *App) SendChangeViewRequest(viewName string) {
	log.Println("[INFO] [GO] Sending change view request to view name: ", viewName)
	if a.conn == nil {
		log.Println("[ERROR] [GO] No connection established, unable to send request")
		return
	}

	encoder := json.NewEncoder(a.conn)

	msg := Message{
		Action:  "control/change_view",
		Message: viewName, // [TODO] Send as JSON object since that will be relevant in the future
	}

	err := encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] [GO] Error encoding change view request: ", err)
		return
	}

	log.Println("[INFO] [GO] Change view request to view " + viewName + " sent.")
}

func (a *App) SendUpdateCalloutRequest(matchInfo SetupMatch) {
	log.Println("[INFO] [GO] Sending update callout request: ", matchInfo)
	if a.conn == nil {
		log.Println("[ERROR] [GO] No connection established, unable to send request")
		return
	}

	matchInfoJson, err := json.Marshal(matchInfo)
	encoder := json.NewEncoder(a.conn)

	if err != nil {
		log.Println("[ERROR] [GO] Error creating JSON message about state info: ", err)
		return
	}

	msg := Message{
		Action:  "callout/update",
		Message: string(matchInfoJson),
	}

	err = encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] [GO] Error encoding update callout request: ", err)
		return
	}
}

// Function created so that wails models generate SetupMatch type
func (a *App) returnSetupMatch(matchInfo SetupMatch) SetupMatch {
	return matchInfo
}
