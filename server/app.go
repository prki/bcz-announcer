package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	listener net.Listener
	conn     net.Conn
	//existingCallouts []SetupMatch
	callouts CalloutStorage
}

type CalloutStorage struct {
	mu               sync.Mutex
	existingCallouts []SetupMatch
}

func (cs *CalloutStorage) AddToStorage(match SetupMatch) {
	cs.mu.Lock()
	cs.existingCallouts = append(cs.existingCallouts, match)
	cs.mu.Unlock()
}

func (cs *CalloutStorage) RemoveCalloutFromStorage(id string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	for idx, match := range cs.existingCallouts {
		if match.CalloutId == id {
			cs.existingCallouts = append(cs.existingCallouts[:idx], cs.existingCallouts[idx+1:]...)
			return true
		}
	}

	log.Println("[WARN] Attempted to remove card id:", id, "but card not found in storage!")
	return false
}

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type AckMessage struct {
	AckAction  string `json:"ack_action"`  // Action which has been ACK'd
	AckMessage string `json:"ack_message"` // Additional message for ACK - e.g. card ID
}

type SetupMatch struct {
	GameName  string `json:"game_name"`
	P1Name    string `json:"p1_name"`
	P2Name    string `json:"p2_name"`
	CalloutId string `json:"callout_id"`
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

func (a *App) handleAckMessages() {
	decoder := json.NewDecoder(bufio.NewReader(a.conn))

	for {
		var ackMsg AckMessage
		err := decoder.Decode(&ackMsg)
		if err != nil {
			log.Println("[GO] [ACKMSG] Error decoding message:", err)
			continue
		}

		log.Println("[GO] [ACKMSG] Received ACK msg:", ackMsg)
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
		a.handleAckMessages()
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

// [TODO] Lock behind mutex in case someone clicks very fast :)
func (a *App) addCalloutToStorage(matchInfo SetupMatch) {
	//a.existingCallouts = append(a.existingCallouts, matchInfo)
	a.callouts.AddToStorage(matchInfo)
	runtime.EventsEmit(a.ctx, "callouts/new", matchInfo)
}

func (a *App) DeleteCalloutCard(cardId string) {
	log.Println("[DEBUG] Cards before delete:", a.callouts.existingCallouts)
	a.callouts.RemoveCalloutFromStorage(cardId)
	// [TODO] Check condition, emit event only on success
	runtime.EventsEmit(a.ctx, "callouts/delete", cardId)
	log.Println("[DEBUG] Cards after delete:", a.callouts.existingCallouts)

}

func (a *App) SendUpdateCallout(p1Name, p2Name, gameName string) *SetupMatch {
	log.Println("[INFO] [GO] Sending update callout request: ", p1Name, p2Name)
	if a.conn == nil {
		log.Println("[ERROR] [GO] No connection established, unable to send request")
		return nil
	}

	log.Println("[DEBUG] [GO] Generating an ID for callout.")
	id, err := uuid.NewV7()
	if err != nil {
		log.Println("[ERROR] [GO] Error generating UUID for callout:", err)
		return nil
	}
	idStr := id.String()
	log.Println("[DEBUG] [GO] Generated UUID:", idStr)

	matchInfo := &SetupMatch{
		GameName:  gameName,
		P1Name:    p1Name,
		P2Name:    p2Name,
		CalloutId: idStr,
	}

	matchInfoJson, err := json.Marshal(matchInfo)
	encoder := json.NewEncoder(a.conn)

	if err != nil {
		log.Println("[ERROR] [GO] Error creating JSON message about state info: ", err)
		return nil
	}

	msg := Message{
		Action:  "callout/update",
		Message: string(matchInfoJson),
	}

	err = encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] [GO] Error encoding update callout request: ", err)
		return nil
	}

	// [TODO] This should be appended to storage only when renderer returns
	// information that it got through and it's there.
	// Leaving it as is for now for debug purposes.
	log.Println("[DEBUG] [GO] Adding callout to storage.")
	a.addCalloutToStorage(*matchInfo)
	//a.existingCallouts = append(a.existingCallouts, *matchInfo)
	// [TODO] Log should also have a mutex lock :-----)
	log.Println("[DEBUG] [GO] Callouts in storage:", a.callouts.existingCallouts)

	return matchInfo
}

// Function created so that wails models generate SetupMatch type
func (a *App) returnSetupMatch(matchInfo SetupMatch) SetupMatch {
	return matchInfo
}
