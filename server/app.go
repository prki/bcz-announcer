package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// [TODO] I don't think we even need the existing callouts structure on backend.
// Will simplify things

// App struct
type App struct {
	ctx      context.Context
	listener net.Listener
	conn     net.Conn
	//existingCallouts []SetupMatch
	callouts CalloutStorage
	playerDb PlayerDb
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

type CalloutUpdateMessage struct {
	CalloutId     string `json:"callout_id"`
	UpdateRequest string `json:"update_request"` // change color, ... => ideally write down options
}

type AckMessage struct {
	AckAction  string `json:"ack_action"`  // Action which has been ACK'd
	AckMessage string `json:"ack_message"` // Additional message for ACK - e.g. card ID
}

type SetupMatch struct {
	GameName      string `json:"game_name"`
	P1Name        string `json:"p1_name"`
	P1CountryCode string `json:"p1_country_code"`
	P2Name        string `json:"p2_name"`
	P2CountryCode string `json:"p2_country_code"`
	CalloutId     string `json:"callout_id"`
}

type GGOBRow struct {
	PlayerName string `json:"player_name"`
	Points     int    `json:"points"`
}

type PlayerDb struct {
	DbConn *sql.DB
}

// Direct DB row
type PlayerRow struct {
	PlayerName  string
	CountryCode sql.NullString
}

// In-app representation. In case of unknown country code, the string is empty (”)
type Player struct {
	PlayerName  string `json:"player_name"`
	CountryCode string `json:"country_code"`
}

// Function fatally fails when the query fails. This is acceptable, since
// it is called only on the startup and in case of failure, the application
// would not be in working state.
func (db *PlayerDb) SelectPlayers() []Player {
	var ret []Player

	sql := "SELECT name, country_code FROM players ORDER BY name COLLATE NOCASE ASC;"
	rows, err := db.DbConn.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Fatal("[ERROR] Error reading players from db:", err)
	}

	for rows.Next() {
		row := PlayerRow{}
		err := rows.Scan(&row.PlayerName, &row.CountryCode)
		if err != nil {
			log.Fatal("[ERROR] Error scanning db:", err)
		}

		player := Player{PlayerName: row.PlayerName}
		if row.CountryCode.Valid {
			player.CountryCode = row.CountryCode.String
		} else {
			player.CountryCode = ""
		}

		ret = append(ret, player)
	}

	return ret
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

	dbConn, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	a.playerDb.DbConn = dbConn

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
			if err == io.EOF {
				log.Println("[ERROR] Got EOF on AckMessage listener, ending loop")
				return
			} // else if json error/timeout/.... For now, retry.
			continue
		}

		if ackMsg.AckAction == "callout/delete" {
			cardId := ackMsg.AckMessage
			log.Println("[GO] [ACKMSG] Received ack on callout delete. Deleting card", cardId)
			runtime.EventsEmit(a.ctx, "callout/delete", cardId)
		}
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
		a.conn = nil // re-set to nil in case conn is lost so that functions don't poop their pants
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
	runtime.EventsEmit(a.ctx, "callout/new", matchInfo)
}

// [TODO] We are repeatedly initializing a JSON encoder. Could be initialized
// on app level already.
func (a *App) sendCalloutCardDeleteRequest(cardId string) error {
	log.Println("[INFO] [GO] Sending callout card delete request, id:", cardId)
	if a.conn == nil {
		log.Println("[ERROR] [GO] No connection established, unable to send request")
		return errors.New("no connection to renderer")
	}

	encoder := json.NewEncoder(a.conn)

	msg := Message{
		Action:  "callout/delete",
		Message: cardId,
	}

	err := encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] [GO] Error encoding card delete request:", err)
		return errors.New("error encoding message")
	}

	log.Println("[INFO] [GO] Callout card id", cardId, "delete request sent.")
	return nil
}

// Exported function for Frontend to delete callout card. Note that it does not
// as a matter of fact delete the card - it only sends the request to the
// renderer.
// The deletion is done only as per ACK receive from renderer. This is to ensure
// that the app never removes a card from admin listing, but it stays on renderer.
// Better to have stale data on admin than on renderer :)
func (a *App) DeleteCalloutCard(cardId string) {
	//log.Println("[DEBUG] Cards before delete:", a.callouts.existingCallouts)
	log.Println("[INFO] Deleting callout card id:", cardId)
	// [TODO] Is it even valuable to have an internal storage? Probably not?
	//a.callouts.RemoveCalloutFromStorage(cardId)
	// [TODO] Check condition, emit event only on success
	// [TODO] Emit only when this is obtained from the renderer as ACK! Should be handled by a
	// different function.
	err := a.sendCalloutCardDeleteRequest(cardId)
	if err != nil {
		log.Println("[ERROR] Attempted to delete callout card id:", cardId, "but connection to renderer failed.")
		return
	}

	//runtime.EventsEmit(a.ctx, "callouts/delete", cardId)
	//log.Println("[DEBUG] Cards after delete:", a.callouts.existingCallouts)
}

// [TODO] Rename to SendNewCallout or something similar. Update should be used for
// status updates, not new cards.
func (a *App) SendUpdateCallout(p1Name, p1CountryCode, p2Name, p2CountryCode, gameName string) *SetupMatch {
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
		GameName:      gameName,
		P1Name:        p1Name,
		P1CountryCode: p1CountryCode,
		P2Name:        p2Name,
		P2CountryCode: p2CountryCode,
		CalloutId:     idStr,
	}

	matchInfoJson, err := json.Marshal(matchInfo)
	if err != nil {
		log.Println("[ERROR] Error building JSON for new callout, id:", idStr, "stopping")
		return nil
	}

	encoder := json.NewEncoder(a.conn)

	msg := Message{
		Action:  "callout/new",
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

func (a *App) SendCalloutUpdate(cardId, updateMsg string) {
	log.Println("[INFO] Sending callout update request, card id:", cardId, "msg:", updateMsg)
	if a.conn == nil {
		log.Println("[ERROR] Attempted to send callout update request for card:", cardId, "but conn was nil")
		return
	}

	msgObj := CalloutUpdateMessage{
		CalloutId:     cardId,
		UpdateRequest: updateMsg,
	}

	msgJson, err := json.Marshal(msgObj)
	if err != nil {
		log.Println("[ERROR] Error building JSON for callout update, id:", cardId, "stopping")
		return
	}

	// [TODO] Encoder should be a part of the state, no reason to initialize every time
	encoder := json.NewEncoder(a.conn)

	msg := Message{
		Action:  "callout/update",
		Message: string(msgJson),
	}

	err = encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] Error encoding update callout request: ", err)
		return
	}

	log.Println("[INFO] Sent callout update request, card id:", cardId)
}

func (a *App) SendFooterUpdate(content string) {
	log.Println("[INFO] Sending footer content update request. Content:", content)
	if a.conn == nil {
		// [TODO] Consider emiting some control event showing an alert/error in frontend
		log.Println("[ERROR] Attempted to send footer update request, conn was nil")
		return
	}

	msg := Message{
		Action:  "footer/update",
		Message: content,
	}

	// [TODO] Create encoder on top level, no reason to initialize every time
	encoder := json.NewEncoder(a.conn)

	err := encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] Error encoding footer update request:", err)
		return
	}

	log.Println("[INFO] Sent footer content update request.")
}

func (a *App) SendTextdisplayUpdate(content string) {
	msg := Message{
		Action:  "textdisplay/update",
		Message: content,
	}

	err := a.sendMessageToRenderer(msg)
	if err == nil {
		log.Println("[ERROR] Attempted to send text display update, but failed")
		// [TODO] Try to figure out toasts/notifs and emit event to frontend
		return
	}

	log.Println("[INFO] Sent text display update message to renderer:", msg)
}

// [TODO] Function is essentially the same as `SendFooterUpdate()`. Considering that,
// isn't there a way to just have a general handling function for sending messages?
// => there should be - these fns just prepare the message and have a generic `sendMessageToRenderer()` fn
func (a *App) SendHeaderUpdate(content string) {
	log.Println("[INFO] Sending header content update request. Content:", content)
	if a.conn == nil {
		log.Println("[ERROR] Attempted to send footer update request, conn was nil")
		return
	}

	msg := Message{
		Action:  "header/update",
		Message: content,
	}

	// [TODO] Create encoder on top level, no reason to initialize every time
	encoder := json.NewEncoder(a.conn)

	err := encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] Error encoding header update request:", err)
		return
	}

	log.Println("[INFO] Sent header update request.")

}

// Main function for sending prepared messages to the renderer.
// Function handles communication. Returns error in case of an error.
// Errors are not differentiated at this point since we don't have
// any error handling of relevance between different types of errors.
func (a *App) sendMessageToRenderer(msg Message) error {
	if a.conn == nil {
		log.Println("[ERROR] Attempted to send message:", msg, "conn was nil.")
		return errors.New("conn error")
	}

	// [TODO] Define/initialize encoder on struct level, not on method/fn call
	encoder := json.NewEncoder(a.conn)

	err := encoder.Encode(msg)
	if err != nil {
		log.Println("[ERROR] Error encoding message:", msg)
		return errors.New("encoding error")
	}

	return nil
}

func (a *App) readGGOBCsv() ([]GGOBRow, error) {
	log.Println("[INFO] Reading GGOB CSV file")

	file, err := os.Open("./ggob.csv")
	if err != nil {
		log.Println("[ERROR] Error opening ggob.csv:", err)
		return nil, err
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Println("[ERROR] Error reading ggob.csv:", err)
		return nil, err
	}

	var ret []GGOBRow
	for _, row := range data {
		pointsInt, err := strconv.Atoi(row[1])
		if err != nil {
			log.Println("[ERROR] Error parsing ggob.csv - points for player " + row[0] + " cannot be casted to int")
			return nil, err
		}

		newRow := GGOBRow{
			PlayerName: row[0],
			Points:     pointsInt,
		}

		ret = append(ret, newRow)
	}

	return ret, nil
}

func (a *App) ReadGGOBCsvPreview() *Message {
	log.Println("[INFO] Fetching GGOB CSV Preview")

	rows, err := a.readGGOBCsv()
	if err != nil {
		log.Println("[ERROR] Error fetching GGOB CSV preview:", err)
		msg := Message{
			Action:  "tableresults",
			Message: err.Error(),
		}

		return &msg
	}

	jsonRows, err := json.Marshal(rows)
	if err != nil {
		log.Println("[ERROR] Error creating json result for ggob.csv:", err)
		return nil
	}

	msg := Message{
		Action:  "tableresults",
		Message: string(jsonRows),
	}

	return &msg
}

func (a *App) SendGGOBCsv() {
	log.Println("[INFO] Sending GGOB CSV to renderer")

	rows, err := a.readGGOBCsv()
	if err != nil {
		log.Println("[ERROR] Error reading GGOB CSV:", err)
		return
	}

	rows = rows[:10] // only send top 10

	jsonRows, err := json.Marshal(rows)
	if err != nil {
		log.Println("[ERROR] Error creating json result for ggob.csv:", err)
		return
	}

	msg := Message{
		Action:  "tableresults/update",
		Message: string(jsonRows),
	}

	err = a.sendMessageToRenderer(msg)
	if err != nil {
		log.Println("[ERROR] Error sending message to renderer:", err)
		return
	}
}

func (a *App) GetPlayers() []Player {
	return a.playerDb.SelectPlayers()
}

// Function created so that wails models generate SetupMatch type
func (a *App) ReturnSetupMatch(matchInfo SetupMatch) SetupMatch {
	return matchInfo
}

// Function created so that wails models generate CalloutUpdateMessage type
func (a *App) ReturnCalloutUpdateMessage(cupdatemsg CalloutUpdateMessage) CalloutUpdateMessage {
	return cupdatemsg
}

// Function created so that wails models generate Message type
func (a *App) ReturnMessage(msg Message) Message {
	return msg
}

// Function created so that wails models generate GGOBRow type
func (a *App) ReturnGGOBRow(row GGOBRow) GGOBRow {
	return row
}
