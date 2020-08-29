package main

import (
	"abhinav/ticket-service/db"
	"abhinav/ticket-service/models"
	"abhinav/ticket-service/utils"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

const (
	APP = "2c"
)

var (
	Config map[string]interface{}
)

func initMain() {
	// Loaling Config Json
	loadConfig()
	os.Setenv("APP", APP)
	if !reflect.ValueOf(Config["env"]).IsValid() {
		os.Setenv("ENV", "DEV")
	} else {
		os.Setenv("ENV", Config["env"].(string))
	}
	// Loading DB connections
	err := db.Init() // bootstrap.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = loadJobs()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func loadJobs() error {
	_ctx := context.Background()
	ctx := context.WithValue(_ctx, 1,
		map[string]interface{}{
			"requestID": "job_queue_load",
			"sessionID": "session_load",
		})
	tickets, err := models.GetPendingTickets(&ctx)
	if err != nil {
		return err
	}
	agents, _ := models.GetCurrentActiveAgents(context.Background())
	jobQueue := utils.GetJobQueue(agents)
	for _, ticket := range tickets {
		job := utils.Job{
			Name: "RESOLVE_TICKET_" + strconv.FormatInt(ticket.TicketID, 10),
			Fun:  models.ProcessTicket, Data: ticket.TicketID,
		}
		jobQueue.Insert(job)
	}
	return nil
}

func loadConfig() {
	configJSON, err := os.Open(envVariables["config"].(string))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer configJSON.Close()
	byteValue, err := ioutil.ReadAll(configJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("Configuration loaded")
}
