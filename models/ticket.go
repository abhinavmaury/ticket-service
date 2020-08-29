package models

import (
	"abhinav/ticket-service/db"
	"context"
	"errors"
	"time"
)

func (userTicket *UserTicket) SaveTicket(ctx *context.Context) error {
	query := `INSERT INTO user_ticket (user_id, phone, query, status) VALUES (?,?,?,?)`
	conn, err := db.UserTicketDB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	row, err := conn.ExecContext(*ctx, query, userTicket.UserID, userTicket.Phone, userTicket.Query, userTicket.Status)
	if err != nil {
		return err
	}
	if userTicket.TicketID, err = row.LastInsertId(); err != nil {
		return err
	}
	return nil
}

func (userTicket *UserTicket) UpdateStatus(ctx *context.Context, status string) error {
	conn, err := db.UserTicketDB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	args := []interface{}{}
	sqlStr := ""
	if status == "in_progress" {
		sqlStr = `update user_ticket set status = ? where id = ?`
		args = append(args, status)
	} else if status == "completed" {
		sqlStr = `update user_ticket set status = ? where id = ?`
		args = append(args, status)
	} else if status == "failed" {
		sqlStr = `update user_ticket set status = ? where id = ?`
		args = append(args, status)
		args = append(args, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		return errors.New("Invalid transition state")
	}

	args = append(args, userTicket.TicketID)
	_, err = conn.ExecContext(*ctx, sqlStr, args...)
	if err != nil {
		return err
	}
	userTicket.Status = status
	return nil
}

func (userTicket *UserTicket) FindByID(ctx *context.Context) error {
	conn, err := db.UserTicketDB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	sql := `select user_id, phone, query from user_ticket where id = ?`
	row := conn.QueryRowContext(*ctx, sql, userTicket.TicketID)
	err = row.Scan(&userTicket.UserID, &userTicket.Phone, &userTicket.Query)
	if err != nil {
		return errors.New("Does not exist")
	}
	return nil
}

func GetPendingTickets(ctx *context.Context) (userTicket []UserTicket, err error) {
	conn, err := db.UserTicketDB.Conn(*ctx)
	if err != nil {
		return nil, errors.New("ERROR_IN_DB_CONNECTION")
	}
	defer conn.Close()
	sql := "select id from user_ticket where status = ? order by created_at"
	rows, err := conn.QueryContext(*ctx, sql, "created")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ticket UserTicket
		err = rows.Scan(&ticket.TicketID)
		if err != nil {
			return nil, err
		}
		userTicket = append(userTicket, ticket)
	}
	return userTicket, nil
}
