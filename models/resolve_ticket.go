package models

import (
	"context"
	"encoding/json"
	"time"
)

type UserTicket struct {
	TicketID int64  `json:"ticket_id,omitempty"`
	UserID   string `json:"user_id"`
	Phone    string `json:"phone"`
	Query    string `json:"query"`
	Status   string `json:"status,omitempty"`
}

func ProcessTicket(ctx context.Context, data interface{}) error {
	var ticketID int64
	dataBytes, err := json.Marshal(data)
	err = json.Unmarshal(dataBytes, &ticketID)
	if err != nil {
		return err
	}
	ticket := &UserTicket{TicketID: ticketID}
	if err = ticket.FindByID(&ctx); err != nil {
		return err
	}
	time.Sleep(10 * time.Second)
	if err = ticket.UpdateStatus(&ctx, `in_progress`); err != nil {
		return err
	}
	time.Sleep(10 * time.Second)
	//do some process with the ticket
	if err = ticket.UpdateStatus(&ctx, `completed`); err != nil {
		return err
	}
	return nil
}
