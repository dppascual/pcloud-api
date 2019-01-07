package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// API stores the config needed to handle requests
type API struct {
	Router *mux.Router
}

// Init initializes an API struct
func (a *API) Init() {
	// Declare a router
	a.Router = mux.NewRouter()

	// Register URL paths and handlers
	a.InitRoutes()
}

// InitRoutes registers Request Handlers
func (a *API) InitRoutes() {
	a.Router.HandleFunc("/devices/{id:[0-9]+}/stats", a.getDeviceStats).Methods("GET")
}

// Run sets the HTTP server's router
func (a *API) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func executeCommand(pipeWriter *io.PipeWriter, args ...string) {
	var cmd *exec.Cmd

	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(args[0])
	}

	cmd.Stdout = pipeWriter

	cmd.Run()
	pipeWriter.Close()
}

func (a *API) getDeviceStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Device ID")
	}

	//Declare a pipe
	pipeReader, pipeWriter := io.Pipe()

	// Run a command
	go executeCommand(pipeWriter, "vmstat")

	// Read pipeReader line by line
	var lastLine string
	scanner := bufio.NewScanner(pipeReader)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		pipeReader.Close()
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	// Close the pipeReader
	pipeReader.Close()

	statsFields := strings.Fields(lastLine)
	if len(statsFields) != 17 {
		log.Printf("Error: Number of fields is different to 17\nContent:\n%s", statsFields)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	// Convert values to integer
	var stats [17]int
	for index, value := range statsFields {
		stats[index], _ = strconv.Atoi(value)
	}

	data := DeviceStats{
		Read: time.Now(),
		Process: Process{
			Running: stats[0],
			Waiting: stats[1],
		},
		Memory: Memory{
			Virtual: stats[2],
			Free:    stats[3],
			Buffer:  stats[4],
			Cache:   stats[5],
		},
		Swap: Swap{
			SwapIn:  stats[6],
			SwapOut: stats[7],
		},
		IO: IO{
			BlocksIn:  stats[8],
			BlocksOut: stats[9],
		},
		System: System{
			Interrupts:    stats[10],
			ContextSwitch: stats[11],
		},
		CPU: CPU{
			Idle:   stats[12],
			User:   stats[13],
			System: stats[14],
			Wait:   stats[15],
			Stolen: stats[16],
		},
	}

	respondWithJSON(w, http.StatusOK, data)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{
		"code":  strconv.Itoa(code),
		"error": message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
